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

func helper(dir string, withFile bool, currList *[]File, indent *int) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			*indent++
			*currList = append(*currList, File{Name: file.Name(), Indent: *indent})
			err := helper(path.Join(dir, file.Name()), withFile, currList, indent)
			if err != nil {
				return err
			}
		} else if withFile {
			*currList = append(*currList, File{Name: file.Name(), Indent: *indent})
		}
	}
	*indent--
	return nil
}

func getFileDirList(dir string, withFile bool) ([]File, error) {
	listDirsFiles := make([]File, 0)
	indent := 0
	err := helper(dir, withFile, &listDirsFiles, &indent)

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
