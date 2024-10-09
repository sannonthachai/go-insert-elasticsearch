package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
)

type FlowNode struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Process struct {
	ID            string     `json:"id"`
	Key           int        `json:"key"`
	PartitionID   int        `json:"partitionId"`
	BpmnProcessID string     `json:"bpmnProcessId"`
	Name          string     `json:"name"`
	Version       int        `json:"version"`
	StartedByForm bool       `json:"startedByForm"`
	FlowNodes     []FlowNode `json:"flowNodes"`
}

type DataWrapper struct {
	Data []Process `json:"data"`
}

func jsonStruct(doc Process) string {
	// Create struct instance of the Elasticsearch fields struct object
	docStruct := &Process{
		ID:            doc.ID,
		Key:           doc.Key,
		PartitionID:   doc.PartitionID,
		BpmnProcessID: doc.BpmnProcessID,
		Name:          doc.Name,
		Version:       doc.Version,
		StartedByForm: doc.StartedByForm,
		FlowNodes:     doc.FlowNodes,
	}

	fmt.Println("\ndocStruct:", docStruct)
	fmt.Println("docStruct TYPE:", reflect.TypeOf(docStruct))

	// Marshal the struct to JSON and check for errors
	b, err := json.Marshal(docStruct)
	if err != nil {
		fmt.Println("json.Marshal ERROR:", err)
		return string(err.Error())
	}
	return string(b)
}

func generateData() []string {
	// Open the JSON file
	jsonFile, err := os.Open("data.json")
	if err != nil {
		log.Fatalf("Failed to open JSON file: %s", err)
	}
	defer jsonFile.Close()

	// Read the file's contents
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %s", err)
	}

	// Unmarshal the JSON into the DataWrapper struct
	var dataWrapper DataWrapper
	err = json.Unmarshal(byteValue, &dataWrapper)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}

	var docs []string

	// Process each item in the data array
	for _, process := range dataWrapper.Data {
		jsonStr := jsonStruct(process)
		docs = append(docs, jsonStr)
	}

	return docs
}
