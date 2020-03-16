// Package errutils contains error utilities
package errutils

// Must handles errors
func Must(err error) {
	if err != nil {
		panic(err)
	}
}
