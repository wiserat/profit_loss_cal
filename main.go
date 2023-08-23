package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Crypto futures P&L calculator\n-----------------------------")
	reader := bufio.NewReader(os.Stdin)
	const bitSize = 64
	for {
		var longOrShort int8
		fmt.Print("\n", "Short(0) or long(1): ")
		_, err := fmt.Scan(&longOrShort)
		if err != nil || !(longOrShort == 0 || longOrShort == 1) {
			fmt.Println(err, "or you didn't enter 0 or 1")
			continue
		}

		fmt.Print("Buy: ")
		s, _ := reader.ReadString('\n')
		s = strings.Replace(s, "\n", "", -1)
		buy, err := strconv.ParseFloat(s, bitSize)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Print("Sell: ")
		s, _ = reader.ReadString('\n')
		s = strings.Replace(s, "\n", "", -1)
		sell, err := strconv.ParseFloat(s, bitSize)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var leverage uint8
		fmt.Print("Leverage: ")
		_, err = fmt.Scan(&leverage)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Print("$$$: ")
		s, _ = reader.ReadString('\n')
		s = strings.Replace(s, "\n", "", -1)
		money, err := strconv.ParseFloat(s, bitSize)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if longOrShort == 0 {
			percentage := (buy - sell) / sell * 100
			if percentage > 0 {
				fmt.Println("\n", "|", "------------------------------")
				fmt.Println(" |", "Price movement:", percentage, "%")
				fmt.Println(" |", "Percentage profit:", percentage*float64(leverage), "%")
				fmt.Println(" |", "Profit:", percentage/100*float64(leverage)*money, "$")
				fmt.Println(" |", "------------------------------")
			} else {
				fmt.Println("\n", "|", "------------------------------")
				fmt.Println(" |", "Price movement:", percentage, "%")
				fmt.Println(" |", "Percentage profit:", percentage*float64(leverage), "%")
				fmt.Println(" |", "Loss:", percentage/100*float64(leverage)*money, "$")
				fmt.Println(" |", "------------------------------")
			}
		} else {
			percentage := (sell - buy) / buy * 100
			if percentage > 0 {
				fmt.Println("\n", "|", "------------------------------")
				fmt.Println(" |", "Price movement:", percentage, "%")
				fmt.Println(" |", "Percentage profit:", percentage*float64(leverage), "%")
				fmt.Println(" |", "Profit:", percentage/100*float64(leverage)*money, "$")
				fmt.Println(" |", "------------------------------")
			} else {
				fmt.Println("\n", "|", "------------------------------")
				fmt.Println(" |", "Price movement:", percentage, "%")
				fmt.Println(" |", "Percentage profit:", percentage*float64(leverage), "%")
				fmt.Println(" |", "Loss:", percentage/100*float64(leverage)*money, "$")
				fmt.Println(" |", "------------------------------")
			}
		}
	}
}
