package coverreportparser

// Line is an interface represents a covered or not line
type Line interface {
	isCovered() bool
	isStatement() bool
}
