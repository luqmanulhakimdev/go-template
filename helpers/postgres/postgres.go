package postgres

import (
	"fmt"
	"strconv"
	"strings"
)

func ReplaceSQL(stmt, pattern string, len int) string {
	pattern += ","
	stmt = fmt.Sprintf(stmt, strings.Repeat(pattern, len))
	n := 0
	for strings.IndexByte(stmt, '?') != -1 {
		n++
		param := "$" + strconv.Itoa(n)
		stmt = strings.Replace(stmt, "?", param, 1)
	}
	return strings.TrimSuffix(stmt, ",")
}

func SubstitutePlaceholder(data string, startInt int) (res string) {
	placeholderCount := strings.Count(data, "?")
	res = data
	for i := startInt; i < startInt+placeholderCount; i++ {
		res = strings.Replace(res, "?", "$"+strconv.Itoa(i), 1)
	}
	return res
}

func BuildStrParams(len int, separator string) string {
	var res string
	for i := 0; i < len; i++ {
		if i != 0 {
			res += separator
		}
		res += "?"
	}
	return res
}

func AppendStrArgs(args []string, currArgs *[]interface{}) {
	for _, v := range args {
		*currArgs = append(*currArgs, v)
	}
}

func AppendIntArgs(args []int, currArgs *[]interface{}) {
	for _, v := range args {
		*currArgs = append(*currArgs, v)
	}
}
