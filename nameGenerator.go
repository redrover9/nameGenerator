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
	fName, lName := genSyl("", "")
    name = fName + lName
	name = strings.ToLower(name)
	fmt.Println(strings.Title(name))
}
func genSyl(nameBeginning, nameEnding string) (string, string) {
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
	lIndex := r.Intn(150)
	fHalf := names[fIndex]
	lHalf := names[lIndex]
	fWLen := len(fHalf)
	lWLen := len(lHalf)
	for k, c := range fHalf {
		if k <= fWLen/2 {
			nameBeginning = nameBeginning + string(c)
		}
	}
	for l, c := range lHalf {
		if l >= lWLen/2 {
			nameEnding = nameEnding + string(c)
		}
	}
    return nameBeginning, nameEnding
}
