package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	gopath "path"
	"strconv"
	"strings"
)

func enumerateDirs(path string) []os.FileInfo {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	dirs := []os.FileInfo{}

	for _, f := range files {
		if f.Mode().IsDir() {
			dirs = append(dirs, f)
		}
	}

	return dirs
}

func enumerateFiles(path string) []string {
	items, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var files []string

	for _, f := range items {
		if !f.Mode().IsDir() {
			files = append(files, f.Name())
		}
	}

	return files
}

func printDirs(dirs []os.FileInfo, chosen []bool, showBrackets bool) {
	for i, dir := range dirs {
		char := ""
		if showBrackets {
			if chosen[i] {
				char = "[X] "
			} else {
				char = "[ ] "
			}
		}
		fmt.Printf("%d. %s%s\n", i, char, dir.Name())
	}
}

func chooseDirs(dirs []os.FileInfo) ([]string, string) {
	chosen := make([]bool, len(dirs))

	reader := bufio.NewReader(os.Stdin)
	input := "DUMMY"

	for len(input) > 0 {
		fmt.Print("\n\n\n")
		printDirs(dirs, chosen, true)
		fmt.Print("--------------------\nChoose INPUT Dirs (seprated by spaces, nothing to continue): ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) == 0 {
			continue
		}

		selections := strings.Fields(input)
		for _, sel := range selections {
			index, _ := strconv.Atoi(sel)

			chosen[index] = !chosen[index]
		}
	}

	var inputDirs []string
	for i, dir := range dirs {
		if chosen[i] {
			inputDirs = append(inputDirs, dir.Name())
		}
	}

	fmt.Print("\n\n\n")
	printDirs(dirs, nil, false)
	fmt.Print("--------------------\nChoose OUTPUT Dir: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, _ := strconv.Atoi(input)

	outputDir := dirs[index].Name()

	return inputDirs, outputDir
}

type action struct {
	from string
	to   string
}

func dryRun(inputDirs []string, outputDir string) []action {

	actions := []action{}

	for i, indir := range inputDirs {
		files := enumerateFiles(indir)
		for _, file := range files {
			newFileName := fmt.Sprintf("D%d.%s", i, file)
			fromPath := gopath.Join(indir, file)
			toPath := gopath.Join(outputDir, newFileName)

			actions = append(actions, action{from: fromPath, to: toPath})
		}
	}

	return actions
}

func check(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	dirs := enumerateDirs("./")
	inputDirs, outputDir := chooseDirs(dirs)

	actions := dryRun(inputDirs, outputDir)

	fmt.Printf("\nThe following renames will occur:\n")
	for _, action := range actions {
		fmt.Printf("[%s] -> [%s]\n", action.from, action.to)
	}

	fmt.Printf("\n %d total moves\n", len(actions))

	fmt.Printf("\nEnter to continue, anything else cancels rename [Enter]: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if len(input) > 0 {
		fmt.Println("Rename canceled!")
		os.Exit(0)
	}

	for _, action := range actions {
		fmt.Printf("Moving [%s] to [%s]\n", action.from, action.to)
		if _, err := os.Stat(action.to); !os.IsNotExist(err) {
			// Destination File Exists
			fmt.Printf("ERROR: %s already exists, aborting\n", action.to)
			os.Exit(1)
		}
		err := os.Rename(action.from, action.to)
		check(err)
	}

	fmt.Printf("\nRemove INPUT directories? Enter = Yes, anything else will exit [Enter]: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if len(input) > 0 {
		fmt.Println("Removal of INPUT directories canceled!")
		os.Exit(0)
	}

	for _, dir := range inputDirs {
		fmt.Printf("Removing: %s\n", dir)
		err := os.Remove(dir)
		check(err)
	}

}
