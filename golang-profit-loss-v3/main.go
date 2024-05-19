package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var datas []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		datas = append(datas, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return []string{}, err
	}

	if len(datas) == 0 {
		return []string{}, nil
	}

	return datas, nil
}

func CalculateProfitLoss(data []string) string {
	// 1/1/2021;income;100000
	// 1/1/2021;expense;50000

	var totalIncome int
	var totalExpense int

	for _, d := range data {
		// split data
		splitedData := SplitData(d)

		// check type
		if splitedData[1] == "income" {
			totalIncome += ConvertToInt(splitedData[2])
		} else {
			totalExpense += ConvertToInt(splitedData[2])
		}
	}

	// calculate profit loss
	profitLoss := totalIncome - totalExpense

	// Karena, hasil hasil total income lebih besar dari total expense, maka terjadi keuntungan. Kita dapat mengembalikan nilai "4/1/2021;profit;500000"
	if profitLoss > 0 {
		return fmt.Sprintf("%s;profit;%d", SplitData(data[len(data)-1])[0], profitLoss)
	}
	// Karena, hasil hasil total income lebih kecil dari total expense, maka terjadi kerugian. Kita dapat mengembalikan nilai "4/1/2021;loss;500000"
	return fmt.Sprintf("%s;loss;%d", SplitData(data[len(data)-1])[0], profitLoss*-1)
}

func SplitData(data string) []string {
	return strings.Split(data, ";")
}

func ConvertToInt(data string) int {
	result, err := strconv.Atoi(data)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
