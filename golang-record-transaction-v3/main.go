package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

type Recap struct {
	Date   string
	Profit string
	Amount int
}

func WriteFile(path string, data string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func SortByDate(arg map[string]Recap) []Recap {
	// Buat slice untuk menyimpan kunci (tanggal) dari map
	var keys []string
	for k := range arg {
		keys = append(keys, k)
	}

	// Urutkan kunci (tanggal) secara ascending
	sort.Strings(keys)

	// Buat slice untuk menyimpan hasil pengurutan
	var sortedRecaps []Recap
	for _, k := range keys {
		sortedRecaps = append(sortedRecaps, arg[k])
	}

	return sortedRecaps
}

func RecordTransactions(path string, transactions []Transaction) error {
	if len(transactions) == 0 {
		return nil
	}

	recaps := make(map[string]Recap)

	for _, t := range transactions {
		// Pisahkan data setiap date yang sama
		// jika key tidak ada, maka buat key baru
		if _, ok := recaps[t.Date]; !ok {
			recaps[t.Date] = Recap{}
		}

		// check type
		if t.Type == "income" {
			// Jika income maka tambahkan ke amount
			amout := recaps[t.Date].Amount + t.Amount
			recaps[t.Date] = Recap{t.Date, "income", amout}
		} else {
			// Jika expense maka kurangkan ke amount
			amout := recaps[t.Date].Amount - t.Amount
			recaps[t.Date] = Recap{t.Date, "expense", amout}
		}
	}

	// lakukan sort by date
	sortedByDate := SortByDate(recaps)

	// buat map jadi array
	var recapsArray []string
	for _, r := range sortedByDate {
		if r.Amount < 0 {
			r.Amount *= -1
			r.Profit = "expense"
		} else {
			r.Profit = "income"
		}
		recapsArray = append(recapsArray, fmt.Sprintf("%s;%s;%d", r.Date, r.Profit, r.Amount))
	}

	// Masukkan data ke file
	data := ""
	for _, r := range recapsArray {
		// jika last jangan tambahkan newline
		if r == recapsArray[len(recapsArray)-1] {
			data += r
		} else {
			data += r + "\n"
		}
	}

	return WriteFile(path, data)
}

func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "income", 200000},
		{"01/01/2021", "income", 1000},
		{"01/01/2021", "expense", 100000},
		{"01/01/2021", "expense", 1000},
		{"02/01/2021", "income", 20000},
		{"02/01/2021", "income", 233000},
		{"02/01/2021", "expense", 3424},
		{"02/01/2021", "expense", 2300},
		{"02/01/2021", "expense", 42000},
		{"03/01/2021", "income", 20000},
		{"03/01/2021", "income", 22000},
		{"03/01/2021", "expense", 22321},
		{"04/01/2021", "income", 24000},
		{"04/01/2021", "income", 20000},
		{"04/01/2021", "expense", 223200},
		{"05/01/2021", "income", 20000},
		{"05/01/2021", "income", 50000},
		{"05/01/2021", "expense", 2213},
		{"06/01/2021", "income", 60000},
		{"06/01/2021", "income", 70000},
		{"06/01/2021", "expense", 4545},
		{"07/01/2021", "income", 80000},
		{"07/01/2021", "income", 110000},
		{"07/01/2021", "income", 120000},
		{"07/01/2021", "income", 111200},
		{"07/01/2021", "expense", 55500},
		{"08/01/2021", "expense", 200000},
		{"10/01/2021", "expense", 20000},
		{"11/01/2021", "expense", 10000},
		{"12/01/2021", "expense", 55500},
		{"13/01/2021", "expense", 55500},
		{"02/01/2021", "expense", 55500},
		{"02/01/2021", "expense", 10000},
		{"14/01/2021", "expense", 20000},
		{"11/01/2021", "expense", 20000},
		{"15/01/2021", "expense", 10000},
		{"16/01/2021", "expense", 20000},
		{"02/01/2021", "expense", 55500},
		{"17/01/2021", "expense", 10000},
		{"06/01/2021", "expense", 20000},
		{"18/01/2021", "expense", 10000},
		{"03/01/2021", "expense", 20000},
		{"04/01/2021", "expense", 10000},
		{"19/01/2021", "expense", 55500},
		{"20/01/2021", "expense", 55500},
		{"21/01/2021", "expense", 10000},
		{"22/01/2021", "expense", 10000},
		{"23/01/2021", "expense", 10000},
		{"24/01/2021", "expense", 10000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
