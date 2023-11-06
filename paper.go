package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func split_txt(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var filelines []string

	for fileScanner.Scan() {
		filelines = append(filelines, fileScanner.Text())
	}
	file.Close()
	return filelines
}

func firstN(str string, n int) string {
	v := []rune(str)
	if n >= len(v) {
		return str
	}
	return string(v[:n])
}

func main() {

	blankFile := split_txt("blank-sheet-of-paper.txt")

	shreddedFile := split_txt("shredded-sheet-of-paper.txt")

	

	var sorted_File []string
	for lineNr := range blankFile {
		blankline := blankFile[lineNr]
		beggining := firstN(blankline, 20)

		for line := range shreddedFile {
			fullLine := shreddedFile[line]
			first := firstN(fullLine, 20)
			if first == beggining {
				sorted_File = append(sorted_File, fullLine)
			}
		}

	}

	file, _ := os.OpenFile("Repaired-File.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)

	datawriter := bufio.NewWriter(file)

	for _, data := range sorted_File {
		_, _ = datawriter.WriteString(data + "\n")

	}
	datawriter.Flush()
	file.Close()
	fmt.Println("Finsished")

}
