package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/huichen/sego"
	"io"
	"math"
	"os"
	"strings"
)

var Segmenter sego.Segmenter
var StopWords map[string]bool
var N SumNum
var Terms map[string]Term

type Term struct {
	A   int32   // c1 contains `x` sum
	B   int32   // c0 contains `x` sum
	C   int32   // c1 not contains `x` sum
	D   int32   // c0 not contains `x` sum
	CHI float64 // CHI
}

type SumNum struct {
	X int32 // c1
	Y int32 // c0
}

func loadDict() {
	dictSrc := "../data/dictionary.txt,../data/sensitive.txt"
	Segmenter.LoadDictionary(dictSrc)
}

func main() {
	loadStopWords()
	loadDict()
	loadData()
	caculateEigenvalue()
	fmt.Println(Terms)
}

func isStopWord(s string) bool {
	if StopWords[s] {
		return true
	} else {
		return false
	}
}

func setOrUpdateTerms(c string, text string) {
	// check if A ++ OR B ++
	term, _ := Terms[text]
	switch c {
	case "1":
		term.A += 1
		break
	case "0":
		term.B += 1
		break
	}
	Terms[text] = term
}

func chineseSegment(c string, s string) {
	segments := Segmenter.Segment([]byte(s))
	if len(segments) > 0 {
		for _, seg := range segments {
			token := seg.Token()
			var text = token.Text()
			if !isStopWord(s) && len(strings.TrimSpace(s)) != 0 {
				setOrUpdateTerms(c, text)
			}
		}
	}
}

func caculateEigenvalue() {
	//CHI
	for key, term := range Terms {
		term.C = N.X - term.A
		term.D = N.Y - term.B
		n := N.X + N.Y
		a := term.A
		b := term.B
		c := term.C
		d := term.D
		term.CHI = float64(n) * math.Pow(float64(a*d-b*c), 2) / float64((a+c)*(b+d)*(a+b)*(c+d))
		Terms[key] = term
	}
}

func loadStopWords() {
	StopWords = make(map[string]bool)

	fi, err := os.Open("../data/stop_words.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)

	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		StopWords[string(a)] = true
	}
}

func loadData() {
	Terms = make(map[string]Term)

	fi, err := os.Open("../data/split/sensitive_approved_split.csv")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	reader := csv.NewReader(fi)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if record[0] == "0" {
			N.X += 1
		} else {
			N.Y += 1
		}
		chineseSegment(record[0], record[1])
	}
}
