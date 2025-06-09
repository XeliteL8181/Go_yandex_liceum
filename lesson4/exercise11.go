package main

import (
	"strings"
)

type Writer interface {
	Write(p []byte) int
}

type Reader interface {
	Read() []byte
}

type ReaderWriter interface {
	Reader
	Writer
}

type UpperReaderWriter struct {
	UpperString string
}

func (u *UpperReaderWriter) Write(testData []byte) int {
	u.UpperString = strings.ToUpper(string(testData))
	return len(u.UpperString)
}

func (u *UpperReaderWriter) Read() []byte {
	return []byte(u.UpperString)
}