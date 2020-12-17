# gosimpleloc

Простой пакет обертка над стандартными библиотеками go:
golang.org/x/text/language, golang.org/x/text/message

Файлы локализации в __YAML__ формате необходимо поместить в папку **translations** 
Название файлов должно содержать код языка (подробней https://pkg.go.dev/golang.org/x/text/language)

## Пример использования

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


## Запуск примера
make run