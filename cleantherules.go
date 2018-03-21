package cleantherules

func init() {
	// r *Rules --> rules.go
	r = New()
}

// Check will check every new modification, for the type we giving to it.
// @param t string Type of this check ( file, line, folder )
func Check(t string) {

}

// checkFile, will take every line in the file to check, one by one.
// @param f string filename we have to check
func checkFile(f string) {

}

// checkLine will check a specifique line and return true if everything is good
// Or false if something happens
// @param l string line we have to check
// @return bool
func checkLine(l string) bool {
	return true
}
