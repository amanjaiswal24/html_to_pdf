package main

import (
	"fmt"
	"log"

	"github.com/elcinzorlu/generate-pdf-and-upload-s3-go/pkg/converter"
)

func main() {
	r := converter.NewRequestPdf("")

	r.LocalFileAccess(true)
	// HTML template path
	templatePath := "templates/student_report.html"

	// Path for saving the generated PDF
	outputPath := "student_report.pdf"

	// HTML template data
	templateData := struct {
		FormNo          string
		StudentName     string
		CurrentBatch    string
		Stream          string
		CourseNameCode  string
		StudyCenterCode string
		MainTests       []Test
		SubjectiveTests []Test
		OtherTests      []Test
	}{
		FormNo:          "1000355023",
		StudentName:     "AMAN JAISWAL",
		CurrentBatch:    "MEL6B",
		Stream:          "JEE (MAIN+ADVANCED)",
		CourseNameCode:  "ENTHUSIAST ADVANCE PHASE-I B(303361)",
		StudyCenterCode: "KOTA, LANDMARK CITY, KUNHARI(1015)",
		MainTests: []Test{
			{
				TestName:   "MAJOR TEST (M36604460)",
				TestDate:   "18 Dec-24",
				Batch:      "MEL6B",
				SA:         5546,
				TMode:      "OFFLINE",
				Physics:    150,
				Chemistry:  149,
				Maths:      338,
				Total:      637,
				Percentage: 88.47,
				Percentile: 93.23,
				TestRank:   "21",
				AIR:        "23",
			},
			{
				TestName:   "MAJOR TEST (M36604460)",
				TestDate:   "18 Dec-24",
				Batch:      "MEL6B",
				SA:         5546,
				TMode:      "OFFLINE",
				Physics:    150,
				Chemistry:  149,
				Maths:      338,
				Total:      637,
				Percentage: 88.47,
				Percentile: 93.23,
				TestRank:   "21",
				AIR:        "23",
			},
			{
				TestName:   "MAJOR TEST (M36604460)",
				TestDate:   "18 Dec-24",
				Batch:      "MEL6B",
				SA:         5546,
				TMode:      "OFFLINE",
				Physics:    150,
				Chemistry:  149,
				Maths:      338,
				Total:      637,
				Percentage: 88.47,
				Percentile: 93.23,
				TestRank:   "21",
				AIR:        "23",
			},
		},
		SubjectiveTests: []Test{
			{
				TestName:   "SUBJECTIVE TEST (M36604460)",
				TestDate:   "18 Dec-24",
				Batch:      "MEL6B",
				SA:         5546,
				TMode:      "OFFLINE",
				Physics:    150,
				Chemistry:  149,
				Maths:      338,
				Total:      637,
				Percentage: 88.47,
				Percentile: 93.23,
				TestRank:   "-",
				AIR:        "-",
			},
			{
				TestName:   "MAJOR TEST (M36604460)",
				TestDate:   "18 Dec-24",
				Batch:      "MEL6B",
				SA:         5546,
				TMode:      "OFFLINE",
				Physics:    150,
				Chemistry:  149,
				Maths:      338,
				Total:      637,
				Percentage: 88.47,
				Percentile: 93.23,
				TestRank:   "21",
				AIR:        "23",
			},
		},
		OtherTests: []Test{
			{
				TestName:   "MINOR TEST (M36013940)",
				TestDate:   "13 Dec-24",
				Batch:      "MEL6B",
				SA:         5938,
				TMode:      "OFFLINE",
				Physics:    144,
				Chemistry:  152,
				Maths:      331,
				Total:      627,
				Percentage: 87.08,
				Percentile: 91.23,
				TestRank:   "-",
				AIR:        "-",
			},
			{
				TestName:   "MAJOR TEST (M36604460)",
				TestDate:   "18 Dec-24",
				Batch:      "MEL6B",
				SA:         5546,
				TMode:      "OFFLINE",
				Physics:    150,
				Chemistry:  149,
				Maths:      338,
				Total:      637,
				Percentage: 88.47,
				Percentile: 93.23,
				TestRank:   "21",
				AIR:        "23",
			},
			{
				TestName:   "MAJOR TEST (M36604460)",
				TestDate:   "18 Dec-24",
				Batch:      "MEL6B",
				SA:         5546,
				TMode:      "OFFLINE",
				Physics:    150,
				Chemistry:  149,
				Maths:      338,
				Total:      637,
				Percentage: 88.47,
				Percentile: 93.23,
				TestRank:   "21",
				AIR:        "23",
			},
		},
	}

	// Parse the HTML template with the provided data
	if err := r.ParseTemplateFile(templatePath, templateData); err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	// Generate the PDF
	if err := r.GeneratePDF(outputPath); err != nil {
		log.Fatalf("Failed to generate PDF: %v", err)
	}

	fmt.Printf("PDF successfully generated and saved at: %s\n", outputPath)
}

type Test struct {
	TestName   string
	TestDate   string
	Batch      string
	SA         int
	TMode      string
	Physics    int
	Chemistry  int
	Maths      int
	Total      int
	Percentage float64
	Percentile float64
	TestRank   string
	AIR        string
}
