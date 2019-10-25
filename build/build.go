// Package build provides version information
package build

// Build is to compile passing -ldflags "-X main.Build <build sha1>"

var (
	// Time holds the string representation of the time when the binary was built.
	Time string
	// User holds the string name of the user who built the binary.
	User string
	// Version holds the software versione
	Version string
	// Build holds the git ref number
	Build string
	// Author is the software author
	Author = "Giuseppe Lo Brutto"
)
