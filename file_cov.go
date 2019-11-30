package coverreportparser

type fileCov struct {
	lines   float64
	covered float64
}

func (fc *fileCov) Add(lines, covered float64) {
	fc.lines += lines
	fc.covered += covered
}

func (fc *fileCov) Coverage() float64 {
	return fc.covered / fc.lines * 100
}
