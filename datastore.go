package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type InProgress struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Done struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// JSON-to-Go is amazeballs https://mholt.github.io/json-to-go/
type Columns struct {
	Todo       []Todo       `json:"todo"`
	InProgress []InProgress `json:"inProgress"`
	Done       []Done       `json:"done"`
}

func ReadFromStorage() Columns {

	var columns Columns

	if _, err := os.Stat("storage.json"); err == nil {
		jsonFile, err := os.Open("storage.json")

		if err != nil {
			fmt.Println("JSON IO ERROR: " + err.Error())
		}

		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		// we unmarshal our byteArray which contains our
		// jsonFile's content into 'columns' which we defined above
		json.Unmarshal(byteValue, &columns)
	}

	return columns
}

func WriteToStorage(m Model) {

	var columns Columns

	for _, element := range m.lists[todo].Items() {
		columns.Todo = append(columns.Todo, Todo{element.(Task).title, element.(Task).description})
	}

	for _, element := range m.lists[inProgress].Items() {
		columns.InProgress = append(columns.InProgress, InProgress{element.(Task).title, element.(Task).description})
	}

	for _, element := range m.lists[done].Items() {
		columns.Done = append(columns.Done, Done{element.(Task).title, element.(Task).description})
	}

	storageFile := "storage.json"

	if _, err := os.Stat(storageFile); err == nil {
		e := os.Remove(storageFile)
		if e != nil {
			fmt.Println(e)
		}
	}

	_, e := os.Create(storageFile)
	if e != nil {
		fmt.Println(e)
	}

	jsonString, _ := json.MarshalIndent(columns, "", " ")
	_ = ioutil.WriteFile(storageFile, jsonString, 0644)
}
