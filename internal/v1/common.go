package v1

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"

	"github.com/keenfury/library_api/internal/v1/book"
	boo "github.com/keenfury/library_api/internal/v1/book"
	// --- replace header text - do not remove ---
)

// Book
func SetupBook(eg *echo.Group) {
	//dl := &boo.DataBook{DB: stor.PsqlDB}
	dl := &boo.DataFileBook{}
	dm := boo.NewManagerBook(dl)
	dh := boo.NewHandlerBook(dm)
	dh.LoadBookRoutes(eg)
	// temp to load the file into memory
	content, err := ioutil.ReadFile("../../data/books.json")
	if err != nil {
		panic("Unable to load books file")
	}
	books := []book.Book{}
	err = json.Unmarshal(content, &books)
	if err != nil {
		panic("Unable to unmarshal books struct")
	}
	dl.Load(books)
}

// --- replace section text - do not remove ---
