package utils

import (
	"fmt"
	"go-journey/basic/helper"
	"strconv"
)

func CallIotaLearnSection() {
	helper.CalledFunction("CallIotaLearnSection")

	strAngka := "124"

	strAngkaConv, err := strconv.Atoi(strAngka)

	if err != nil {
		fmt.Println("Error when convert strAngkaConv", err)
		return
	}

	fmt.Println("strAngka ", strAngka, "converted to int", strAngkaConv)

	intAngka := 123

	intAngkaConv := strconv.Itoa(intAngka)

	fmt.Println("Integer angka ", intAngka, " setelah di conversi ke string ", intAngkaConv)

	price := "19.99"
	priceFloat, err := strconv.ParseFloat(price, 64) // string ke float64

	if err != nil {
		fmt.Println("Error saat parse ke float", err)
		return
	}

	fmt.Println("price string=", price, "pric float ", priceFloat)
}
