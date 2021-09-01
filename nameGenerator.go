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
	naming(1)
	naming(2)
	naming(3)
	naming(4)
}
func naming(syllables int) {
	//rVIndex := r.Intn(6)
	//rCIndex := r.Intn(20)
	var v [6]string
	v[0] = "a"
	v[1] = "e"
	v[2] = "i"
	v[3] = "o"
	v[4] = "u"
	v[5] = "y"
	var con [20]string
	con[0] = "b"
	con[1] = "c"
	con[2] = "d"
	con[3] = "f"
	con[4] = "g"
	con[5] = "h"
	con[6] = "j"
	con[7] = "k"
	con[8] = "l"
	con[9] = "m"
	con[10] = "n"
	con[11] = "p"
	con[12] = "q"
	con[13] = "r"
	con[14] = "s"
	con[15] = "t"
	con[16] = "v"
	con[17] = "w"
	con[18] = "x"
	con[19] = "z"
	syl := ""
	for i := 1; i <= syllables; i++ {
		syl = syl + genSyl("")
	}
	fmt.Println(strings.ToUpper(syl))
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
	var names [150]string
	for scanner.Scan() {
		names[j] = scanner.Text()
		j++
	}

	fIndex := r.Intn(150)
	fHalf := names[fIndex]
	fWLen := len(fHalf)
	for k, c := range fHalf {
		if k < fWLen/2 {
			nameChunk = nameChunk + string(c)
		}
    }
	return nameChunk
}
