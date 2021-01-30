package main

import (
	"bufio"
	"fmt"
	"interview_refacotryid/Soal_1/models"
	"interview_refacotryid/Soal_1/utils"
	"os"
	"strconv"
	"time"
)

var Products []models.Product

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputDetails(reader)
}

func inputDetails(reader *bufio.Reader) {
	fmt.Print("Nama Warung: ")
	namaWarung, err := reader.ReadString('\n')
	if err != nil {
		utils.ErrorHandler("Hapar periksa kembali data yang dimasukan. ", err)
	}

	fmt.Print("Tanggal: ")
	tanggal, _ := reader.ReadString('\n')
	tanggalTransaksi, err := time.Parse("02/01/2006 15:04:05", utils.TrimReturn(tanggal))
	if err != nil {
		utils.ErrorHandler("Hapar periksa kembali data yang dimasukan. ", err)
	}

	fmt.Print("Kasir: ")
	kasir, err := reader.ReadString('\n')
	if err != nil {
		utils.ErrorHandler("Hapar periksa kembali data yang dimasukan. ", err)
	}

	var i = 1

	for {
		fmt.Print("Produk " + strconv.Itoa(i) + ": ")
		product, _ := reader.ReadString('\n')

		if product == "\n" {
			break
		}

		fmt.Print("Harga: ")
		harga, _ := reader.ReadString('\n')

		hargaProduct, err := strconv.ParseInt(utils.TrimReturn(harga), 10, 64)
		if err != nil {
			utils.ErrorHandler("Hapar periksa kembali data yang dimasukan. ", err)
		}

		var productDetail = models.Product{
			NamaProduct:  utils.TrimReturn(product),
			HargaProduct: hargaProduct,
		}
		Products = append(Products, productDetail)
		i++
	}

	detailWarung := models.Warung{
		NamaWarung: utils.TrimReturn(namaWarung),
		Tanggal:    tanggalTransaksi,
		Kasir:      utils.TrimReturn(kasir),
		Products:   Products,
	}

	printer := utils.Printer{Data: detailWarung}
	printer.PrintData()
}