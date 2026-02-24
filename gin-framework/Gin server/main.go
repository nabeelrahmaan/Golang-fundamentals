package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)
type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

var(
	newid int = 1
	Users []User
)

func main(){
	r := gin.Default()

	r.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message" : "hello world",
		})
	})

	api := r.Group("/api")
	api.GET("/users", GetUsers)
	api.POST("/users", PostUser)

	fmt.Println("Server running on port 5000...")
	r.Run(":5000")
}

func GetUsers(c *gin.Context){
	c.JSON(http.StatusOK, Users)
}

func PostUser(c *gin.Context){
	var user User

	if err:= c.ShouldBindJSON(&user) ; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}
	user.Id = newid
	newid++
	Users = append(Users, user) 

	c.JSON(http.StatusCreated, user)
}