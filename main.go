package main

import (
	"log/slog"
	"os"

	"github.com/fnoopv/yuefu/api"
	"github.com/fnoopv/yuefu/middleware"
	"github.com/fnoopv/yuefu/model"
	"github.com/fnoopv/yuefu/repository"
	"github.com/fnoopv/yuefu/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := gorm.Open(sqlite.Open("yuefu.db?cache=shared"), &gorm.Config{})
	if err != nil {
		slog.Error("connect database failed", "errpr", err.Error())
		os.Exit(1)
	}

	if err := db.AutoMigrate(new(model.Library)); err != nil {
		slog.Error("sync database failed", "error", err.Error())
		os.Exit(1)
	}

	libraryRepository := repository.NewLibraryRepositoryXormImpl(db)
	libraryService := service.NewLibraryServiceImpl(libraryRepository)
	libraryAPI := api.NewLibraryAPI(libraryService)

	router := gin.Default()
	router.Use(middleware.Response())

	router.GET("/library", libraryAPI.List)
	router.POST("/library", libraryAPI.Create)
	router.PATCH("/library", libraryAPI.Update)
	router.DELETE("/library/:id", libraryAPI.Delete)

	if err := router.Run(":8080"); err != nil {
		slog.Error("start server failed", "error", err.Error())
		os.Exit(1)
	}
}
