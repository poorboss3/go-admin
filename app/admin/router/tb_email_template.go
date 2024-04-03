package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerEmailTemplateRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerEmailTemplateNoCheck)
}

// registerEmailTemplateRouter
func registerEmailTemplateRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.EmailTemplate{}
	r := v1.Group("/email-template").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
	}
}

// registerCustomerRouterNoCheck
func registerEmailTemplateNoCheck(v1 *gin.RouterGroup) {
	api := apis.EmailTemplate{}
	r := v1.Group("/email-template")
	{
		r.GET("/:id", actions.PermissionAction(), api.Get)
	}
}
