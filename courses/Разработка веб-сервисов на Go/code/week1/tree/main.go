package main

import (
	"fmt"
	"io"
	"os"
)

// dirTree выводит список каталогов в указанной директории.
// Если значение file равно true, то тогда выведутся и
// файлы в каталогах.
func dirTree(out io.Writer, dirName string, file bool) error {
	return nil
}

func main() {
	// out := os.Stdout
	// err := dirTree(out, "testdata", true)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	files, _ := os.ReadDir("testdata")
	fmt.Println(files)

	for _, f := range files {
		if f.IsDir() {
			fmt.Println(f.Name(), f)
			files2, err := os.ReadDir(f.Name())
			if err != nil {
				fmt.Println("got error: (%v)", err)
			}
			fmt.Println(files2)
			for _, f1 := range files2 {
				fmt.Println(f1.Name())
			}
		}
	}
}
