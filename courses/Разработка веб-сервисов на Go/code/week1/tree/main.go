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
	noFileSize    = -1
)

type File struct {
	Name     string
	Indent   int
	IsDir    bool
	FileSize int
}

func (f File) getFileSizeFormat() string {
	if f.FileSize == 0 {
		return "empty"
	}
	return fmt.Sprintf("(%d)b", f.FileSize)
}

func main() {
	dirTree(os.Stdout, "testdata", true)
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
		if dirFileList[i].Indent == 1 {
			fmt.Fprintf(out, "%s%s (%t)\n", indentSymbol, dirFileList[i].Name, isLastFileInCurrentLevel(dirFileList[i:]))
			continue
		}
		char := pickChar(dirFileList[i:])
		tab := createTabIndent(dirFileList[i].Indent - 1)
		if file && !dirFileList[i].IsDir {
			fileSize := dirFileList[i].getFileSizeFormat()
			fmt.Fprintf(out, "%s%s%s%s %s (%t)\n", boundSymbol, tab, char, dirFileList[i].Name, fileSize, isLastFileInCurrentLevel(dirFileList[i:]))
			continue
		}
		fmt.Fprintf(out, "%s%s%s%s (%t)\n", boundSymbol, tab, char, dirFileList[i].Name, isLastFileInCurrentLevel(dirFileList[i:]))
	}

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

func isLastFileInCurrentLevel(files []File) bool {
	if len(files) == 1 {
		return true
	}
	currLevel := files[0].Indent

	for i := 1; i < len(files); i++ {
		if files[i].Indent == currLevel {
			return false
		}
		if files[i].Indent < currLevel {
			return true
		}
	}
	return true
}

func pickChar(files []File) string {
	if isLastFileInCurrentLevel(files) {
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
			fileSize := int(fileInfo.Size())
			*currList = append(*currList, File{Name: file.Name(), Indent: *indent + 1, FileSize: fileSize})
		}
	}
	*indent--
	return nil
}
