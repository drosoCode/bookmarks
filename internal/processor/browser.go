package processor

import (
	"github.com/playwright-community/playwright-go"
)

var instance *playwright.Playwright
var browser *playwright.Browser

func getBrowser() (*playwright.Browser, error) {
	var err error = nil
	if instance == nil {
		instance, err = playwright.Run()
		if err != nil {
			return nil, err
		}
	}
	if browser == nil {
		b, err := (*instance).Chromium.Launch()
		browser = &b
		if err != nil {
			return nil, err
		}
	}
	return browser, nil
}

func close() error {
	if err := (*browser).Close(); err != nil {
		return err
	}
	if err := (*instance).Stop(); err != nil {
		return err
	}
	return nil
}