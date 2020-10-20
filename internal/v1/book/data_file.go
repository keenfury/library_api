package book

type (
	DataFileBook struct{}
)

var books []Book
var id int64

func init() {
	books = []Book{}
	id = int64(1)
}

func (d *DataFileBook) Get(boo *Book) error {
	for _, b := range books {
		if b.Id == boo.Id {
			boo.Author = b.Author
			boo.ImageLink = b.ImageLink
			boo.Lang = b.Lang
			boo.Link = b.Link
			boo.Pages = b.Pages
			boo.Title = b.Title
			boo.Year = b.Year
		}
	}
	return nil
}

func (d *DataFileBook) List(boo *[]Book) error {
	*boo = books
	return nil
}

func (d *DataFileBook) Post(boo *Book) error {
	boo.Id = id
	id++
	books = append(books, *boo)
	return nil
}

func (d *DataFileBook) Put(boo *Book) error {
	for i, b := range books {
		if b.Id == boo.Id {
			books[i].Author = boo.Author
			books[i].ImageLink = boo.ImageLink
			books[i].Lang = boo.Lang
			books[i].Link = boo.Link
			books[i].Pages = boo.Pages
			books[i].Title = boo.Title
			books[i].Year = boo.Year
		}
	}
	return nil
}

func (d *DataFileBook) Delete(boo *Book) error {
	for i, b := range books {
		if b.Id == boo.Id {
			books = append(books[:i], books[i+1:]...)
		}
	}
	return nil
}

func (d *DataFileBook) Load(boo []Book) error {
	for i := range boo {
		boo[i].Id = id
		id++
	}
	books = append(books, boo...)
	return nil
}
