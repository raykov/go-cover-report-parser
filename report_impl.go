package coverreportparser

import (
	"fmt"
	"sort"
)

type report struct {
	files map[string]File
}

func (report *report) Add(fileName string, line Line) {
	if report.files == nil {
		report.files = make(map[string]File)
	}

	f, ok := report.files[fileName]
	if !ok {
		report.files[fileName] = &file{}
		f = report.files[fileName]
	}

	f.Add(line)
}

func (report *report) Coverage() float64 {
	var countLines, coveredLines float64

	for _, file := range report.files {
		count, covered := file.Info()

		countLines += count
		coveredLines += covered
	}

	return coveredLines / countLines * 100
}

func (report *report) String() string {
	fileNames := make([]string, 0, len(report.files))
	for fileName := range report.files {
		fileNames = append(fileNames, fileName)
	}

	sort.Slice(fileNames, func(i, j int) bool { return fileNames[i] < fileNames[j] })

	str := ""
	for _, fileName := range fileNames {
		str += fmt.Sprintf("%.2f%%: \t%s\n", report.files[fileName].Coverage(), fileName)
	}
	return str
}
