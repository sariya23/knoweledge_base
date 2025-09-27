package echo

import "strings"

func ConcatFor(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func ConcatBuilder(args []string, total int) string {
	sep := ""
	sb := strings.Builder{}
	sb.Grow(total)
	for _, v := range args {
		_, _ = sb.WriteString(sep + v)
		sep = " "
	}
	return sb.String()
}

func ConcatAndAllocBuilder(args []string) string {
	sep := ""
	sb := strings.Builder{}
	total := 0
	for i := 0; i < len(args); i++ {
		total += len(args[i])
	}
	sb.Grow(total)
	for _, v := range args {
		_, _ = sb.WriteString(sep + v)
		sep = " "
	}
	return sb.String()
}

func ConcatJoin(args []string) string {
	return strings.Join(args, " ")
}
