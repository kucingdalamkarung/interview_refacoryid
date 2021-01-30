package main

import (
	"encoding/json"
	"fmt"
	"interview_refacotryid/Soal_3/model"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

//Find items in Meeting Room.
//Find all electronic devices.
//Find all furnitures.
//Find all items was purchased at 16 Januari 2020.
//Find all items with brown color.

func main() {
	var basePath, _ = os.Getwd()
	var dataPath = fmt.Sprintf("%s/%s", basePath, "Soal_3/data.json")
	jsonFile, err := os.Open(dataPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer jsonFile.Close()

	value, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	var items model.Item
	_ = json.Unmarshal(value, &items)

	fmt.Println("No. 1")
	fmt.Println(strings.Repeat("-", 50))
	findMeetingRoomItems(items)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println()

	fmt.Println("No. 2")
	fmt.Println(strings.Repeat("-", 50))
	findItemByType(items, "electronic")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println()

	fmt.Println("No. 3")
	fmt.Println(strings.Repeat("-", 50))
	findItemByType(items, "furniture")
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println()

	fmt.Println("No. 4")
	fmt.Println(strings.Repeat("-", 50))
	findItemPurchasedAt(items)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println()

	fmt.Println("No. 5")
	fmt.Println(strings.Repeat("-", 50))
	findItemWithBrownColor(items)
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println()
}

func findMeetingRoomItems(data model.Item) {
	for _, val := range data {
		if val.Placement.Name == "Meeting Room" {
			fmt.Println("Inventory Id:", val.InventoryID)
			fmt.Println("Name:", val.Name)
			fmt.Println()
		}
	}
}

func findItemByType(data model.Item, typeName string) {
	for _, val := range data {
		if val.Type == typeName {
			fmt.Println("Inventory Id:", val.InventoryID)
			fmt.Println("Name:", val.Name)
			fmt.Println()
		}
	}
}

func findItemPurchasedAt(data model.Item) {
	for _, val := range data {
		unixTime := time.Unix(int64(val.PurchasedAt), 0)
		if unixTime.Day() == 16 && unixTime.Month() == 01 && unixTime.Year() == 2020 {
			fmt.Println("Inventory Id:", val.InventoryID)
			fmt.Println("Name:", val.Name)
			fmt.Println()
		}
	}
}

func findItemWithBrownColor(data model.Item) {
	for _, val := range data {
		for _, tag := range val.Tags {
			if tag == "brown" {
				fmt.Println("Inventory Id:", val.InventoryID)
				fmt.Println("Name:", val.Name)
				fmt.Println()
			}
		}
	}
}
