package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/crufter/goquery"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	strItem := os.Args[1:]
	var strName []string
	var strValue []string

	if len(strItem) != 1 { //Проверка на кол-во принимаемых аргументов
		fmt.Println("Error - invalid input method. Must be articul of item")
		return
	}
	for _, v := range strItem[0] { //Проверка на кол-во валидности аргумента
		if v < 48 || v > 57 {
			fmt.Println("Error - invalid input method. Must be numerlic")
			return
		}
	}
	res, _ := http.Get("https://shop.kz/search/?q=" + strItem[0])
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find(".bx_detail_chars .bx_detail_chars_i").Each(func(i int, s *goquery.Selection) {
		q := s.Find(".glossary-term").Text()
		strName = append(strName, q)
		p := s.Find(".bx_detail_chars_i_field").Text()
		strValue = append(strValue, p)
		fmt.Printf("%s %s\n", q, p)
	})
	// fmt.Println(strName,)
	// fmt.Println(strValue,)
	for i := range strName {
		fmt.Print(strName[i], ", ")
	}
	fmt.Println()
	for i := range strValue {
		fmt.Print(strValue[i], ", ")
	}
	fmt.Println()
}
