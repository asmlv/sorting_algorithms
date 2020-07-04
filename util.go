package algorithms

// Order for sorting
type Order = func(int, int) bool

// Less is a predicate for ascending order
func Less(a, b int) bool {
	return a < b
}

// Greater is a predicate for descending order
func Greater(a, b int) bool {
	return a > b
}
