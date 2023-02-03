package processor

import (
	"context"
	"fmt"
	"os"
	"path"
	"strconv"
	"sync"

	"github.com/drosocode/bookmarks/internal/config"
	"github.com/drosocode/bookmarks/internal/database"
	"github.com/playwright-community/playwright-go"
)

func Process(data BookmarkQueueData) error {
	fmt.Printf("processing %s\n", data.Link)

	b, err := getBrowser()
	if err != nil {
		return err
	}
	page, err := (*b).NewPage()
	if err != nil {
		return err
	}

	saveBaseDir := path.Join(config.Data.CachePath, strconv.FormatInt(data.ID, 10))
	saveDir := path.Join(saveBaseDir, "html")
	os.MkdirAll(saveDir, os.ModePerm)

	if data.Save {
		sData = SharedData{
			SaveDir:            saveDir,
			SavedFiles:         []string{},
			SavedEditablePaths: []string{},
			SavedDocument:      false,
			WG:                 sync.WaitGroup{},
		}
		page.On("response", saveFile)
	}
	if _, err := page.Goto(data.Link, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateNetworkidle,
	}); err != nil {
		return err
	}

	database.DB.SetBookmarkData(context.Background(), database.SetBookmarkDataParams{
		Name:        getTitle(page),
		Description: getDescription(page),
		ID:          data.ID,
	})
	getImage(page, path.Join(saveBaseDir, "image.png"))

	if data.Save {
		sData.WG.Wait()
		patchFiles()
	}

	page.Close()
	return nil
}
