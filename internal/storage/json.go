package storage

import (
	"catcatgo/internal/model"
	"encoding/json"
	"os"
)

func Save(path string, functions []model.Function) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(functions)
}

func Load(path string) ([]model.Function, error){
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var functions []model.Function
	err = json.NewDecoder(file).Decode(&functions)
	return functions, err
}
