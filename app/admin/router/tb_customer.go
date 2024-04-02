package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerCustomerRouter)
	routerNoCheckRole = append(routerNoCheckRole, registerCustomerRouterNoCheck)
}

// registerCustomerRouter
func registerCustomerRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Customer{}
	r := v1.Group("/customer").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		//r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
	}
}

// registerCustomerRouterNoCheck
func registerCustomerRouterNoCheck(v1 *gin.RouterGroup) {
	api := apis.Customer{}
	r := v1.Group("/customer")
	{
		r.POST("", api.Insert)
	}
}
