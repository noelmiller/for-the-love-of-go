package bookstore

import (
	"errors"
	"fmt"
)

type Book struct {
	Title           string
	Author          string
	Copies          int
	Id              int
	PriceCents      int
	DiscountPercent int
}

type Catalog map[int]Book

// var catalog []Book

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func (c Catalog) GetAllBooks() []Book {
	result := []Book{}
	for _, b := range c {
		result = append(result, b)
	}
	return result
}

func (c Catalog) GetBook(Id int) (Book, error) {
	b, ok := c[Id]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", Id)
	}
	return b, nil
}

func (b Book) NetPriceCents() int {
	saving := b.PriceCents * b.DiscountPercent / 100
	return b.PriceCents - saving
}
