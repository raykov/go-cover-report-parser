package coverreportparser

type file struct {
	lines []Line
}

func (f *file) Add(line Line) {
	f.lines = append(f.lines, line)
}

func (f *file) Coverage() float64 {
	count, covered := f.Info()

	return covered / count * 100
}

func (f *file) Info() (count, covered float64) {
	for _, line := range f.lines {
		if !line.isStatement() {
			continue
		}

		count++
		if line.isCovered() {
			covered++
		}
	}

	return count, covered
}
