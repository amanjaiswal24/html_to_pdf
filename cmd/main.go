package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/elcinzorlu/generate-pdf-and-upload-s3-go/pkg/converter"
)

const (
	S3_BUCKET = "my-bucket" // Bucket
	S3_REGION = "us-east-1" // Region
)

func main() {
	r := converter.NewRequestPdf("")

	r.LocalFileAccess(true)
	//html template path
	templatePath := "templates/student_report.html"

	//path for download pdf
	outputPath := "student_report.pdf"

	//html template data
	templateData := struct {
		Title           string
		AcademicSession string
		Address         string
		FormNo          string
		StudentName     string
		CurrentBatch    string
		Stream          string
		CourseNameCode  string
		StudyCenterCode string
		Tests           []Test
	}{
		Title:           "Student Performance Report",
		AcademicSession: "2024-2025",
		Address:         "SAMYAK, LANDMARK CITY, KUNHARI, BUNDI ROAD, KOTA, KOTA (RAJASTHAN)",
		FormNo:          "1000355023",
		StudentName:     "ADITYA SAHU",
		CurrentBatch:    "MEL6B",
		Stream:          "PRE-MEDICAL",
		CourseNameCode:  "ENTHUSIAST ADVANCE PHASE-I B(303361)",
		StudyCenterCode: "KOTA, LANDMARK CITY, KUNHARI(1015)",
		Tests: []Test{
			{
				TestName:   "MAJOR TEST (M36604460)",
				TestDate:   "18 Dec-24",
				Batch:      "MEL6B",
				SA:         5546,
				TMode:      "OFFLINE",
				Physics:    150,
				Chemistry:  149,
				Biology:    338,
				Total:      637,
				Percentage: 88.47,
				Percentile: 93.23,
				TestRank:   "-",
				AIR:        "-",
			},
			{
				TestName:   "MAJOR TEST (M36013940)",
				TestDate:   "13 Dec-24",
				Batch:      "MEL6B",
				SA:         5938,
				TMode:      "OFFLINE",
				Physics:    144,
				Chemistry:  152,
				Biology:    331,
				Total:      627,
				Percentage: 87.08,
				Percentile: 91.23,
				TestRank:   "-",
				AIR:        "-",
			},
			{
				TestName:   "MAJOR TEST (M36013940)",
				TestDate:   "13 Dec-24",
				Batch:      "MEL6B",
				SA:         5938,
				TMode:      "OFFLINE",
				Physics:    144,
				Chemistry:  152,
				Biology:    331,
				Total:      627,
				Percentage: 87.08,
				Percentile: 91.23,
				TestRank:   "-",
				AIR:        "-",
			}, {
				TestName:   "MAJOR TEST (M36013940)",
				TestDate:   "13 Dec-24",
				Batch:      "MEL6B",
				SA:         5938,
				TMode:      "OFFLINE",
				Physics:    144,
				Chemistry:  152,
				Biology:    331,
				Total:      627,
				Percentage: 87.08,
				Percentile: 91.23,
				TestRank:   "-",
				AIR:        "-",
			}, {
				TestName:   "MAJOR TEST (M36013940)",
				TestDate:   "13 Dec-24",
				Batch:      "MEL6B",
				SA:         5938,
				TMode:      "OFFLINE",
				Physics:    144,
				Chemistry:  152,
				Biology:    331,
				Total:      627,
				Percentage: 87.08,
				Percentile: 91.23,
				TestRank:   "-",
				AIR:        "-",
			}, {
				TestName:   "MAJOR TEST (M36013940)",
				TestDate:   "13 Dec-24",
				Batch:      "MEL6B",
				SA:         5938,
				TMode:      "OFFLINE",
				Physics:    144,
				Chemistry:  152,
				Biology:    331,
				Total:      627,
				Percentage: 87.08,
				Percentile: 91.23,
				TestRank:   "-",
				AIR:        "-",
			}, {
				TestName:   "MAJOR TEST (M36013940)",
				TestDate:   "13 Dec-24",
				Batch:      "MEL6B",
				SA:         5938,
				TMode:      "OFFLINE",
				Physics:    144,
				Chemistry:  152,
				Biology:    331,
				Total:      627,
				Percentage: 87.08,
				Percentile: 91.23,
				TestRank:   "-",
				AIR:        "-",
			},
		},
	}

	if err := r.ParseTemplateFile(templatePath, templateData); err != nil {
		log.Fatal(err)
	}
	if err := r.GeneratePDF(outputPath); err != nil {
		log.Fatal(err)
	}
	fmt.Println("pdf generated successfully")

	sess, err := session.NewSession(&aws.Config{Region: aws.String(S3_REGION)})
	if err != nil {
		log.Fatalf("session.NewSession - filename: %v, err: %v", outputPath, err)
	}

	handler := converter.S3Handler{
		Session: sess,
		Bucket:  S3_BUCKET,
	}

	err = handler.UploadFile("student_report.pdf", outputPath)
	if err != nil {
		log.Fatalf("UploadFile - filename: %v, err: %v", outputPath, err)
	}
	log.Println("UploadFile - success")
}

type Test struct {
	TestName   string
	TestDate   string
	Batch      string
	SA         int
	TMode      string
	Physics    int
	Chemistry  int
	Biology    int
	Total      int
	Percentage float64
	Percentile float64
	TestRank   string
	AIR        string
}
