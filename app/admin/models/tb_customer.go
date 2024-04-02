package models

import (
	"go-admin/common/models"
)

type Customer struct {
	models.Model

	Email    string `json:"email" gorm:"type:varchar(500);comment:Email"`
	Name     string `json:"name" gorm:"type:varchar(200);comment:Name"`
	AmazonID string `json:"amazonID" gorm:"type:varchar(200);comment:AmazonID;column:amazonID"`
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
