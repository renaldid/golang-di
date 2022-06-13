package main

import (
	"github.com/go-playground/validator/v10"
	"golang-rest-api/app"
	"golang-rest-api/controller"
	"golang-rest-api/helper"
	"golang-rest-api/middleware"
	"golang-rest-api/repository"
	"golang-rest-api/service"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	//1. Membuat objek bernama categoryRepository
	//2. Memanggil constructor bernama NewCategoryRepository
	//3. Tanpa ada dependency
	categoryRepository := repository.NewCategoryRepository()

	//1. Membuat object bernama categoryService
	//2. Memanggil constructor bernama NewCategoryService
	//3. Meng-inject dependency bernama categoryRepository, db, validate
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	//1. Membuat object bernama categoryController
	//2. Memanggil constructor bernama NewCategoryController
	//3. Meng-inject dependency bernama categoryService
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
