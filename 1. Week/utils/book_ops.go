package utils

import (
	"fmt"
	"strings"
)

// The function that gets all books in the library.
func List_book(books []string) {

	if len(books) == 0 {

		fmt.Println("Kitaplıkta kitap bulunmamaktadır.")

	} else {

		for _, book := range books {

			fmt.Println(book)

		}

	}

}

// The function that searches books within written arg inputs
func Search_book(search_params string, books []string) {

	if len(search_params) == 0 {

		fmt.Println("Lütfen aramak istediğiniz kitabın adını yazınız.")

	} else {
		// Sonuçlar liste şeklinde istenildiğinden, İşlemler sonrası eşleşen sonuçları 'result' slice'ının içine append edeceğiz.
		result := []string{}

		for _, book := range books {
			//Büyük-küçük harf duyarlılığını kaldırmak için hem kitapları hem de arama parametrelerini lowercase'e çeviriyoruz.
			if strings.Contains(strings.ToLower(book), strings.ToLower(search_params)) {
				result = append(result, book)
			}
		}

		if len(result) == 0 {
			fmt.Println("Aradığınız kriterde kitap bulunamamıştır.")
		} else {
			fmt.Println(result)
		}

	}
}
