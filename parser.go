package main

import (
	"errors"
	"strconv"
	"strings"
)

func parsePorts(input string) ([]string, error) {
	ports := []string{}

	if input == "" {
		return ports, nil
	}

	input = strings.ReplaceAll(input, " ", "")

	parts := strings.Split(input, ",")

	for _, part := range parts {
		if strings.ContainsAny(part, "-") {
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) != 2 {
				return ports, errors.New("invalid range")
			}

			start, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				return ports, errors.New("invalid range")
			}

			end, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				return ports, errors.New("invalid range")
			}

			if start > end {
				return ports, errors.New("invalid range")
			}

			for i := start; i <= end; i++ {
				ports = append(ports, strconv.Itoa(i))
			}
		} else {
			ports = append(ports, part)
		}
	}

	return ports, nil
}
