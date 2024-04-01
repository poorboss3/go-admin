package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type CustomerGetPageReq struct {
	dto.Pagination     `search:"-"`
    Email string `form:"email"  search:"type:exact;column:Email;table:tb_customer" comment:""`
    Name string `form:"name"  search:"type:exact;column:Name;table:tb_customer" comment:""`
    AmazonID string `form:"amazonID"  search:"type:exact;column:AmazonID;table:tb_customer" comment:""`
    CustomerOrder
}

type CustomerOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:tb_customer"`
    Email string `form:"emailOrder"  search:"type:order;column:Email;table:tb_customer"`
    Name string `form:"nameOrder"  search:"type:order;column:Name;table:tb_customer"`
    AmazonID string `form:"amazonIDOrder"  search:"type:order;column:AmazonID;table:tb_customer"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_customer"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_customer"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_customer"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_customer"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_customer"`
    
}

func (m *CustomerGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type CustomerInsertReq struct {
    Id int `json:"-" comment:""` // 
    Email string `json:"email" comment:""`
    Name string `json:"name" comment:""`
    AmazonID string `json:"amazonID" comment:""`
    common.ControlBy
}

func (s *CustomerInsertReq) Generate(model *models.Customer)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Email = s.Email
    model.Name = s.Name
    model.AmazonID = s.AmazonID
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *CustomerInsertReq) GetId() interface{} {
	return s.Id
}

type CustomerUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Email string `json:"email" comment:""`
    Name string `json:"name" comment:""`
    AmazonID string `json:"amazonID" comment:""`
    common.ControlBy
}

func (s *CustomerUpdateReq) Generate(model *models.Customer)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Email = s.Email
    model.Name = s.Name
    model.AmazonID = s.AmazonID
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *CustomerUpdateReq) GetId() interface{} {
	return s.Id
}

// CustomerGetReq 功能获取请求参数
type CustomerGetReq struct {
     Id int `uri:"id"`
}
func (s *CustomerGetReq) GetId() interface{} {
	return s.Id
}

// CustomerDeleteReq 功能删除请求参数
type CustomerDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *CustomerDeleteReq) GetId() interface{} {
	return s.Ids
}
