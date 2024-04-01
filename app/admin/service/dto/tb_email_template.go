package dto

import (

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type EmailTemplateGetPageReq struct {
	dto.Pagination     `search:"-"`
    Subject string `form:"subject"  search:"type:exact;column:Subject;table:tb_email_template" comment:""`
    Context string `form:"context"  search:"type:exact;column:Context;table:tb_email_template" comment:""`
    Address string `form:"address"  search:"type:exact;column:address;table:tb_email_template" comment:""`
    EmailTemplateOrder
}

type EmailTemplateOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:tb_email_template"`
    Subject string `form:"subjectOrder"  search:"type:order;column:Subject;table:tb_email_template"`
    Context string `form:"contextOrder"  search:"type:order;column:Context;table:tb_email_template"`
    Address string `form:"addressOrder"  search:"type:order;column:address;table:tb_email_template"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_email_template"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_email_template"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_email_template"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_email_template"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_email_template"`
    
}

func (m *EmailTemplateGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type EmailTemplateInsertReq struct {
    Id int `json:"-" comment:""` // 
    Subject string `json:"subject" comment:""`
    Context string `json:"context" comment:""`
    Address string `json:"address" comment:""`
    common.ControlBy
}

func (s *EmailTemplateInsertReq) Generate(model *models.EmailTemplate)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Subject = s.Subject
    model.Context = s.Context
    model.Address = s.Address
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *EmailTemplateInsertReq) GetId() interface{} {
	return s.Id
}

type EmailTemplateUpdateReq struct {
    Id int `uri:"id" comment:""` // 
    Subject string `json:"subject" comment:""`
    Context string `json:"context" comment:""`
    Address string `json:"address" comment:""`
    common.ControlBy
}

func (s *EmailTemplateUpdateReq) Generate(model *models.EmailTemplate)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Subject = s.Subject
    model.Context = s.Context
    model.Address = s.Address
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *EmailTemplateUpdateReq) GetId() interface{} {
	return s.Id
}

// EmailTemplateGetReq 功能获取请求参数
type EmailTemplateGetReq struct {
     Id int `uri:"id"`
}
func (s *EmailTemplateGetReq) GetId() interface{} {
	return s.Id
}

// EmailTemplateDeleteReq 功能删除请求参数
type EmailTemplateDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *EmailTemplateDeleteReq) GetId() interface{} {
	return s.Ids
}
