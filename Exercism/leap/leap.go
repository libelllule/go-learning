// Package leap determines whether a given year is a leap year or not.
package leap

// IsLeapYear returns true if the given year is a leap year, and false otherwise.
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && year%100 != 0)
}
