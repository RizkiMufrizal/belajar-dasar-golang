package app

import (
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", controller.FindAll)
	router.GET("/api/categories/:id", controller.FindById)
	router.POST("/api/categories", controller.Create)
	router.PUT("/api/categories/:id", controller.Update)
	router.DELETE("/api/categories/:id", controller.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
