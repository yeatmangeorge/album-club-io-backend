/*
	Package for converting uints.
*/
package util

import "strconv"

// UintToString takes a base 10 uint value and returns a human readable string.
func UintToString(value uint) string {
	return strconv.FormatUint(uint64(value), 10)
}

// IntToString takes a base 10 int value and returns a human readable string.
func IntToString(value int) string {
	return strconv.FormatInt(int64(value), 10)
}
