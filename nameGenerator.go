package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	naming(2)
    naming(2)
	naming(2)
}
func naming(syllables int) {
	name := ""
	for i := 1; i <= syllables; i++ {
		name = name + genSyl("")
	}
	fmt.Println(strings.ToUpper(name))
}
func genSyl(nameChunk string) string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	n, err := os.Open("names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer n.Close()
	scanner := bufio.NewScanner(n)
	j := 0
	var names [1150]string
	for scanner.Scan() {
		names[j] = scanner.Text()
		j++
	}

	fIndex := r.Intn(150)
	fHalf := names[fIndex]
	fWLen := len(fHalf)
	for k, c := range fHalf {
		if k <= fWLen/2 {
			nameChunk = nameChunk + string(c)
		}
    }
	return nameChunk
}
