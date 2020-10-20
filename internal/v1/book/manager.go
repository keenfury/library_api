package book

import (
	"encoding/json"
	"fmt"

	g "github.com/Jeffail/gabs"
	ae "github.com/keenfury/library_api/internal/api_error"
)

type (
	DataBookAdapter interface {
		Get(*Book) error
		List(*[]Book) error
		Post(*Book) error
		Put(*Book) error
		Delete(*Book) error
	}

	ManagerBook struct {
		dataBook DataBookAdapter
	}
)

func NewManagerBook(cboo DataBookAdapter) *ManagerBook {
	return &ManagerBook{dataBook: cboo}
}

func (m *ManagerBook) Get(boo *Book) error {
	if boo.Id < 1 {
		return ae.MissingParamError("Id")
	}
	return m.dataBook.Get(boo)
}

func (m *ManagerBook) List(boo *[]Book) error {
	return m.dataBook.List(boo)
}

func (m *ManagerBook) Post(boo *Book) error {
	if !boo.Author.Valid {
		return ae.MissingParamError("Author")
	}
	if len(boo.Author.ValueOrZero()) > 100 {
		return ae.StringLengthError("Author", 100)
	}
	if len(boo.ImageLink.ValueOrZero()) > 250 {
		return ae.StringLengthError("ImageLine", 250)
	}
	if len(boo.Lang.ValueOrZero()) > 50 {
		return ae.StringLengthError("Lang", 50)
	}
	if len(boo.Link.ValueOrZero()) > 250 {
		return ae.StringLengthError("Link", 250)
	}
	if !boo.Pages.Valid {
		return ae.MissingParamError("Pages")
	}
	if !boo.Title.Valid {
		return ae.MissingParamError("Title")
	}
	if len(boo.Title.ValueOrZero()) > 250 {
		return ae.StringLengthError("Title", 250)
	}
	if !boo.Year.Valid {
		return ae.MissingParamError("Year")
	}

	return m.dataBook.Post(boo)
}

func (m *ManagerBook) Put(body []byte) error {
	jParsed, errP := g.ParseJSON(body)
	if errP != nil {
		return ae.BindError(errP)
	}
	idFlt, okIdFlt := jParsed.Search("id").Data().(float64)
	if !okIdFlt {
		return ae.MissingParamError("Id")
	}
	id := int64(idFlt)

	boo := &Book{Id: id}
	errGet := m.Get(boo)
	if errGet != nil {
		return errGet
	}
	fmt.Printf("%+v\n", boo)
	// Author
	author, okAuthor := jParsed.Search("author").Data().(string)
	if okAuthor {
		boo.Author.Scan(author)
	}
	// ImageLink
	imageLink, okImageLink := jParsed.Search("image_link").Data().(string)
	if okImageLink {
		boo.ImageLink.Scan(imageLink)
	}
	// Lang
	lang, okLang := jParsed.Search("lang").Data().(string)
	if okLang {
		boo.Lang.Scan(lang)
	}
	// Link
	link, okLink := jParsed.Search("link").Data().(string)
	if okLink {
		boo.Link.Scan(link)
	}
	// Pages
	pages, okPages := jParsed.Search("pages").Data().(float64)
	if okPages {
		boo.Pages.Scan(int64(pages))
	}
	// Title
	title, okTitle := jParsed.Search("title").Data().(string)
	if okTitle {
		boo.Title.Scan(title)
	}
	// Year
	year, okYear := jParsed.Search("year").Data().(float64)
	if okYear {
		boo.Year.Scan(int64(year))
	}

	return m.dataBook.Put(boo)
}

func (m *ManagerBook) Delete(boo *Book) error {
	if boo.Id < 1 {
		return ae.MissingParamError("Id")
	}
	return m.dataBook.Delete(boo)
}

func ValidJson(jsonValue json.RawMessage) bool {
	bValue, err := jsonValue.MarshalJSON()
	if err != nil {
		return false
	}
	check := make(map[string]interface{}, 0)
	if errCheck := json.Unmarshal(bValue, &check); errCheck != nil {
		return false
	}
	return true
}
