package utils

import (
	"fmt"
	"strconv"
	"strings"

	library "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-asargin-dev/models/library_models"
)

// The function that gets all books in the library.
func List_book(books []library.Book) {

	if len(books) == 0 {

		fmt.Println("Kitaplıkta kitap bulunmamaktadır.")

	} else {

		fmt.Print("\nSırasıyla ID - Kitap Adı - Stok Kodu - ISBN Numarası - Stok Adedi - Sayfa Sayısı - Fiyat - Silindi\n\n")

		for _, book := range books {

			fmt.Printf("%d - %s - %d - %d - %d Adet - %d Sayfa - %d TL - %t\n", book.Id, book.BookName, book.StockCode, book.IsbnNo, book.StockQuantity, book.NumberOfPages, book.Price, book.IsDeleted)

		}

	}

}

// The function that searches books within written arg inputs
func Search_book(search_params string, books []library.Book) {

	if len(search_params) == 0 {

		fmt.Println("Lütfen aramak istediğiniz kitabın adını yazınız.")

	} else {
		// Sonuçlar liste şeklinde istenildiğinden, İşlemler sonrası eşleşen sonuçları 'result' slice'ının içine append edeceğiz.
		result := []string{}

		for _, book := range books {
			//Büyük-küçük harf duyarlılığını kaldırmak için hem kitapları hem de arama parametrelerini lowercase'e çeviriyoruz.
			//İkinci ödev eklentileriyle beraber kullanıcı Stok Kodu ve ISBN no ile de arama yapabilmesi sağlanmıştır.
			if strings.Contains(strings.ToLower(book.BookName), strings.ToLower(search_params)) {
				result = append(result, book.BookName)

			} else if strings.Contains(strings.ToLower(strconv.Itoa(book.IsbnNo)), strings.ToLower(search_params)) {
				result = append(result, book.BookName)

			} else if strings.Contains(strings.ToLower(strconv.Itoa(book.StockCode)), strings.ToLower(search_params)) {
				result = append(result, book.BookName)
			}
		}

		if len(result) == 0 {
			fmt.Println("Aradığınız kriterde kitap bulunamamıştır.")
		} else {
			for _, value := range result {
				fmt.Println(value)
			}
		}

	}
}
