package drivers

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type JsonDriver struct {
	filepath string
}

func NewJsonDriver(filepath string) *JsonDriver {
	return &JsonDriver{filepath: filepath}
}

func (j *JsonDriver) Write(object interface{}) error {
	file, err := json.MarshalIndent(object, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(j.filepath, file, 0644)
}

func (j *JsonDriver) DeleteFile() error {
	return os.Remove(j.filepath)
}
