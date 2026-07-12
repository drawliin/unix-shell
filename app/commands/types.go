package commands

type shellToken struct {
	text   string
	quoted bool
}

type stdPath struct {
	path string
	stdout bool
	stderr bool
	appendEOF bool
}
