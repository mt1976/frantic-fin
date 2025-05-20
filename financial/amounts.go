package financial

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// ParseFinancialAbbreviationToInt parses financial strings like "1b2m3k" into their numeric equivalent.
func ParseFinancialAbbreviationToInt(input string) (int64, error) {
	if input == "" {
		return 0, errors.New("input is empty")
	}

	// Normalize to lowercase
	input = strings.ToLower(input)

	// Match patterns like "1.5k", "2b", etc.
	re := regexp.MustCompile(`(\d+(\.\d+)?)([kmbt])`)
	matches := re.FindAllStringSubmatch(input, -1)

	// Build a string from matches to verify full coverage
	combinedMatches := ""
	for _, match := range matches {
		combinedMatches += match[0]
	}

	// If some input is not matched, it is malformed
	if combinedMatches != input {
		return 0, fmt.Errorf("invalid input format near: '%s'", strings.TrimPrefix(input, combinedMatches))
	}

	var total int64
	for _, match := range matches {
		numStr := match[1]
		suffix := match[3]

		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number '%s': %v", numStr, err)
		}

		switch suffix {
		case "k":
			total += int64(num * 1_000)
		case "m":
			total += int64(num * 1_000_000)
		case "b":
			total += int64(num * 1_000_000_000)
		case "t":
			total += int64(num * 1_000_000_000_000)
		default:
			return 0, fmt.Errorf("unsupported suffix '%s'", suffix)
		}
	}

	return total, nil
}

// ParseFinancialAbbreviationToFloat parses financial strings like "1.5k" into their numeric equivalent.
func ParseFinancialAbbreviationToFloat(input string) (float64, error) {
	intVal, err := ParseFinancialAbbreviationToInt(input)
	if err != nil {
		return 0, err
	}
	// Convert int64 to float64
	floatVal := float64(intVal)
	return floatVal, nil
}

func ParseFinancialAbbreviationToString(input string) (string, error) {
	intVal, err := ParseFinancialAbbreviationToInt(input)
	if err != nil {
		return "", err
	}
	// Convert int64 to string
	strVal := strconv.FormatInt(intVal, 10)
	return strVal, nil
}

func ParseFinancialAbbreviationToAmountString(input string) (string, error) {
	floatVal, err := ParseFinancialAbbreviationToFloat(input)
	if err != nil {
		return "", err
	}
	// Convert int64 to string
	strVal := FormatAmount(floatVal, "GBP")
	return strVal, nil
}
