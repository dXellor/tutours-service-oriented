package enum

type userRole int

const (
	Administrator userRole = iota
	Author
	Tourist
)
