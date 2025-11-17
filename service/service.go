package service

import (
	"encoding/json"
	"fmt"
	"os"
	"week2/models"
)

var employees []models.Employee

func EmployeeGetAll(name string) ([]models.Employee, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("cannot read employee file: %v", err)
	}

	err = json.Unmarshal(data, &employees)
	if err != nil {
		return nil, fmt.Errorf("cannot parse employee file: %v", err)
	}

	//check
	err = Validator(employees)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func EmployeeGetById(name string, employeeId int) (models.Employee, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return models.Employee{}, fmt.Errorf("cannot read employee file: %v", err)
	}
	err = json.Unmarshal(data, &employees)
	if err != nil {
		return models.Employee{}, fmt.Errorf("cannot parse employee file: %v", err)
	}

	//check
	err = Validator(employees)
	if err != nil {
		return models.Employee{}, err
	}

	for _, employee := range employees {
		if employee.Id == employeeId {
			return employee, nil
		}
	}

	return models.Employee{}, nil
}
