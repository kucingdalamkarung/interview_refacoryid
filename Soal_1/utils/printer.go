package utils

import (
	"fmt"
	"interview_refacotryid/Soal_1/models"
	"strconv"
	"strings"
)

type Printer struct {
	Data interface{}
}

func (p Printer) PrintData() {
	data := p.Data.(models.Warung)
	w := 30
	fmt.Printf(fmt.Sprintf("%%-%ds", w/2), fmt.Sprintf(fmt.Sprintf("%%%ds", w/2),data.NamaWarung))
	fmt.Println()
	fmt.Println("Tanggal:", data.Tanggal.Format("02/01/2006 15:04:05"))
	fmt.Println("Name Kasir:", data.Kasir)
	fmt.Println(strings.Repeat(string('='), 30))
	for _, product := range data.Products {
		lenProduct := countString(product.NamaProduct)
		lenHarga := countString("Rp"+strconv.FormatInt(product.HargaProduct, 10))
		totalLen := lengthString(lenHarga, lenProduct)
		fmt.Printf("%s%sRp.%d\n", product.NamaProduct, printDot(totalLen-2), product.HargaProduct)
		fmt.Println()
	}
	totalPrice := getTotalPrice(data.Products)
	totalLength := strconv.Itoa(int(totalPrice))
	fmt.Printf("Total%sRp.%d", printDot((len(totalLength) + 5) - 2), totalPrice)
}

func countString(str string) int {
	return len(str)
}

func lengthString(data ...int) int {
	sum := 0

	for i := range data {
		sum += data[i]
	}

	return sum
}

func printDot(length int) string {
	return strings.Repeat(string('.'), 30-length)
}

func getTotalPrice(product []models.Product) int64 {
	var sum int64 = 0
	for _, val := range product {
		sum += val.HargaProduct
	}

	return sum
}