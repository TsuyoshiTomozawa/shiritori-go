package main

import "fmt"

func main() {
	currencies := make(map[string]string)

	currencies["日本"] = "JPY"
	currencies["USA"] = "USD"
	currencies["EU"] = "EUR"
	currencies["中国"] = "CNY"

	// キーを指定して値を取得
	fmt.Println(currencies["日本"])
	fmt.Println(currencies["中国"])

	fmt.Println("==すべて取得=======")

	for country, currency := range currencies {
		fmt.Println(country, currency)
	}

}
