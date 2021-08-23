package main

import (
	"fmt"

	run "github.com/iwita/simple-sw/pkg/runtime"
)

func main() {
	fmt.Println("Execution starts here!")
	r := run.NewRuntime("/home/dgiagos/goprojects/demo-go-workflow/testfiles/test-workflow.yaml")
	r.Start()
}
