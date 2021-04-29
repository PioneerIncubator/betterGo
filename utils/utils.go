package utils

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
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
			log.Fatal(err)
		}
		increased := i + first
		return test[0] + separator + strconv.Itoa(increased)
	}

	return str + separator + strconv.Itoa(first)
}
