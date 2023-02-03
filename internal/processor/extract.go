package processor

import (
	"strings"

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
