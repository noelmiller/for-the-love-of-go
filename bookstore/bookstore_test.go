package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
		Id:     1,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
	want := 1
	result, err := bookstore.Buy(book)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:  "Spart Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err := bookstore.Buy(book)
	if err == nil {
		t.Error("want error buying from zero copies, got nil")
	}
}

func TestGetAllBooks(t *testing.T) {
	catalog := bookstore.Catalog{
		1: {Title: "For the Love of Go"},
		2: {Title: "The Power of Go: Tools"},
	}

	want := []bookstore.Book{
		{Title: "For the Love of Go"},
		{Title: "The Power of Go: Tools"},
	}

	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool {
		return got[i].Id < got[j].Id
	})

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{
		1: {Id: 1, Title: "For the Love of Go"},
		2: {Id: 2, Title: "For the Love of God"},
	}
	want := bookstore.Book{Id: 1, Title: "For the Love of Go"}

	got, err := catalog.GetBook(1)

	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBookBadIdReturnsError(t *testing.T) {
	t.Parallel()
	catalog := bookstore.Catalog{}
	_, err := catalog.GetBook(3)
	if err == nil {
		t.Fatal(err)
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		PriceCents:      50,
		DiscountPercent: 50,
	}

	want := 25

	got := book.NetPriceCents()

	//	if err != nil {
	//		t.Fatal(err)
	//	}

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
