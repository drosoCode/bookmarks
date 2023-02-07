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
		fmt.Println("processing error: " + err.Error())
		return err
	}
	page, err := (*b).NewPage()
	if err != nil {
		fmt.Println("processing error: " + err.Error())
		return err
	}
	defer page.Close()

	saveBaseDir := path.Join(config.Data.CachePath, strconv.FormatInt(data.ID, 10))
	saveDir := path.Join(saveBaseDir, "html")
	err = os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		fmt.Println("processing error: " + err.Error())
		return err
	}

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
		fmt.Println("processing error: " + err.Error())
		return err
	}

	err = database.DB.SetBookmarkData(context.Background(), database.SetBookmarkDataParams{
		Name:        getTitle(page),
		Description: getDescription(page),
		ID:          data.ID,
	})
	if err != nil {
		fmt.Println("processing error: " + err.Error())
		return err
	}

	err = getImage(page, path.Join(saveBaseDir, "image.png"))
	if err != nil {
		fmt.Println("processing error: " + err.Error())
		return err
	}

	if data.Save {
		sData.WG.Wait()
		patchFiles()
	}

	fmt.Println("processing ok")
	return nil
}
