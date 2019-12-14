package main

import (
	"strings"
	"strconv"
	"math"
)

func ParseDimension(dimension string) []int {
	list := strings.Split(dimension, "x")
	
	result := make([]int, len(list))
	for i, s := range list {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		result[i] = v
	}
	return result
}

// 29x13x26
func FindWrappingArea(dimension string) int {
	list := ParseDimension(dimension)
	if len(list) != 3 {
		panic("Wrong length of list dimensions")
	}

	result := 0
	minArea := math.MaxInt32
	for i, a := range list {
		for j := i + 1; j < len(list); j += 1 {
			b := list[j]

			area := a * b 
			result += 2 * area

			if minArea > area { minArea = area }
		}
	}

	result += minArea
	return result
}

func FindWrappingAreaList(dims []string) int {
	result := 0
	for _, dim := range dims {
		result += FindWrappingArea(dim)
	}
	return result
}

func FindRobbinLength(dim string) int {
	list := ParseDimension(dim)
	if len(list) != 3 {
		panic("Wrong length of list dimensions")
	}

	minSum := math.MaxInt32
	for i, a := range list {
		for j := i + 1; j < len(list); j += 1 {
			b := list[j]

			sum := a + b 

			if minSum > sum { minSum = sum }
		}
	}

	result := minSum * 2 + list[0] * list[1] * list[2]
	return result
}

func FindRobbinLengthList(dims []string) int {
	result := 0
	for _, dim := range dims {
		result += FindRobbinLength(dim)
	}
	return result
}