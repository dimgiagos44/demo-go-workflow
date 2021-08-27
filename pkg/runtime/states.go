package runtime

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/itchyny/gojq"
	"github.com/serverlessworkflow/sdk-go/model"
)


//event state handler
func handleEventState(state *model.EventState, r *Runtime) error {
	fmt.Println("--> Current state Event:", state.GetName())
	if (state.GetTransition() != nil) {
		newStateName := state.Transition.NextState
		ns := findNewStateObject(newStateName, r)
		fmt.Println("New State = ", ns)
		r.begin(ns)
		return nil
	}
	fmt.Println("This is the end..")
	return nil
}

func handleOperationState(state *model.OperationState, r *Runtime) error {
	fmt.Println("--> Current state Operation:", state.GetName())
	// TODO
	// Check for the action Mode (default: sequential)
	if (state.ActionMode == "sequential") {
		fmt.Println("This OperationState is sequential")
		return nil
	}
	fmt.Println("This OperationState is parallel")
	return nil
}

//inject state handler
func HandleInjectState(state *model.InjectState, r *Runtime) error {
	fmt.Println("--> Current state Inject: ", state.GetName())
	if (state.GetTransition() != nil) {
		newStateName  := state.Transition.NextState
		ns := findNewStateObject(newStateName, r)
		//fmt.Println("New State Name = ", newStateName)
		fmt.Println("New State = ", ns)
		r.begin(ns)
		return nil
	}
	fmt.Println("This is the end..")
	return nil
}

func findNewStateObject(name string, r *Runtime) model.State {
	fmt.Println("Searching the next State: ")
	states := r.Workflow.States
	for _, state := range states {
		if  (name == state.GetName()){
			return state
		}
	}
	fmt.Println("Next state not found")
	return nil
}

func HandleDataBasedSwitch(state *model.DataBasedSwitchState, in []byte, r *Runtime) error {
	for _, cond := range state.DataConditions {
		fmt.Println(cond.GetCondition())
		switch cond.(type) {
		case *model.TransitionDataCondition:
			var result map[string]interface{}
			json.Unmarshal(in, &result)
			op, _ := gojq.Parse(cond.GetCondition())
			iter := op.Run(result)
			v, _ := iter.Next()
			if err, ok := v.(error); ok {
				log.Fatalln(err)
			}
			// fmt.Printf("%v\n", v)
			if v.(bool) {
				fmt.Println("GOTO", cond.(*model.TransitionDataCondition).Transition.NextState)
				newStateName  := cond.(*model.TransitionDataCondition).Transition.NextState
				//ns := r.Workflow.States[2].(*model.InjectState)
				ns := findNewStateObject(newStateName, r)
				//fmt.Println("New State Name = ", newStateName)
				fmt.Println("New State = ", ns)
				r.begin(ns)

			} else {
				fmt.Println("Not True")
			}
			// test := map[string]interface{}{"foo": []interface{}{"age", 2, 3}}

			// fmt.Println("Result is:", string(res))

			// return cond.(*model.TransitionDataCondition).Transition.NextState
			// if this condition is true
			// HandleTransition(state, ns)
			//find next state object
			// InferType()
		case *model.EndDataCondition:
			fmt.Println(cond.(*model.EndDataCondition).End)
			fmt.Println("This is the end..")
			// this is the end, you know
		}

	}
	return nil
}
