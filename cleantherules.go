package cleantherules

var r *Rules

// Rules is the object will conaint all the rules the user want to check
type Rules struct {
}

// New return an initialized Rules instances
func New() *Rules {
	r := new(Rules)

	return r
}
