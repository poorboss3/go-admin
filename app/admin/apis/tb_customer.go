package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type Customer struct {
	api.Api
}

// GetPage 获取Customer列表
// @Summary 获取Customer列表
// @Description 获取Customer列表
// @Tags Customer
// @Param email query string false ""
// @Param name query string false ""
// @Param amazonID query string false ""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Customer}} "{"code": 200, "data": [...]}"
// @Router /api/v1/customer [get]
// @Security Bearer
func (e Customer) GetPage(c *gin.Context) {
    req := dto.CustomerGetPageReq{}
    s := service.Customer{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.Customer, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Customer失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Customer
// @Summary 获取Customer
// @Description 获取Customer
// @Tags Customer
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Customer} "{"code": 200, "data": [...]}"
// @Router /api/v1/customer/{id} [get]
// @Security Bearer
func (e Customer) Get(c *gin.Context) {
	req := dto.CustomerGetReq{}
	s := service.Customer{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Customer

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Customer失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建Customer
// @Summary 创建Customer
// @Description 创建Customer
// @Tags Customer
// @Accept application/json
// @Product application/json
// @Param data body dto.CustomerInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/customer [post]
// @Security Bearer
func (e Customer) Insert(c *gin.Context) {
    req := dto.CustomerInsertReq{}
    s := service.Customer{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建Customer失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Customer
// @Summary 修改Customer
// @Description 修改Customer
// @Tags Customer
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.CustomerUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/customer/{id} [put]
// @Security Bearer
func (e Customer) Update(c *gin.Context) {
    req := dto.CustomerUpdateReq{}
    s := service.Customer{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改Customer失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除Customer
// @Summary 删除Customer
// @Description 删除Customer
// @Tags Customer
// @Param data body dto.CustomerDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/customer [delete]
// @Security Bearer
func (e Customer) Delete(c *gin.Context) {
    s := service.Customer{}
    req := dto.CustomerDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除Customer失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
