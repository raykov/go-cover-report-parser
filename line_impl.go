package coverreportparser

type line struct {
	index     float64
	statement bool
	covered   bool
}

func (l *line) isCovered() bool {
	return l.covered
}

func (l *line) isStatement() bool {
	return l.statement
}
