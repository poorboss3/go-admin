package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type CustomerGetPageReq struct {
	dto.Pagination `search:"-"`
	Email          string `form:"email"  search:"type:exact;column:Email;table:tb_customer" comment:""`
	Name           string `form:"name"  search:"type:exact;column:Name;table:tb_customer" comment:""`
	AmazonID       string `form:"amazonID"  search:"type:exact;column:AmazonID;table:tb_customer" comment:""`
	CustomerOrder
}

type CustomerOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:tb_customer"`
	Email       string `form:"emailOrder"  search:"type:order;column:Email;table:tb_customer"`
	Name        string `form:"nameOrder"  search:"type:order;column:Name;table:tb_customer"`
	AmazonID    string `form:"amazonIDOrder"  search:"type:order;column:AmazonID;table:tb_customer"`
	PhoneNumber string `form:"phoneNumberOrder"  search:"type:order;column:phone_number;table:tb_customer"`
	Address1    string `form:"address1Order"  search:"type:order;column:address1;table:tb_customer"`
	Address2    string `form:"address2Order"  search:"type:order;column:address2;table:tb_customer"`
	City        string `form:"cityOrder"  search:"type:order;column:city;table:tb_customer"`
	State       string `form:"stateOrder"  search:"type:order;column:state;table:tb_customer"`
	ZipCode     string `form:"zipCodeOrder"  search:"type:order;column:zip_code;table:tb_customer"`
	Country     string `form:"countryOrder"  search:"type:order;column:country;table:tb_customer"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_customer"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_customer"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_customer"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_customer"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_customer"`
}

func (m *CustomerGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type CustomerInsertReq struct {
	Id          int    `json:"-" comment:""` //
	Email       string `json:"email" comment:""`
	Name        string `json:"name" comment:""`
	AmazonID    string `json:"amazonID" comment:""`
	PhoneNumber string `json:"phoneNumber" comment:""`
	Address1    string `json:"address1" comment:""`
	Address2    string `json:"address2" comment:""`
	City        string `json:"city" comment:""`
	State       string `json:"state" comment:""`
	ZipCode     string `json:"zipCode" comment:""`
	Country     string `json:"country" comment:""`
	common.ControlBy
}

func (s *CustomerInsertReq) Generate(model *models.Customer) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Email = s.Email
	model.Name = s.Name
	model.AmazonID = s.AmazonID
	model.PhoneNumber = s.PhoneNumber
	model.Address1 = s.Address1
	model.Address2 = s.Address2
	model.City = s.City
	model.State = s.State
	model.ZipCode = s.ZipCode
	model.Country = s.Country
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *CustomerInsertReq) GetId() interface{} {
	return s.Id
}

type CustomerUpdateReq struct {
	Id          int    `uri:"id" comment:""` //
	Email       string `json:"email" comment:""`
	Name        string `json:"name" comment:""`
	AmazonID    string `json:"amazonID" comment:""`
	PhoneNumber string `json:"phoneNumber" comment:""`
	Address1    string `json:"address1" comment:""`
	Address2    string `json:"address2" comment:""`
	City        string `json:"city" comment:""`
	State       string `json:"state" comment:""`
	ZipCode     string `json:"zipCode" comment:""`
	Country     string `json:"country" comment:""`
	common.ControlBy
}

func (s *CustomerUpdateReq) Generate(model *models.Customer) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Email = s.Email
	model.Name = s.Name
	model.AmazonID = s.AmazonID
	model.PhoneNumber = s.PhoneNumber
	model.Address1 = s.Address1
	model.Address2 = s.Address2
	model.City = s.City
	model.State = s.State
	model.ZipCode = s.ZipCode
	model.Country = s.Country
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *CustomerUpdateReq) GetId() interface{} {
	return s.Id
}

// TbCustomerGetReq 功能获取请求参数
type CustomerGetReq struct {
	Id int `uri:"id"`
}

func (s *CustomerGetReq) GetId() interface{} {
	return s.Id
}

// TbCustomerDeleteReq 功能删除请求参数
type CustomerDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *CustomerDeleteReq) GetId() interface{} {
	return s.Ids
}
