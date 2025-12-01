package main

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"text/template"

	"alexi.ch/aoc/2025/lib"
)

const TEMPLATE_DIR = "templates"
const PROBLEMS_FOLDER = "problems"

type TemplateData struct {
	ProblemNumber    int
	ProblemNumberStr string
	PackageName      string
	Title            string
	Cwd              string
}

func main() {
	if len(os.Args) < 2 {
		panic("Please provide at least 1 problem number")
	}

	nr := os.Args[1]
	title := ""
	if len(os.Args) >= 3 {
		title = strings.Join(os.Args[2:], " ")
	}
	title = strings.ToTitle(title)

	cwd, err := os.Getwd()
	lib.Check(err)

	// create files for each given nr:
	problemNr, err := strconv.Atoi(nr)

	lib.Check(err)
	tplData := TemplateData{
		ProblemNumber:    problemNr,
		ProblemNumberStr: fmt.Sprintf("%02d", problemNr),
		PackageName:      fmt.Sprintf("day%02d", problemNr),
		Title:            title,
		Cwd:              cwd,
	}

	// create problem files:
	err = createProblemFile(tplData)
	lib.Check(err)

	// create entry in main (aoc.go):
	processMainEntrypoint(tplData)

	// create empty data files:
	createDataFiles(tplData)
}

func createProblemFile(tplData TemplateData) error {
	// Create template:
	tplDir := path.Join(tplData.Cwd, TEMPLATE_DIR)
	tplFile := path.Join(tplDir, "day.go.tpl")
	tplName := path.Base(tplFile)
	tpl := template.Must(template.New(tplName).
		Funcs(template.FuncMap{
			"format": func(format string, value any) string {
				return fmt.Sprintf(format, value)
			},
		}).
		ParseFiles(tplFile))
	// Create problem file:
	cwd := tplData.Cwd
	problemFolder := path.Join(cwd, PROBLEMS_FOLDER, tplData.PackageName)
	fmt.Printf("Creating new problem from template: %s in folder %s\n", tplData.ProblemNumberStr, problemFolder)

	err := os.MkdirAll(problemFolder, 0750)
	outputPath := path.Join(problemFolder, fmt.Sprintf("day%02d.go", tplData.ProblemNumber))
	lib.Check(err)

	if lib.FileExists(outputPath) {
		fmt.Printf("File already exists: %s\n", outputPath)
		return fmt.Errorf("file already exists: %s", outputPath)
	}
	fh, err := os.Create(outputPath)
	lib.Check(err)
	defer fh.Close()

	err = tpl.Execute(fh, tplData)
	lib.Check(err)

	return nil
}

func processMainEntrypoint(tplData TemplateData) {
	mainName := path.Join(tplData.Cwd, "aoc.go")
	main, err := os.ReadFile(mainName)
	lib.Check(err)
	mainStr := string(main)

	fmt.Printf("Replacing template comments in %s\n", mainName)

	// find all occurences of "//template:*", and process them:
	re := regexp.MustCompile(`//template:(.*)`)
	matches := re.FindAllStringSubmatch(mainStr, -1)
	for _, match := range matches {
		tplStr := match[1]
		tpl := template.Must(template.New("tpl").
			Funcs(template.FuncMap{
				"format": func(format string, value any) string {
					return fmt.Sprintf(format, value)
				},
			}).
			Parse(tplStr))

		var buffer bytes.Buffer
		err = tpl.ExecuteTemplate(&buffer, "tpl", tplData)
		lib.Check(err)
		output := match[0] + "\n" + buffer.String()

		mainStr = strings.ReplaceAll(mainStr, match[0], output)
	}

	err = os.WriteFile(mainName, []byte(mainStr), 0755)
	lib.Check(err)
}

func createDataFiles(tplData TemplateData) {
	dataDir := path.Join(tplData.Cwd, "data")
	fmt.Printf("Creating data files in %s\n", dataDir)
	dataFile := fmt.Sprintf("%02d-data.txt", tplData.ProblemNumber)
	testDataFile := fmt.Sprintf("%02d-test-data.txt", tplData.ProblemNumber)

	dfh, err := os.Create(path.Join(dataDir, dataFile))
	lib.Check(err)
	defer dfh.Close()

	tfh, err := os.Create(path.Join(dataDir, testDataFile))
	lib.Check(err)
	defer tfh.Close()
}
