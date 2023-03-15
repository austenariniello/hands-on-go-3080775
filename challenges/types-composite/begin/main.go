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
	title string
	author
}

// define a library type to track a list of books
type library struct {
	books map[string][]book
}

// define addBook to add a book to the library
func (l *library) addBook(newBook book) {
	l.books[newBook.author.name] = append(l.books[newBook.author.name], newBook)
}

// define a lookupByAuthor function to find books by an author's name
func (l library) lookupByAuthor(authorName string) (authorBooks []book) {
	authorBooks = l.books[authorName]
	return
}

func main() {
	// create a new library
	l := library{
		books: make(map[string][]book),
	}

	fmt.Println("Library Initialization: ", l.books)

	// add 2 books to the library by the same author
	tm := author{
			name: "Toni Morrison",
		}

	l.addBook(book{
		title: "The Bluest Eye",
		author: tm,
	})

	l.addBook(book{
		title: "Beloved",
		author: tm,
	})

	// add 1 book to the library by a different author
	jb := author{
			name: "James Baldwin",
		}

	l.addBook(book{
		title: "Giovanni's Room",
		author: jb,
	})

	// dump the library with spew
	spew.Dump(l)

	// lookup books by known author in the library
	tmBooks := l.lookupByAuthor("Toni Morrison")

	fmt.Println("-------------------------------------")
	spew.Dump(tmBooks)
	fmt.Println("-------------------------------------")

	// print out the first book's title and its author's name
	if len(tmBooks) > 0 {
		fmt.Println("First Book Title:", tmBooks[0].title, "by", tmBooks[0].author.name)
	}

}
