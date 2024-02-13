package main

import (
	"fmt"

	"workflow-orchestrator-go/internal/workflow"
)

func main() {
	ok := workflow.Handle(workflow.Event{ID: "1", Payload: "work"})
	fmt.Println(ok)
}
