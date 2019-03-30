package service

import (
	"strings"
)

// SortDirection represents whether we are sorting Ascending or Descending
type SortDirection int

const (
	// SortAscending means to sort in Ascending order
	SortAscending SortDirection = iota
	// SortDescending means to sort in Descending order
	SortDescending
	// SortNatural means to sort in whatever direction is natural for the field
	SortNatural
)

// SortField represents a single field to sort against
type SortField struct {
	Field     string
	Direction SortDirection
}

// ParseSorts will parse a string representing the fields to sort and produce a slice of SortField structs
// The input string should be comma separated, where each entry is a field name that is optionally prefixed with a + or -
// A + prefix indicates to sort in Ascending Order
// A - prefix indicates to sort in Descending Order
// No prefix indicates to sort in Natural Order for the field - typically Ascending but not always
func ParseSorts(input string) []SortField {
	result := []SortField{}
	for _, field := range strings.Split(input, ",") {
		trimmed := strings.TrimSpace(field)
		if trimmed != "" {
			switch trimmed[0] {
			case '+':
				result = append(result, SortField{Field: trimmed[1:], Direction: SortAscending})
			case '-':
				result = append(result, SortField{Field: trimmed[1:], Direction: SortDescending})
			default:
				result = append(result, SortField{Field: trimmed, Direction: SortNatural})
			}
		}
	}
	return result
}
