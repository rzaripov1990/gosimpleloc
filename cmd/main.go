package main

import (
	"fmt"
	gsl "gosimpleloc"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	err := gsl.InitDictionary("./translations")
	if err != nil {
		panic(err)
	}

	en := gsl.LString("EN", "hello", "gosimpleloc")
	fmt.Println(en)

	ru := gsl.LString("RU", "hello", "gosimpleloc")
	fmt.Println(ru)

	kz := gsl.LString("KZ", "hello", "gosimpleloc")
	fmt.Println(kz)

	std := message.NewPrinter(language.Russian).Sprintf("hello", "gosimpleloc")
	fmt.Print(std)
}
