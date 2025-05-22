package main

import (
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/middleware"
)

// karena server tidak mempunyai provider, kita membuat diatas sini sendiri
func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	//db := app.NewDB()
	//validate := validator.New()
	//categoryRepository := repository.NewCategoryRepository()
	//categoryService := service.NewCategoryService(categoryRepository, db, validate)
	//categoryController := controller.NewCategoryController(categoryService)
	//router := app.NewRouter(categoryController)
	//authMiddleware := middleware.NewAuthMiddleware(router)

	//hanya perlu ini
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
