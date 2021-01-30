package main

import (
	"encoding/json"
	"fmt"
	"interview_refacotryid/Soal_2/model"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	var basePath, _ = os.Getwd()
	var dataPath = fmt.Sprintf("%s/%s", basePath, "Soal_2/data.json")
	jsonFile, err := os.Open(dataPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer jsonFile.Close()

	value, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	var users model.User
	_ = json.Unmarshal(value, &users)

	fmt.Println("No.1")
	fmt.Println(strings.Repeat("-", 50))
	findUserWithoutPhoneNumber(users)
	fmt.Println(strings.Repeat("-", 50))
	println()

	fmt.Println("No.2")
	fmt.Println(strings.Repeat("-", 50))
	findUserHaveArticles(users)
	fmt.Println(strings.Repeat("-", 50))
	println()

	fmt.Println("No.3")
	fmt.Println(strings.Repeat("-", 50))
	findUserWithSpecificName(users)
	fmt.Println(strings.Repeat("-", 50))
	println()

	fmt.Println("No.4")
	fmt.Println(strings.Repeat("-", 50))
	findUserWithArticlesSpecificYear(users)
	fmt.Println(strings.Repeat("-", 50))
	println()

	fmt.Println("No.5")
	fmt.Println(strings.Repeat("-", 50))
	findUserBornOn(users)
	fmt.Println(strings.Repeat("-", 50))
	println()

	fmt.Println("No.6")
	fmt.Println(strings.Repeat("-", 50))
	findSpecificArticles(users)
	fmt.Println(strings.Repeat("-", 50))
	println()

	fmt.Println("No.7")
	fmt.Println(strings.Repeat("-", 50))
	findArticlesPublishBefore(users)
	fmt.Println(strings.Repeat("-", 50))
}

func findUserWithoutPhoneNumber(data model.User) {
	for _, val := range data {
		if len(val.Profile.Phones) == 0 {
			fmt.Println("Id: ", val.ID)
			fmt.Println("Username: ", val.Username)
		}
	}
}

func findUserHaveArticles(data model.User) {
	for _, val := range data {
		if len(val.Articles) != 0 {
			fmt.Println("Id: ", val.ID)
			fmt.Println("Username: ", val.Username)
		}
	}
}

func findUserWithSpecificName(data model.User) {
	for _, val := range data {
		if strings.Contains(strings.ToLower(val.Profile.FullName), "annis") {
			fmt.Println("Id: ", val.ID)
			fmt.Println("Username: ", val.Username)
		}
	}
}

func findSpecificArticles(data model.User) {
	for _, val := range data {
		for _, article := range val.Articles {
			if strings.Contains(strings.ToLower(article.Title), "tips") {
				fmt.Println("Title:", article.Title)
				fmt.Println("Publish:", article.PublishedAt)
				fmt.Println()
			}
		}
	}
}

func findUserWithArticlesSpecificYear(data model.User) {
	var layout = "2006-01-02T15:04:05"
	for _, val := range data {
		for _, article := range val.Articles {
			publish, err := time.Parse(layout, article.PublishedAt)
			if err != nil {
				fmt.Println(err)
			}

			if publish.Year() == 2020 {
				fmt.Println("Id: ", val.ID)
				fmt.Println("Username: ", val.Username)
			}
		}
	}
}

func findUserBornOn(data model.User) {
	layout := "2006-01-02"
	for _, val := range data {
		born, err := time.Parse(layout, val.Profile.Birthday)
		if err != nil {
			fmt.Println(err)
		}
		if born.Year() == 1986 {
			fmt.Println("Id: ", val.ID)
			fmt.Println("Username: ", val.Username)
		}
	}
}

func findArticlesPublishBefore(data model.User) {
	var layout = "2006-01-02T15:04:05"
	for _, val := range data {
		for _, article := range val.Articles {
			publish, err := time.Parse(layout, article.PublishedAt)
			if err != nil {
				fmt.Println(err)
			}

			if publish.Year() < 2019 && publish.Month() < 8 {
				fmt.Println("Title:", article.Title)
				fmt.Println("Publish:", article.PublishedAt)
			}
		}
	}
}