package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type (
	Report struct {
		Id       string  `json:"id"`
		Name     string  `json:"name"`
		Date     string  `json:"date"`
		Semester int     `json:"semester"`
		Studies  []Study `json:"studies"`
	}
	Study struct {
		StudyName   string `json:"study_name"`
		StudyCredit int    `json:"study_credit"`
		Grade       string `json:"grade"`
	}
)

func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	result, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	var report Report
	data, err := ReadFile(filename)
	if err != nil {
		return report, err
	}
	if err := json.Unmarshal(data, &report); err != nil {
		return report, err
	}
	return report, nil
}

var mapPenilaian = map[string]float64{
	"A":  4.0,
	"AB": 3.5,
	"B":  3.0,
	"BC": 2.5,
	"C":  2.0,
	"CD": 1.5,
	"D":  1.0,
	"DE": 0.5,
	"E":  0.0,
}

func GradePoint(report Report) float64 {
	var totalGrade float64
	var totalCredit float64
	if len(report.Studies) == 0 {
		return 0.0
	}
	for _, study := range report.Studies {
		grade, ok := mapPenilaian[study.Grade]
		if !ok {
			continue
		}
		totalGrade += grade * float64(study.StudyCredit)
		totalCredit += float64(study.StudyCredit)
	}
	return totalGrade / totalCredit
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
