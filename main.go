/**
 *  created by malevinsky
 *  email: malevinskaya2000@gmail.com
 *  telegram: @theos_deus
 *  date: 10.05.2021
 */

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Currency struct {
	Success   bool   `json:"success"`
	Timestamp int    `json:"timestamp"`
	Base      string `json:"base"`
	Date      string `json:"date"`
	Rates     Rates  `json:"rates"`
}

type Rates struct {
	USD float64 `json:"USD"`
	AUD float64 `json:"AUD"`
	CAD float64 `json:"CAD"`
	PLN float64 `json:"PLN"`
	MXN float64 `json:"MXN"`
}

func main() {

	printing()

	var i int
	fmt.Scan(&i)

	url := "http://api.exchangeratesapi.io/v1/latest?access_key=9c484230306ca3014e2eb4c8575de8df&symbols=USD,AUD,CAD,PLN,MXN&format=1"

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	Currency1 := Currency{}
	jsonErr := json.Unmarshal(body, &Currency1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	switching(i, Currency1)
}

func printing() int {
	/**
	Печатаем условия в отдельной функции, чтобы main выглядел более структурно.
	*/

	fmt.Println("\n\n\nКурс какой валюты к евро вы хотите узнать?\n1. USD(доллар)\n2. AUD(австралийский доллар)\n3. CAD(канадский доллар)\n4. PLN(польский злотый)\n5. MXN(мексиканское песо)\n\nНаберите номер валюты:\n")
	return 1
}

func switching(i int, Currency1 Currency) int {
	/**
	Тут мы проверяем какую валюту выбрал пользователь. Использую обычный switch, чтобы избежать множества if'ов.
	*/

	switch i {
	case 1:
		fmt.Printf("\nКурс доллара к евро: %v\n\n", Currency1.Rates.USD)
	case 2:
		fmt.Printf("\nКурс австралийского доллара к евро: %v\n\n", Currency1.Rates.AUD)
	case 3:
		fmt.Printf("\nКурс канадского доллара к евро: %v\n\n", Currency1.Rates.CAD)
	case 4:
		fmt.Printf("\nКурс польского злотого к евро: %v\n\n", Currency1.Rates.PLN)
	case 5:
		fmt.Printf("\nКурс мексиканского песо к евро: %v\n\n", Currency1.Rates.MXN)
	}
	return 1
}
