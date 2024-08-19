package uniq

import (
	"bufio"
	"fmt"
	"io"
)

func Unique(in io.Reader, out io.Writer) error {
	input := bufio.NewScanner(in)
	var prev string

	for input.Scan() {
		txt := input.Text()

		if txt == prev {
			continue
		}

		if txt < prev {
			return fmt.Errorf("файл не отсортирован")
		}

		prev = txt
		fmt.Fprintln(out, txt)
	}
	return nil
}
