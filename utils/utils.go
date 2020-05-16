package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IncrementString(str string, separator string, first int) string {
	if separator == "" {
		separator = "_"
	}

	if first == 0 || first < 0 {
		first = 1
	}

	test := strings.SplitN(str, separator, 2)
	expect := 2
	if len(test) >= expect {
		i, err := strconv.Atoi(test[1])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		increased := i + first
		return test[0] + separator + strconv.Itoa(increased)
	}

	return str + separator + strconv.Itoa(first)
}
