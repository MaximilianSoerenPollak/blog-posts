package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Date   string `json:"date,omitempty"`
	Task   string `json:"task,omitempty"`
	Status string `json:"status,omitempty"`
}

func ReadTaskDBFile(filepath string) ([]Task, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Error parsing json: ", err)
		return nil, err
	}
	return tasks, nil
}

func SaveTaskDBFile(filepath string, tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling tasks:", err)
		return err
	}
	err = os.WriteFile(filepath, data, 0644)
	if err != nil {
		fmt.Println("Error writing file", err)
		return err
	}
	return nil
}
