// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package twofer

// Returns the string "One for ..., one for me"
// The dots being replaced by the name if not null else it is replaced by you
func ShareWith(name string) string {
	var placeholder string
	if name != "" {
		placeholder = name
	} else {
		placeholder = "you"
	}

	return "One for " + placeholder + ", one for me."
}
