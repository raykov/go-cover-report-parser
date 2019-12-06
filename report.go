package coverreportparser

// Report is an interface represents a coverage report
type Report interface {
	Add(string, Line)
	Coverage() float64
	String() string
}
