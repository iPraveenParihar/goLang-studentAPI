package data

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
)

const studentDataFileName = "student-data.yaml"

type StudentData []Student

type Student struct {
	Name  string `yaml:"name"`
	Age   int    `yaml:"age"`
	Class string `yaml:"class"`
}

func WriteData(student Student) (bool, error) {
	glog.Info("writing data")

	existingStudentData, err := ReadData()
	if err != nil {
		return false, err
	}

	newData := append(existingStudentData, student)
	newStudentData, err := yaml.Marshal(newData)

	if err != nil {
		return false, nil
	}
	if err = os.WriteFile(studentDataFileName, newStudentData, fs.ModeAppend); err != nil {
		return false, nil
	}

	glog.Infof("New Data: %+v\n", newData)
	glog.Error("writing completed!")
	return true, nil
}

func ReadData() (StudentData, error) {
	glog.Info("reading data")

	studentDataFilePath, _ := filepath.Abs("../studentAPI/data/student-data.yaml")
	file, err := ioutil.ReadFile(studentDataFilePath)
	if err != nil {
		return nil, err
	}

	var studentData StudentData
	if err := yaml.Unmarshal(file, &studentData); err != nil {
		return nil, err
	}

	glog.Infof("Stduent Data: %+v\n", studentData)
	glog.Info("read completed!")
	return studentData, nil
}
