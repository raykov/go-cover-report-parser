package coverreportparser

// File is an interface represents a covered file
type File interface {
	Add(Line)
	Coverage() float64
	Info() (float64, float64)
}
