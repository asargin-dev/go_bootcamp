package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-1-asargin-dev/utils"
)

func main() {

	args := os.Args

	books := []string{
		"A Tale of Two Cities",
		"The Hobbit",
		"The Little Prince",
		"The Alchemist",
		"Harry Potter and the Chamber of Secrets",
	}

	//Checking args and routing associated functions
	if len(args) > 1 {

		switch args[1] {

		case "list":

			utils.List_book(books)

		case "search":
			// search argümanı sonrası girilececek değerleri Join ederek string bir arama querysi oluşturarak
			// kitaplarla beraber search_book fonksiyonumuza parametre olarak gönderiyoruz.
			utils.Search_book(strings.Join(args[2:], " "), books)

		}

	} else {

		fmt.Printf("%s uygulamasında kullanabileceğiniz komutlar : \n search => arama işlemi için \n list => listeleme işlemi için\n", "Kitaplık")

	}

}
