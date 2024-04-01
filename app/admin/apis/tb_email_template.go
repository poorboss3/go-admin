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

type EmailTemplate struct {
	api.Api
}

// GetPage 获取EmailTemplate列表
// @Summary 获取EmailTemplate列表
// @Description 获取EmailTemplate列表
// @Tags EmailTemplate
// @Param subject query string false ""
// @Param context query string false ""
// @Param address query string false ""
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.EmailTemplate}} "{"code": 200, "data": [...]}"
// @Router /api/v1/email-template [get]
// @Security Bearer
func (e EmailTemplate) GetPage(c *gin.Context) {
    req := dto.EmailTemplateGetPageReq{}
    s := service.EmailTemplate{}
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
	list := make([]models.EmailTemplate, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取EmailTemplate失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取EmailTemplate
// @Summary 获取EmailTemplate
// @Description 获取EmailTemplate
// @Tags EmailTemplate
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.EmailTemplate} "{"code": 200, "data": [...]}"
// @Router /api/v1/email-template/{id} [get]
// @Security Bearer
func (e EmailTemplate) Get(c *gin.Context) {
	req := dto.EmailTemplateGetReq{}
	s := service.EmailTemplate{}
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
	var object models.EmailTemplate

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取EmailTemplate失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建EmailTemplate
// @Summary 创建EmailTemplate
// @Description 创建EmailTemplate
// @Tags EmailTemplate
// @Accept application/json
// @Product application/json
// @Param data body dto.EmailTemplateInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/email-template [post]
// @Security Bearer
func (e EmailTemplate) Insert(c *gin.Context) {
    req := dto.EmailTemplateInsertReq{}
    s := service.EmailTemplate{}
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
		e.Error(500, err, fmt.Sprintf("创建EmailTemplate失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改EmailTemplate
// @Summary 修改EmailTemplate
// @Description 修改EmailTemplate
// @Tags EmailTemplate
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.EmailTemplateUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/email-template/{id} [put]
// @Security Bearer
func (e EmailTemplate) Update(c *gin.Context) {
    req := dto.EmailTemplateUpdateReq{}
    s := service.EmailTemplate{}
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
		e.Error(500, err, fmt.Sprintf("修改EmailTemplate失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除EmailTemplate
// @Summary 删除EmailTemplate
// @Description 删除EmailTemplate
// @Tags EmailTemplate
// @Param data body dto.EmailTemplateDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/email-template [delete]
// @Security Bearer
func (e EmailTemplate) Delete(c *gin.Context) {
    s := service.EmailTemplate{}
    req := dto.EmailTemplateDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除EmailTemplate失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
