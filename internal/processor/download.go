package processor

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/playwright-community/playwright-go"
)

func downloadFile(url string, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

var savedFiles []string
var savedEditablePaths []string
var requestedUrl string
var saveDir string

func saveFile(response playwright.Response) {
	url := response.URL()
	pos := strings.LastIndex(url, "/")
	file := url[pos+1:]
	res := response.Request().ResourceType()
	if res == "document" {
		file = "index.html"
	}
	path := path.Join(saveDir, file)

	if (res == "image" || res == "script" || res == "stylesheet" || res == "document") && response.Status() == 200 {
		go func() {
			data, err := response.Body()
			if err != nil {
				return
			}
			err = os.WriteFile(path, data, 0644)
			if err != nil {
				return
			}

			savedFiles = append(savedFiles, file)
			if res == "script" || res == "document" {
				savedEditablePaths = append(savedEditablePaths, path)
			}
		}()
	}
}

func patchFiles() {
	for _, file := range savedFiles {
		for _, x := range savedEditablePaths {
			f, err := ioutil.ReadFile(x)
			if err == nil {
				fstr := string(f[:])
				r := `[-a-zA-Z0-9@:%._\+\/~#=]*\/` + regexp.QuoteMeta(file)
				rg1 := regexp.MustCompile(" " + r + " ")
				rg2 := regexp.MustCompile(`"` + r + `"`)
				rg3 := regexp.MustCompile(`"` + r + ` `)
				rg4 := regexp.MustCompile(` ` + r + `"`)
				fstr = rg1.ReplaceAllString(fstr, " ./"+file+" ")
				fstr = rg2.ReplaceAllString(fstr, "\"./"+file+"\"")
				fstr = rg3.ReplaceAllString(fstr, "\"./"+file+" ")
				fstr = rg4.ReplaceAllString(fstr, " ./"+file+"\"")
				os.WriteFile(x, []byte(fstr), 0644)
			}
		}
	}
}
