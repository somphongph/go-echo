package router

import "github.com/labstack/echo"

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//create groups
	// adminGroup := e.Group("/admin")

	//set all middlewares
	// middlewares.SetMainMiddleWares(e)
	// middlewares.SetAdminMiddlewares(adminGroup)

	//set main routes
	// api.MainGroup(e)

	//set groupRoutes
	// api.AdminGroup(adminGroup)

	// Routes
	// v1 := e.Group("/api/v1")
	// {
	// 	v1.POST("/members", api.PostMember())
	// 	v1.GET("/members", api.GetMembers())
	// 	v1.GET("/members/:id", api.GetMember())
	// }
	return e
}
