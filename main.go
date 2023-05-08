package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"gopkg.in/yaml.v3"
)

const studentDataFileName = "student-data.yaml"

type StudentData []Student

type Student struct {
	Name	string 	`yaml:"name"`
	Age		int		`yaml:"age"`
	Class	string	`yaml:"class"`
}


func writeData(student Student) (bool, error) {
	log.Println("writing data")

	existingStudentData, err := readData()
	if err != nil {
		return false, err
	}

	newData := append(existingStudentData, student)
	newStudentData, err := yaml.Marshal(newData);

	if err != nil {
		return false, nil
	}
	if err = os.WriteFile(studentDataFileName, newStudentData, fs.ModeAppend); err != nil {
		return false, nil
	}

	log.Printf("New Data: %+v\n", newData)
	log.Fatal("writing completed!")
	return true, nil
}

func readData() (StudentData, error) {
	log.Println("reading data")

	file, err := ioutil.ReadFile(studentDataFileName)
	if err != nil {
		return nil, err
	}

	var studentData StudentData	
	if err := yaml.Unmarshal(file, &studentData); err != nil {
		return nil, err
	}

	log.Printf("Stduent Data: %+v\n", studentData)
	log.Print("read completed!")
	return studentData, nil
}

func main() {
	log.SetOutput(os.Stderr)
	
	readData()
	student := Student{
		Name: "rohit",
		Age: 25,
		Class: "BNMIT",
	}
	writeData(student)	
}