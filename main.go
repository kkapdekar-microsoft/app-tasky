package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	controller "github.com/jeffthorne/tasky/controllers"
	"github.com/joho/godotenv"
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func openFileReadContent () {
	fileContent, _ := os.ReadFile("wizexercise.txt")
	fmt.Print(string(fileContent))
}

func main() {
	godotenv.Overload()

	openFileReadContent()
	
	router := gin.Default()
	router.LoadHTMLGlob("assets/*.html")
	router.Static("/assets", "./assets")

	router.GET("/", index)
	router.GET("/todos/:userid", controller.GetTodos)
	router.GET("/todo/:id", controller.GetTodo)
	router.POST("/todo/:userid", controller.AddTodo)
	router.DELETE("/todo/:userid/:id", controller.DeleteTodo)
	router.DELETE("/todos/:userid", controller.ClearAll)
	router.PUT("/todo", controller.UpdateTodo)


	router.POST("/signup", controller.SignUp)
	router.POST("/login", controller.Login)
	router.GET("/todo", controller.Todo)

	router.Run(":8080" )

}
