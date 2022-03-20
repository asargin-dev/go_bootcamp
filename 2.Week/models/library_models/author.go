package library_models

// Created Struct for Author Informations

type Author struct {
	Name string
	Book string
}

// Author model for creating Author
func (a *Author) CreateAuthor(bookName string) Author {

	author_slice := []string{"Charles Dickens", "J.R.R Tolkien", "Antoine de Saint-Exupéry", "Paulo Coelho", "J.K. Rowling"}

	var author Author

	switch bookName {

	case "A Tale of Two Cities":
		author.Name = author_slice[0]
		author.Book = bookName

	case "The Hobbit":
		author.Name = author_slice[1]
		author.Book = bookName

	case " The Little Prince":
		author.Name = author_slice[2]
		author.Book = bookName

	case "The Little Prince":
		author.Name = author_slice[3]
		author.Book = bookName
	case "Harry Potter and the Chamber of Secrets":
		author.Name = author_slice[4]
		author.Book = bookName
	}

	return author
}
