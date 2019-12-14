package main

import (
	"io/ioutil"
	"strings"
)

func ReadFile(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func ReadLines(path string) []string {
	s := ReadFile(path)
	arr := strings.Split(s, "\n")
	return arr
}