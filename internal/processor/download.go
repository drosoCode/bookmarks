package processor

import (
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"

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

type SharedData struct {
	SavedFiles         []string
	SavedEditablePaths []string
	SavedDocument      bool
	SaveDir            string
	WG                 sync.WaitGroup
}

var sData SharedData

func saveFile(response playwright.Response) {
	url := response.URL()
	pos := strings.LastIndex(url, "/")
	file := url[pos+1:]
	res := response.Request().ResourceType()
	if res == "document" && !sData.SavedDocument {
		file = "index.html"
		sData.SavedDocument = true
	}
	path := path.Join(sData.SaveDir, file)

	if (res == "image" || res == "script" || res == "stylesheet" || res == "document") && response.Status() == 200 {
		sData.WG.Add(1)
		go func() {
			defer sData.WG.Done()
			data, err := response.Body()
			if err != nil {
				return
			}
			err = os.WriteFile(path, data, 0644)
			if err != nil {
				return
			}

			sData.SavedFiles = append(sData.SavedFiles, file)
			if res == "script" || res == "document" {
				sData.SavedEditablePaths = append(sData.SavedEditablePaths, path)
			}
		}()
	}
}

func patchFiles() {
	for _, file := range sData.SavedFiles {
		for _, x := range sData.SavedEditablePaths {
			f, err := os.ReadFile(x)
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
