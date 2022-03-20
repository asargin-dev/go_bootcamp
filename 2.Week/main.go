package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	library "github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-asargin-dev/models/library_models"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-asargin-dev/utils"
)

// Book Structımıza ait metodumuzu kullanabilmek için bir struct nesnesi oluşturuyoruz.
var book library.Book

var Books []library.Book

func init() {

	books := []string{
		"A Tale of Two Cities",
		"The Hobbit",
		"The Little Prince",
		"The Alchemist",
		"Harry Potter and the Chamber of Secrets",
	}

	/* Program başladıktan sonra tanımladığımız string slice'ı
	models/library_models altındaki CreateLibrary constructor fonksiyonuna parametre
	olarak gönderilir ve structlar oluşturularak globalde tanımladığımız Books struct
	slice'ına return edilir.
	*/

	Books = book.CreateLibrary(books)

}

func main() {

	/*
		Önceki ödevimizde gereksinim kadarıyla console'a tek tek girilen komutlar ile
		işlemlerimizi gerçekleştiriyorduk.

		Bu ödevimizde her defasında go run main.go args... komutları verilerimizi yeniden oluşturacağından
		satın alma, stok düşme ve silme gibi işlemlerde inşa edilmiş veri üzerinde değişim gerektiğinden
		'bufio' paketinden Reader oluşturarak kullanıcı terminali terkedene kadar işlem yapılmasını sağlayalım.
	*/

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("---------------------")

	fmt.Println(
		`Hoşgeldiniz, Kitapp uygulamasında kullanabileceğiniz komutlar :

	Kitap satın almak için => buy (Örn. buy ID Quantity -- buy 1 10)

	Kitap aramak için => search  (Örn. search gollum)

	Kitapları listelemek için => list

	Kitaplıktan kitap silmek için => delete (Örn. delete ID Quantity )
	
	Uygulamadan çıkmak için => exit`)

	fmt.Println("---------------------")

	//Exit komutu ile işlemin break edilebilmesi için kod bloğumuzu isimlendiriyoruz.

InputLoop:

	//Kullanıcının her bir girdisini döngü ile handle ediyoruz.

	for {

		fmt.Print("->")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Bir hata oluştu.")
		}

		input = strings.Replace(input, "\n", "", -1)

		input_args := strings.Split(input, " ")

		//Checking args and routing associated functions
		switch strings.ToLower(input_args[0]) {

		//Satın alma
		case "buy":
			book.BuyBook(input_args[1], input_args[2])

		//Silme
		case "delete":
			Books = book.DeleteBook(input_args[1])

		//Kitap Listeleme
		case "list":
			utils.List_book(Books)

		//Kitap Arama
		case "search":
			// search argümanı sonrası girilececek değerleri Join ederek string bir arama querysi oluşturarak
			// kitaplarla beraber search_book fonksiyonumuza parametre olarak gönderiyoruz.
			utils.Search_book(strings.Join(input_args[1:], " "), Books)

		case "exit":
			break InputLoop

		default:
			fmt.Println("Geçersiz bir komut girdiniz.Lütfen tekrar deneyiniz.")
		}
	}
}
