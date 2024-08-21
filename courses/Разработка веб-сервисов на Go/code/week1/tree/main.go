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
	isDir  bool
}

// dirTree выводит список каталогов в указанной директории.
// Если значение file равно true, то тогда выведутся и
// файлы в каталогах.
func dirTree(out io.Writer, dirName string, file bool) error {
	dirFileList, err := getFileDirList(dirName, file)

	if err != nil {
		return err
	}

	for i := 0; i < len(dirFileList)-1; i++ {
		if dirFileList[i].Indent == 1 {
			fmt.Fprintf(out, "%s%s\n", indentSymbol, dirFileList[i].Name)
		} else if (dirFileList[i+1].Indent > dirFileList[i].Indent || dirFileList[i+1].Indent < dirFileList[i].Indent) && !dirFileList[i].isDir {
			fmt.Fprintf(out, "%s%s%s%s\n", boundSymbol, strings.Repeat("\t", dirFileList[i].Indent-1), subFileSymbol, dirFileList[i].Name)
		} else {
			fmt.Fprintf(out, "%s%s%s%s\n", boundSymbol, strings.Repeat("\t", dirFileList[i].Indent-1), indentSymbol, dirFileList[i].Name)
		}
	}

	fmt.Fprintf(out, "%s%s", subFileSymbol, dirFileList[len(dirFileList)-1].Name)
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

func helper(dir string, withFile bool, currList *[]File, indent *int) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			*indent++
			*currList = append(*currList, File{Name: file.Name(), Indent: *indent, isDir: true})
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

func main() {
	l, err := getFileDirList("testdata", true)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\n%v\n", l)

	dirTree(os.Stdout, "testdata", true)
}
