package main

import (
	"fmt"
	"os"
	"path"
)

// var listDirsFiles = make([]string, 0)

const (
	indentSymbol  = "├───"
	boundSymbol   = "│"
	indentValue   = "\t"
	subFileSymbol = "└───"
)

type File struct {
	Name   string
	Indent int
}

// dirTree выводит список каталогов в указанной директории.
// Если значение file равно true, то тогда выведутся и
// файлы в каталогах.
// func dirTree(out io.Writer, dirName string, file bool) error {
// 	dirFileList, err := getFileDirList(dirName, file)

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func helper(dir string, withFile bool, currList *[]string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			*currList = append(*currList, file.Name())
			err := helper(path.Join(dir, file.Name()), withFile, currList)
			if err != nil {
				return err
			}
		} else if withFile {
			*currList = append(*currList, file.Name())
		}
	}
	return nil
}

func getFileDirList(dir string, withFile bool) ([]string, error) {
	listDirsFiles := make([]string, 0)

	err := helper(dir, withFile, &listDirsFiles)

	if err != nil {
		return nil, err
	}

	return listDirsFiles, nil
}

func main() {
	list, err := getFileDirList("testdata", false)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}
