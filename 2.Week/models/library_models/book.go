package library_models

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Created Struct for Book Informations

type Book struct {
	Id, StockCode, IsbnNo, NumberOfPages, StockQuantity, Price int
	BookName                                                   string
	AuthorInfo                                                 Author
	IsDeleted                                                  bool
}

var Books []Book

/* Constructor */
func (b *Book) CreateLibrary(books_slice []string) []Book {

	rand.Seed(time.Now().UnixNano())

	for index, value := range books_slice {
		var book Book
		book.Id = index + 1
		book.BookName = value
		book.StockCode = randomNumberGenerator(50000, 60000)
		book.IsbnNo = randomNumberGenerator(10000, 20000)
		book.NumberOfPages = randomNumberGenerator(30, 3000)
		book.StockQuantity = randomNumberGenerator(1, 100)
		book.Price = randomNumberGenerator(5, 250)
		book.IsDeleted = false
		book.AuthorInfo = book.AuthorInfo.CreateAuthor(value)
		Books = append(Books, book)
	}

	return Books

}

type Deletable interface {
	DeleteBook(id_string string)
}

// Kitap silme işlemlerini gerçekleştiren fonksiyon
func (b *Book) DeleteBook(id_string string) []Book {

	//İnputlar string tipinde geldiğinden int'e çevirip kontrollerini sağlıyoruz.
	id, err := strconv.Atoi(id_string)

	if err != nil {
		fmt.Println("Dikkat: ID'yi sayı olarak giriniz.")

	} else {
		// Eğer id kitaplığımızdaki id'lerden biriyse
		if id > 0 && id <= (len(Books)-1) {
			// Books indexi 0'dan id'imiz 1 den başladığından id-1 değerine sahip işlem yapılması istenen index değişkenimizi oluşturuyoruz.
			index := id - 1

			if !Books[index].IsDeleted {

				message_flag := false

				for _, v := range Books {

					if v.Id == id {

						message_flag = true
						Books[index].IsDeleted = true
						Books = append(Books[:index], Books[index+1:]...)

					}
				}

				if message_flag {

					fmt.Println("İşlem başarıyla gerçekleşti. Kitaplığın yeni hali :")

					fmt.Print("\nSırasıyla ID - Kitap Adı - Stok Kodu - Stok Adedi - Sayfa Sayısı - Fiyat - Silindi\n\n")

					for _, book := range Books {

						fmt.Printf("%d - %s - %d - %d Adet - %d Sayfa - %d TL - %t\n", book.Id, book.BookName, book.StockCode, book.StockQuantity, book.NumberOfPages, book.Price, book.IsDeleted)

					}

				} else {
					fmt.Println("Kitap daha önceden silinmiştir.")
				}

			} else {
				fmt.Println("Kitap daha önceden silinmiştir.")
			}

		} else {
			fmt.Println("Dikkat: Girdiğiniz ID'nin kitaplıkta bulunan kitap ID'lerinden biri olduğundan emin olunuz.")
		}
	}
	return Books
}

// Kitap satın alma işlemlerini gerçekleştiren fonksiyon
func (b *Book) BuyBook(id_string, quantity_string string) {

	//İnputlar string tipinde geldiğinden int'e çevirip kontrollerini sağlıyoruz.
	id, err := strconv.Atoi(id_string)

	if err != nil {
		fmt.Println("Dikkat: ID'yi sayı olarak giriniz.")

	} else {

		// Eğer id kitaplığımızdaki id'lerden biriyse
		if id > 0 && id <= (len(Books)-1) {

			quantity, err := strconv.Atoi(quantity_string)

			if err != nil {
				fmt.Println("Dikkat: Miktarı sayı olarak girdiğinizden emin olunuz.")
			}

			// Hali hazırda bulunan stok adedimizi 'stock_quantity' deeğişkenine atıyoruz.
			stock_quantity := Books[id-1].StockQuantity

			if quantity <= stock_quantity {

				// Miktarı aşmadığı takdirde satın alma işlemini gerçekleştiriyoruz.
				Books[id-1].StockQuantity = stock_quantity - quantity
				fmt.Printf("%s kitabından %d adet alındı. İyi günlerde okumanız dileğiyle.\n", Books[id-1].BookName, quantity)

			} else {
				fmt.Printf("\nDikkat: Malesef stoklarımızda talep ettiğiniz kadar kitap bulunmamaktadır.\nStoğumuzdaki miktar : %d\n", Books[id-1].StockQuantity)
			}

		} else {

			fmt.Println("Dikkat: Girdiğiniz ID'nin kitaplıkta bulunan kitap ID'lerinden biri olduğundan emin olunuz.")

		}
	}

}

func randomNumberGenerator(min, max int) int {

	random_number := rand.Intn(max-min+1) + min

	return random_number

}
