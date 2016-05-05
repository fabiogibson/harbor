package core

import (
	"github.com/fatih/color"
	"os/exec"
	"path/filepath"
	"strings"
)

//
//
//
func RunTests(filename string) error {
	if !strings.HasSuffix(filename, ".go") {
		return nil
	}

	color.Blue("===> Running tests for %s", filename)

	pkg := filepath.Dir(filename)
	testOutput, err := exec.Command("go", "test", "./"+pkg).CombinedOutput()

	if err != nil {
		color.Red("--- Tests Failed")
		color.Red("%s", testOutput)
		color.Set(color.BgRed)
		color.White("Go back to work!")
	} else {
		color.Green("Test Result: %s", testOutput)
		color.Set(color.BgGreen)
		color.White("Good job! Keep going...")
	}

	return err
}
