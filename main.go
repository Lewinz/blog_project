package main

import (
	"blog_project/blog"
	"blog_project/components/config"
	"blog_project/components/db"
	"blog_project/filter"
	"blog_project/models"
	"blog_project/user"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func initConfig() {
	// -------------------------------------------------------
	// 初始化配置文件，从命令行中服务参数
	config.SetConfigType("yaml")

	var configPath string
	flag.StringVar(&configPath, "f", "", "配置文件路径'-f /mnt/config'")
	flag.Parse()

	file, err := os.Open(configPath)
	if err != nil {
		return
	}
	defer file.Close()

	if err = config.ReadConfig(file); err != nil {
		panic(fmt.Errorf("fatal error read config file: %s", err))
	}

	// 初始化数据库配置
	driver := db.Driver{}
	if err = config.UnmarshalKey("database", &driver); err != nil {
		panic(fmt.Errorf("fatal error unmarsha config: %s", err))
	}

	db.Register(driver)
}

func main() {
	initConfig()

	models.AutoMigrate()

	engine := gin.Default()

	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "the page 404",
		})
	})

	engine.GET("/createToken", filter.CreateToken)

	engine.POST("/user/register", user.Create) // 用户注册

	group := engine.Group("/blog", filter.AuthCheck)
	{
		group.POST("/create", blog.Create)
	}

	srv := &http.Server{
		Addr:    ":" + config.GetString("app.port"),
		Handler: engine,
	}

	if config.GetString("app.mode") == gin.DebugMode {
		fmt.Printf("Listening and serving HTTP on %s\n", srv.Addr)
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("Failed to listen and serve %v", err)
			os.Exit(1)
		}
	} else {
		gin.SetMode(gin.ReleaseMode)
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("Failed to listen and serve %v", err)
			os.Exit(1)
		}
	}
}
