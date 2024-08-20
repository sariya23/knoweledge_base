package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

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
func dirTree(out io.Writer, dirName string, file bool) error {
	dirFileList, err := getFileDirList(dirName, file)

	if err != nil {
		return err
	}

	for i := 0; i < len(dirFileList); i++ {
		fmt.Fprintf(out, "*%s%s\n", strings.Repeat("-", dirFileList[i].Indent), dirFileList[i].Name)
	}
	return nil
}

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
			*currList = append(*currList, File{Name: file.Name(), Indent: *indent + 1})
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
	l, err := getFileDirList("testdata", true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n%v\n", l)

	dirTree(os.Stdout, "testdata", true)
}
