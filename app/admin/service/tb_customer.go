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

type Customer struct {
	service.Service
}

// GetPage 获取Customer列表
func (e *Customer) GetPage(c *dto.CustomerGetPageReq, p *actions.DataPermission, list *[]models.Customer, count *int64) error {
	var err error
	var data models.Customer

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("CustomerService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Customer对象
func (e *Customer) Get(d *dto.CustomerGetReq, p *actions.DataPermission, model *models.Customer) error {
	var data models.Customer

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetCustomer error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Customer对象
func (e *Customer) Insert(c *dto.CustomerInsertReq) error {
    var err error
    var data models.Customer
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("CustomerService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Customer对象
func (e *Customer) Update(c *dto.CustomerUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Customer{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("CustomerService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Customer
func (e *Customer) Remove(d *dto.CustomerDeleteReq, p *actions.DataPermission) error {
	var data models.Customer

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveCustomer error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
