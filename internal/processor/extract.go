package processor

import (
	"github.com/playwright-community/playwright-go"
	"fmt"
)

func screen(page playwright.Page, path string) error {
	if _, err := page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String(path),
	}); err != nil {
		return err
	}
	return nil
}

func download(page playwright.Page) error {
	page.On("response", func(response playwright.Response) {
		//fmt.Printf("<< %v %s\n", response.Status(), response.URL())
	})
	return nil
}

func getValue(page playwright.Page, selector string) string {
	fmt.Println(selector)
	titleElement, err := page.QuerySelectorAll(selector)
	fmt.Println(titleElement)
	if err == nil && len(titleElement) > 0 {
		title, err := titleElement[0].TextContent()
		if err != nil {
			return title
		}
	}
	return ""
}

func getTitle(page playwright.Page) string {
	val := getValue(page, "meta[name=\"twitter:title\"]")
	if val != "" {
		return val
	}
	val = getValue(page, "meta[name=\"og:title\"]")
	if val != "" {
		return val
	}
	return getValue(page,"title")
}

func getDescription(page playwright.Page) string {
	val := getValue(page, "meta[name=\"twitter:description\"]")
	if val != "" {
		return val
	}
	val = getValue(page, "meta[name=\"og:description\"]")
	if val != "" {
		return val
	}
	return getValue(page,"meta[name=description]")
}

func getImageUrl(page playwright.Page) string {
	val := getValue(page, "meta[name=\"twitter:image:src\"]")
	if val != "" {
		return val
	}
	return getValue(page, "meta[name=\"og:image\"]")
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

	//download(page)
	_, err = page.WaitForSelector("title", playwright.PageWaitForSelectorOptions{
		State: playwright.WaitForSelectorStateAttached,
	})
	if err != nil {
		return err
	}

	if _, err := page.Goto(link, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		return err
	}

	fmt.Println(getTitle(page))
	//fmt.Println(getDescription(page))
	//getImage(page, "cache/test.png")
	
	page.Close()
	close()
	return nil
}

