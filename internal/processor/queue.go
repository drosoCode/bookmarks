package processor

import "fmt"

type BookmarkQueueData struct {
	ID   int64
	Link string
	Save bool
}

var bookmarkQueue chan BookmarkQueueData

func AddBookmark(id int64, link string, save bool) {
	bookmarkQueue <- BookmarkQueueData{
		ID:   id,
		Link: link,
		Save: save,
	}
}

func StartProcessor() {
	bookmarkQueue = make(chan BookmarkQueueData, 50)
	for {
		x := <-bookmarkQueue
		err := Process(x)
		if err != nil {
			fmt.Println(err)
		}
		if len(bookmarkQueue) == 0 {
			close()
		}
	}
}
