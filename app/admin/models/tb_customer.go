package models

import (
	"go-admin/common/models"
)

type Customer struct {
	models.Model

	Email       string `json:"email" gorm:"type:varchar(500);comment:Email"`
	Name        string `json:"name" gorm:"type:varchar(200);comment:Name"`
	AmazonID    string `json:"amazonID" gorm:"type:varchar(200);comment:AmazonID;column:amazonID"`
	PhoneNumber string `json:"phoneNumber" gorm:"type:varchar(100);comment:phoneNumber;column:phone_number"`
	Address1    string `json:"address1" gorm:"type:varchar(1000);comment:address1;column:address1"`
	Address2    string `json:"address2" gorm:"type:varchar(1000);comment:address2;column:address2"`
	City        string `json:"city" gorm:"type:varchar(100);comment:city;column:city"`
	State       string `json:"state" gorm:"type:varchar(100);comment:state;column:state"`
	ZipCode     string `json:"zipCode" gorm:"type:varchar(100);comment:zipCode;column:zip_code"`
	Country     string `json:"country" gorm:"type:varchar(100);comment:country;column:country"`
	models.ModelTime
	models.ControlBy
}

func (Customer) TableName() string {
	return "tb_customer"
}

func (e *Customer) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Customer) GetId() interface{} {
	return e.Id
}
