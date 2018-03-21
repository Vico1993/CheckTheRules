package cleantherules

import (
	"fmt"
)

var ctr *CleanTheRules

func init() {
	// r *Rules --> rules.go
	r = NewRules()
	ctr = NewCleanTheRules()
}

// CleanTheRules will be the object, will contain all information we need,
// Like filename, filepath.
type CleanTheRules struct {
	filename string
	filepath string
	line     string
}

// NewCleanTheRules return an initialized CleanTheRules instances
func NewCleanTheRules() *CleanTheRules {
	ctr := new(CleanTheRules)

	// Default Parameter
	ctr.filename = ""
	ctr.filepath = ""
	ctr.line = ""

	return ctr
}

// Check will check every new modification, for the type we giving to it.
// @param t string Type of this check ( file, line, folder )
func Check(t string) {
	switch t {
	case "file":
		checkFile(ctr.filename)
	case "line":
		checkFile(ctr.line)
	default:
		fmt.Println("Default not set.")
	}
}

// checkFile, will take every line in the file to check, one by one.
// @param f string filename we have to check
func checkFile(f string) {
	// First Get file
	// If existe check every line in it
}

// checkLine will check a specifique line and return true if everything is good
// Or false if something happens
// @param l string line we have to check
// @return bool
func checkLine(l string) bool {
	// check each line
	// First, check all General Settings
	// Second, check specifique case.

	return true
}

// SetFileName explicitly defines filename.
// @params filename string set the filename to check
func SetFileName(filename string) { ctr.setFileName(filename) }
func (ctr *CleanTheRules) setFileName(filename string) {
	ctr.filename = filename
}

// SetFilePath explicitly defines filepath.
// @params filepath string set the filepath to check
func SetFilePath(filepath string) { ctr.setFilePath(filepath) }
func (ctr *CleanTheRules) setFilePath(filepath string) {
	ctr.filepath = filepath
}

// SetLineToCheck explicitly set the line to check
// @params line string set the line to check
func SetLineToCheck(line string) { ctr.setLine(line) }
func (ctr *CleanTheRules) setLine(line string) {
	ctr.line = line
}
