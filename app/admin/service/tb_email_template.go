package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type EmailTemplate struct {
	service.Service
}

// GetPage 获取EmailTemplate列表
func (e *EmailTemplate) GetPage(c *dto.EmailTemplateGetPageReq, p *actions.DataPermission, list *[]models.EmailTemplate, count *int64) error {
	var err error
	var data models.EmailTemplate

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("EmailTemplateService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取EmailTemplate对象
func (e *EmailTemplate) Get(d *dto.EmailTemplateGetReq, p *actions.DataPermission, model *models.EmailTemplate) error {
	var data models.EmailTemplate

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetEmailTemplate error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建EmailTemplate对象
func (e *EmailTemplate) Insert(c *dto.EmailTemplateInsertReq) error {
    var err error
    var data models.EmailTemplate
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("EmailTemplateService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改EmailTemplate对象
func (e *EmailTemplate) Update(c *dto.EmailTemplateUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.EmailTemplate{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("EmailTemplateService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除EmailTemplate
func (e *EmailTemplate) Remove(d *dto.EmailTemplateDeleteReq, p *actions.DataPermission) error {
	var data models.EmailTemplate

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveEmailTemplate error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
