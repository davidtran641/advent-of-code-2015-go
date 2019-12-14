package main

import (
	"crypto/md5"
	"strconv"
	"encoding/hex"
	"io"
	"strings"
	"fmt"
)

func makeString(pattern string, repeat int) string {
	result := ""
	for i := 0; i < repeat; i++ {
		result += pattern
	}

	return result
}

func FindSuitableHash(start string, numZero int) int {
	prefix := makeString("0", numZero)
	num := 1
	
	for {
		key := start + strconv.Itoa(num)

		hasher := md5.New()
		io.WriteString(hasher, key)
		hash := hasher.Sum(nil)
		value := hex.EncodeToString(hash)

		if strings.HasPrefix(value, prefix) {
			return num
		}

		num += 1
		
	}
}

func main() {
	fmt.Println(FindSuitableHash("abcdef", 5))
}