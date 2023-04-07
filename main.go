package main

import (
	"fmt"
	"path/filepath"

	"github.com/anyingiit/GoReactResourceManagement/db"
	"github.com/anyingiit/GoReactResourceManagement/globalVars"
	"github.com/anyingiit/GoReactResourceManagement/middleware"
	"github.com/anyingiit/GoReactResourceManagement/router"
	"github.com/anyingiit/GoReactResourceManagement/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// fmt.Println("Hello, playground")
	projectRootPath, err := filepath.Abs(filepath.Dir("./main.go"))
	if err != nil {
		panic(fmt.Errorf("failed to get project root path, %v", err))
	}
	fmt.Println(projectRootPath)
	err = globalVars.SetProjectRootPath(projectRootPath)
	if err != nil {
		panic(fmt.Errorf("failed to set project root path, %v", err))
	}

	config, err := utils.ReadConfigFile(filepath.Join(projectRootPath, "config", "config.yml"))
	if err != nil {
		panic(fmt.Errorf("failed to read config file, %v", err))
	}
	err = globalVars.SetProjectConfig(config)
	if err != nil {
		panic(fmt.Errorf("failed to set project config, %v", err))
	}

	db, err := db.InitDB(&config.Database)
	if err != nil {
		panic(fmt.Errorf("failed to init db, %v", err))
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// 创建gin实例
	g := gin.New()

	// 注册中间件
	g.Use(gin.Logger(), gin.Recovery(), middleware.ErrorHandler())

	//init router
	router.InitRouter(g, db)

	// create server address string, and print it
	serverAddress := fmt.Sprintf("%s:%d", config.Server.Ip, config.Server.Port)
	fmt.Println("Server is running on", serverAddress)

	g.Run(serverAddress)
}

//TODO: 首次运行时, 应该执行一些操作
