// challenges/types-composite/begin/main.go
package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// define an author type with a name
type author struct {
	name string
}

// define a book type with a title and author
type book struct {
	author author
	title  string
}

// define a library type to track a list of books
type library map[string][]book

// define addBook to add a book to the library
func (l library) addBook(b book) {
	l[b.author.name] = append(l[b.author.name], b)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthorName(name string) []book {
	return l[name]
}

func main() {
	// create a new library
	var lib = make(library)

	// add 2 books to the library by the same author
	mt := author{name: "Mark Twain"}
	lib.addBook(book{
		title:  "Tom Sawyer",
		author: mt,
	})
	lib.addBook(book{
		title:  "Huckleberry Finn",
		author: mt,
	})

	// add 1 book to the library by a different author
	lib.addBook(book{
		title:  "The Da Vinci Code",
		author: author{name: "Dan Brown"},
	})

	// dump the library with spew
	spew.Dump(lib)

	// lookup books by known author in the library
	books := lib.lookupByAuthorName("Mark Twain")
	spew.Dump(books)

	// print out the first book's title and its author's name
	if len(books) != 0 {
		b := books[0]
		fmt.Printf("%s - %s\n", b.title, b.author.name)
	}
}
