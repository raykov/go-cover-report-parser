package coverreportparser

// Result is an interface represents a result of coverage check
type Result interface {
	String() string
	IsFailure() bool
}
