package book

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/labstack/echo"

	ae "github.com/keenfury/library_api/internal/api_error"
)

type (
	ManagerBookAdapter interface {
		Get(*Book) error
		List(*[]Book) error
		Post(*Book) error
		Put([]byte) error
		Delete(*Book) error
	}

	HandlerBook struct {
		managerBook ManagerBookAdapter
	}
)

func NewHandlerBook(mboo ManagerBookAdapter) *HandlerBook {
	return &HandlerBook{managerBook: mboo}
}

func (h *HandlerBook) LoadBookRoutes(eg *echo.Group) {
	eg.GET("/book/:id", h.Get)
	eg.GET("/book/list", h.List)
	eg.POST("/book", h.Post)
	eg.PUT("/book", h.Put)
	eg.DELETE("/book/:id", h.Delete)
}

func (h *HandlerBook) Get(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		bindErr := ae.ParseError("Invalid param value, not a number")
		return c.JSON(bindErr.StatusCode, bindErr.BodyError())
	}
	book := &Book{Id: id}
	if err := h.managerBook.Get(book); err != nil {
		be := err.(ae.ApiError).BodyError()
		return c.JSON(be.StatusCode, be)
	}
	return c.JSON(http.StatusOK, book)
}

func (h *HandlerBook) List(c echo.Context) error {
	books := &[]Book{}
	if err := h.managerBook.List(books); err != nil {
		be := err.(ae.ApiError).BodyError()
		return c.JSON(be.StatusCode, be)
	}
	return c.JSON(http.StatusOK, books)
}

func (h *HandlerBook) Post(c echo.Context) error {
	boo := &Book{}
	if err := c.Bind(boo); err != nil {
		bindErr := ae.BindError(err)
		return c.JSON(bindErr.StatusCode, bindErr.BodyError())
	}
	if err := h.managerBook.Post(boo); err != nil {
		be := err.(ae.ApiError).BodyError()
		return c.JSON(be.StatusCode, be)
	}
	return c.JSON(http.StatusOK, boo)
}

func (h *HandlerBook) Put(c echo.Context) error {
	body, _ := ioutil.ReadAll(c.Request().Body)
	if err := h.managerBook.Put(body); err != nil {
		be := err.(ae.ApiError).BodyError()
		return c.JSON(be.StatusCode, be)
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *HandlerBook) Delete(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		bindErr := ae.ParseError("Invalid param value, not a number")
		return c.JSON(bindErr.StatusCode, bindErr.BodyError())
	}
	book := &Book{Id: id}
	if err := h.managerBook.Delete(book); err != nil {
		be := err.(ae.ApiError).BodyError()
		return c.JSON(be.StatusCode, be)
	}
	return c.NoContent(http.StatusNoContent)
}
