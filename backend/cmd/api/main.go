// backend/cmd/api/main.go

// package declaration defines the code's namespace
package main //tell Go that this is a independent program instead of a tool set for other to use

import (
	"log" //standard lib, for print out log

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

//served as the program's entry point
func main(){
	//initiate a new Fiber instance 
	app := fiber.New(fiber.Config{ 
		AppName: "Spark MVP v0 - Backend", //set the app name 
	})

	//add the middleware to print out all request for debug 它會紀錄每一次進來的 HTTP 請求，像是警衛，所有request 進入API前都會經過
	app.Use(logger.New(logger.Config{
		Format: "${time} ${status} - ${method} ${path} ${latency} \n",
		TimeFormat: "2006-01-02 15:04:05", 
		TimeZone: "Asia/Taipei",
	}))


	//define the routing 定義路由 (Routing)
	//路由定義了「當使用者輸入什麼網址時，伺服器該回傳什麼」。
	// health check 健康檢查端點
	app.Get("api/v1/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "healthy",
			"version": "0.1.0",
			"message": "Spark MVP backend working 🔥",
		})
	})

	//首頁端點
	app.Get("/", func(c * fiber.Ctx) error{ 
		return c.SendString("Welcome to Spark MVP v0.") //只回傳純文字
	})
	//這個 c 就是 Context。它是一個巨大的物件，裝載了這一次 HTTP 請求的所有資訊：client的 IP, data, and what to send back 

	//啟動伺服器 (Listen & Serve)
	port:= "8080"
	log.Printf("Server started from http://localhost:%s", port) //showing message to hint starting in terminal 
	if err := app.Listen(":" +port); err != nil{
		log.Fatalf("Fail to start: %v", err) //print the error value and terminate 
	}
}