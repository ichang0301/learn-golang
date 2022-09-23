package search

// SearchAlgorithm is an interface to sort a list.
// Supported data type is 'int', 'float32', 'float64', 'string'.
type SearchAlgorithm interface {
	Search(interface{}) (int, error)
}
