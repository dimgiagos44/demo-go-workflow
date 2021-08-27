package main

import (
	"fmt"

	run "github.com/dimgiagos44/demo-go-workflow/pkg/runtime"
)

func main() {
	fmt.Println("Runtime execution starts here!-----------------------")
	r := run.NewRuntime("./testfiles/test-workflow.yaml", run.WithInputFile("./testfiles/applicant.json"))
	r.Start()
	fmt.Println("Runtime execution ends here!-------------------------")
}
