package models

import (

	"go-admin/common/models"

)

type EmailTemplate struct {
    models.Model
    
    Subject string `json:"subject" gorm:"type:varchar(200);comment:Subject"` 
    Context string `json:"context" gorm:"type:text;comment:Context"` 
    Address string `json:"address" gorm:"type:varchar(100);comment:Address"` 
    models.ModelTime
    models.ControlBy
}

func (EmailTemplate) TableName() string {
    return "tb_email_template"
}

func (e *EmailTemplate) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *EmailTemplate) GetId() interface{} {
	return e.Id
}