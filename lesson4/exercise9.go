package main

import (
	"time"
)

type User struct {
	ID int
	Name string
	Email string
	Age int
}

type Report struct {
	Date string
	ReportID int
	User
}

func CreateReport(user User, reportDate string) Report {
	reportId := int(time.Now().UnixNano()) + user.ID
	return Report{User: user, ReportID: reportId, Date: reportDate}
}

func GenerateUserReports(users []User, reportDate string) []Report {
	reports := make([]Report, 0, len(users))
	for _, user := range users {
		report := CreateReport(user, reportDate)
		reports = append(reports, report)
	}

	return reports
}