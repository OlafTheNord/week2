package service

import (
	"encoding/json"
	"fmt"
	"os"
	"week2/models"
)

var employees []models.Employee

func EmployeeGetAll(name string) (error, []models.Employee) {
	data, err := os.ReadFile(name)
	if err != nil {
		return fmt.Errorf("cannot read employee file: %v", err), nil
	}

	err = json.Unmarshal(data, &employees)
	if err != nil {
		return fmt.Errorf("cannot parse employee file: %v", err), nil
	}
	/*
		//check
		err = Validator(employees)
		if err != nil {
			return err, nil
		}

	*/
	return nil, employees
}

func EmployeeGetById(name string, employeeId int) (error, models.Employee) {
	data, err := os.ReadFile(name)
	if err != nil {
		return fmt.Errorf("cannot read employee file: %v", err), models.Employee{}
	}
	err = json.Unmarshal(data, &employees)
	if err != nil {
		return fmt.Errorf("cannot parse employee file: %v", err), models.Employee{}
	}

	//check
	err = Validator(employees)
	if err != nil {
		return err, models.Employee{}
	}

	for _, employee := range employees {
		if employee.Id == employeeId {
			return nil, employee
		}
	}

	return nil, models.Employee{}
}
