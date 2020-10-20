package book

import (
	"gopkg.in/guregu/null.v3"
)

type Book struct {
	Id        int64       `db:"id" json:"id"`
	Author    null.String `db:"author" json:"author"`
	ImageLink null.String `db:"image_link" json:"imageLink"`
	Lang      null.String `db:"lang" json:"language"`
	Link      null.String `db:"link" json:"link"`
	Pages     null.Int    `db:"pages" json:"pages"`
	Title     null.String `db:"title" json:"title"`
	Year      null.Int    `db:"year" json:"year"`
}
