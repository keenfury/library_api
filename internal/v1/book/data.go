package book

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	// temp
	ae "github.com/keenfury/library_api/internal/api_error"
)

type (
	DataBook struct {
		DB *sqlx.DB
	}
)

func (d *DataBook) Get(boo *Book) error {
	sqlGet := `
		select
			id,
			author,
			image_link,
			lang,
			link,
			pages,
			title,
			year
		from book where id = $1`
	if errDB := d.DB.Get(boo, sqlGet, boo.Id); errDB != nil {
		return ae.DBError("Book Get: unable to get record.", errDB)
	}
	return nil
}

func (d *DataBook) List(boo *[]Book) error {
	sqlList := `
		select
			id,
			author,
			image_link,
			lang,
			link,
			pages,
			title,
			year
		from book order by id`
	if errDB := d.DB.Select(boo, sqlList); errDB != nil {
		return ae.DBError("Book List: unable to select records.", errDB)
	}
	return nil
}

func (d *DataBook) Post(boo *Book) error {
	sqlPost := `
		insert into book (
			author,
			image_link,
			lang,
			link,
			pages,
			title,
			year
		) values (
			:author,
			:image_link,
			:lang,
			:link,
			:pages,
			:title,
			:year
		) returning id`
	rows, errDB := d.DB.NamedQuery(sqlPost, boo)
	if errDB != nil {
		return ae.DBError("Book Post: unable to insert record.", errDB)
	}
	defer rows.Close()
	var lastId int64
	if rows.Next() {
		rows.Scan(&lastId)
	}
	boo.Id = lastId

	return nil
}

func (d *DataBook) Put(boo *Book) error {
	sqlPut := `
		update book set
			author = :author,
			image_link = :image_link,
			lang = :lang,
			link = :link,
			pages = :pages,
			title = :title,
			year = :year
		where id = :id`
	if _, errDB := d.DB.NamedExec(sqlPut, boo); errDB != nil {
		return ae.DBError("Book Put: unable to update record.", errDB)
	}
	return nil
}

func (d *DataBook) Delete(boo *Book) error {
	sqlDelete := `
		delete from book where id = $1`
	if _, errDB := d.DB.Exec(sqlDelete, boo.Id); errDB != nil {
		return ae.DBError("Book Delete: unable to delete record.", errDB)
	}
	return nil
}

func (d *DataBook) Load(boo []Book) error {
	fmt.Println("load")
	sqlTruncate := `truncate table book`
	_, err := d.DB.Exec(sqlTruncate)
	if err != nil {
		return fmt.Errorf("Error in truncating book table: %s", err)
	}
	for _, b := range boo {
		err := d.Post(&b)
		if err != nil {
			return fmt.Errorf("Error inserting book: %s", err)
		}
	}
	return nil
}
