package domain

import "github.com/diogoqds/banking/errs"

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
