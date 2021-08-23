package main

import (
	"fmt"
	"github.com/serverlessworkflow/sdk-go/model"
	"github.com/serverlessworkflow/sdk-go/parser"
)

func ParseWorkflow(filePath string) (*model.Workflow, error){
	workflow, err := parser.FromFile(filePath)
	if err != nil {
		return nil, err
	}
	return workflow, nil
}

func main(){
	workflow, err := ParseWorkflow("/home/dgiagos/goprojects/demo-go-workflow/testfiles/test-workflow.yaml")
	fmt.Println("Workflow states [0] = ", workflow.States[0])
	fmt.Println("Workflow states [1]  = ", workflow.States[1])
	fmt.Println("Error = ", err)
	a := workflow.States[0]
	b := workflow.States[1]
	fmt.Println(a.GetName(), b.GetName())

}
