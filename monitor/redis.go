package monitor

import (
	"fmt"
	"strings"
)

func RdsInfo(data []string) {
	s := " {"

	for _, value := range data {
		_t := strings.Split(value, "\r")
		_s := strings.TrimSpace(_t[0])

		if len(_s) > 1 {

			s += " \"" + _s + "\":{"
			for i, _v := range _t[1:] {
				org := strings.Split(_v, ":")
				if len(org) > 1 {
					if len(_t[1:]) == (i + 1) {
						s += "\"" + strings.Trim(org[0], "\n") + "\":\"" + org[1] + "\""
					} else {
						s += "\"" + strings.Trim(org[0], "\n") + "\":\"" + org[1] + "\","
					}
				}

			}
			s += "},"
		}

	}
	fmt.Println(s + "}")

}
