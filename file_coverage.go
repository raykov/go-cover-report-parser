package coverreportparser

type fileCoverage struct {
	lines   float64
	covered float64
}

func (fc *fileCoverage) Add(lines, covered float64) {
	fc.lines += lines
	fc.covered += covered
}

func (fc *fileCoverage) Coverage() float64 {
	return fc.covered / fc.lines * 100
}
