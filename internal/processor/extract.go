package processor

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/playwright-community/playwright-go"
)

func screen(page playwright.Page, path string) error {
	if _, err := page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String(path),
	}); err != nil {
		return err
	}
	return nil
}

func getValue(page playwright.Page, selector string, prop string) string {
	js := "document.querySelector(\"" + strings.ReplaceAll(selector, "\"", "\\\"") + "\")." + prop
	val, err := page.Evaluate(js, playwright.FrameEvaluateOptions{})
	if err == nil {
		return val.(string)
	}
	return ""
}

func getTitle(page playwright.Page) string {
	val := getValue(page, "meta[name=\"twitter:title\"]", "content")
	if val != "" {
		return val
	}
	val = getValue(page, "meta[name=\"og:title\"]", "content")
	if val != "" {
		return val
	}
	return getValue(page, "title", "textContent")
}

func getDescription(page playwright.Page) string {
	val := getValue(page, "meta[name=\"twitter:description\"]", "content")
	if val != "" {
		return val
	}
	val = getValue(page, "meta[name=\"og:description\"]", "content")
	if val != "" {
		return val
	}
	return getValue(page, "meta[name=description]", "content")
}

func getImageUrl(page playwright.Page) string {
	val := getValue(page, "meta[name=\"twitter:image:src\"]", "content")
	if val != "" {
		return val
	}
	return getValue(page, "meta[name=\"og:image\"]", "content")
}

func getImage(page playwright.Page, path string) error {
	if val := getImageUrl(page); val != "" {
		return downloadFile(val, path)
	} else {
		return screen(page, path)
	}
}

func Test(link string) error {
	b, err := getBrowser()
	if err != nil {
		return err
	}
	page, err := (*b).NewPage()
	if err != nil {
		return err
	}
	savedFiles = []string{}
	savedEditablePaths = []string{}
	requestedUrl = link
	saveDir = "cache/1/html"
	os.MkdirAll(saveDir, os.ModePerm)

	page.On("response", saveFile)
	if _, err := page.Goto(link, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		return err
	}

	fmt.Println(getTitle(page))
	fmt.Println(getDescription(page))
	getImage(page, "cache/test.png")
	time.Sleep(2 * time.Second)
	patchFiles()

	page.Close()
	close()
	return nil
}
