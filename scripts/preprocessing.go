package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Text struct {
	Text string `json:"text"`
}

func WriteCsvFromDest(src string, dest string) {
	fi, err := os.Open(src)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)

	des, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer des.Close()
	w := csv.NewWriter(des)

	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		text := Text{}
		json.Unmarshal(a, &text)
		s := []string{text.Text}

		w.Write(s)
		if err := w.Error(); err != nil {
			fmt.Printf("error writing csv:", err)
		}
	}

	w.Flush()
}

func WriteCsvFromSplit(src1 string, src2 string, dest string, tag string) {
	temp := [][]string{{}}

	f1, err := os.Open(src1)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer f1.Close()
	f2, err := os.Open(src2)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer f2.Close()

	br1 := bufio.NewReader(f1)
	br2 := bufio.NewReader(f2)

	des, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer des.Close()
	w := csv.NewWriter(des)
	for {
		a, _, c := br1.ReadLine()
		if c == io.EOF {
			break
		}
		var s []string
		if tag != "" {
			s = []string{tag, string(a)}
		} else {
			s = []string{string(a)}
		}
		temp = append(temp, s)
	}

	for {
		a, _, c := br2.ReadLine()
		if c == io.EOF {
			break
		}
		var s []string
		if tag != "" {
			s = []string{tag, string(a)}
		} else {
			s = []string{string(a)}
		}
		temp = append(temp, s)
	}

	for _, value := range temp {
		w.Write(value)
		if err := w.Error(); err != nil {
			fmt.Printf("error writing csv:", err)
		}
	}
	w.Flush()
}

func RawFileSplit() {
	WriteCsvFromDest("../data/raw/sensitive_posts.csv", "../data/split/sensitive_posts_split.csv")
	WriteCsvFromDest("../data/raw/sensitive_comments.csv", "../data/split/sensitive_comments_split.csv")
	WriteCsvFromDest("../data/raw/approved_posts.csv", "../data/split/approved_posts_split.csv")
	WriteCsvFromDest("../data/raw/approved_comments.csv", "../data/split/approved_comments_split.csv")
	WriteCsvFromDest("../data/raw/deleted_posts.csv", "../data/split/deleted_posts_split.csv")
	WriteCsvFromDest("../data/raw/deleted_comments.csv", "../data/split/deleted_comments_split.csv")
}

func WriteCsvForPreprocessing() {
	RawFileSplit()
	WriteCsvFromSplit("../data/split/approved_posts_split.csv", "../data/split/approved_comments_split.csv", "../data/split/sensitive_approved.csv", "1")
	WriteCsvFromSplit("../data/split/deleted_posts_split.csv", "../data/split/deleted_comments_split.csv", "../data/split/sensitive_deleted.csv", "0")
	WriteCsvFromSplit("../data/split/sensitive_approved.csv", "../data/split/sensitive_deleted.csv", "../data/pre/training.csv", "")
}

func main() {
	WriteCsvForPreprocessing()
}
