package service

import (
	"fmt"
	"week2/models"
)

func Validator(e []models.Employee) error {
	for _, employee := range e {
		if employee.Id == 0 {
			return fmt.Errorf("cannot get employee with empty id: %v", employee.Name)
		}
		if employee.Name == "" {
			return fmt.Errorf("cannot get employee with empty name: %v", employee.Id)
		}
		if employee.HireDate == "" {
			return fmt.Errorf("cannot get employee with empty hire date: %v %v", employee.Id, employee.Name)
		}
		if employee.Salary <= 0 {
			return fmt.Errorf("cannot get employee with empty salary: %v", employee.Id)
		}
	}
	return nil
}
