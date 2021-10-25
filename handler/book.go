package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":   "Ahmad Usamah",
		"Alamat": "Jakarta",
	})
}

func (h *bookHandler) GetBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	b, err := h.bookService.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"http":   http.StatusOK,
		"data":   bookResponse,
	})
}

func (h *bookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	var booksResponse []book.BookResponse

	for _, b := range books {
		bookResponse := book.BookResponse{
			ID:          b.ID,
			Title:       b.Title,
			Description: b.Description,
			Price:       b.Price,
			Rating:      b.Rating,
		}

		booksResponse = append(booksResponse, bookResponse)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "Success",
		"http":       http.StatusOK,
		"total_data": len(booksResponse),
		"data":       booksResponse,
	})
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var bookRequset book.BookRequest

	err := c.ShouldBindJSON(&bookRequset)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"http":   http.StatusOK,
		"data":   book,
	})
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on Field %s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.Update(id, bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"http":   http.StatusOK,
		"data":   bookResponse,
	})
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)

	b, err := h.bookService.Delete(int(id))

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"errors": err,
		})
		return
	}

	bookResponse := book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Description: b.Description,
		Price:       b.Price,
		Rating:      b.Rating,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "Delete Success",
		"http":   http.StatusOK,
		"data":   bookResponse,
	})
}
