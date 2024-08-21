package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

const (
	indentSymbol  = "├───"
	boundSymbol   = "│"
	indentValue   = "\t"
	subFileSymbol = "└───"
	noFileSize    = ""
)

type File struct {
	Name     string
	Indent   int
	IsDir    bool
	FileSize string
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
			continue
		}
		char := pickChar(dirFileList[i], dirFileList[i+1])
		tab := createTabIndent(dirFileList[i].Indent - 1)
		if file && !dirFileList[i].IsDir {
			fileSize := getFileSizeFormat(dirFileList[i])
			fmt.Fprintf(out, "%s%s%s%s %s\n", boundSymbol, tab, char, dirFileList[i].Name, fileSize)
			continue
		}
		fmt.Fprintf(out, "%s%s%s%s\n", boundSymbol, tab, char, dirFileList[i].Name)
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

func getFileSizeFormat(file File) string {
	if file.FileSize == "0" {
		return "empty"
	}
	return fmt.Sprintf("(%s)b", file.FileSize)
}

func pickChar(curr, next File) string {
	if curr.IsDir && next.Indent > curr.Indent {
		return indentSymbol
	}
	if next.Indent < curr.Indent || next.Indent > curr.Indent {
		return subFileSymbol
	}
	return indentSymbol
}

func createTabIndent(indent int) string {
	if indent >= 2 {
		tabs := strings.Repeat("\t", indent)
		return strings.Join(strings.Split(tabs, ""), "|")
	}
	return strings.Repeat("\t", indent)
}

func helper(dir string, withFile bool, currList *[]File, indent *int) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			*indent++
			*currList = append(*currList, File{Name: file.Name(), Indent: *indent, IsDir: true, FileSize: noFileSize})
			err := helper(path.Join(dir, file.Name()), withFile, currList, indent)
			if err != nil {
				return err
			}
		} else if withFile {
			fileInfo, err := os.Stat(path.Join(dir, file.Name()))
			if err != nil {
				return err
			}
			fileSize := strconv.Itoa(int(fileInfo.Size()))
			*currList = append(*currList, File{Name: file.Name(), Indent: *indent + 1, FileSize: fileSize})
		}
	}
	*indent--
	return nil
}

func main() {
	_, err := getFileDirList("testdata", true)
	if err != nil {
		fmt.Println(err)
	}

	dirTree(os.Stdout, "testdata", true)
}
