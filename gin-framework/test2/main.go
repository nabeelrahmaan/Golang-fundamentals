package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
}

var Users []User
var newid int =1

func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context){
		c.JSON(http.StatusOK, Users)
	} )

	r.POST("/users", func(c *gin.Context) {
		var user User

		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":"invalid JSON",
			})
			return
		}

		user.Id = newid
		newid++
		Users = append(Users, user)
		c.JSON(http.StatusOK, user)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		idparam := c.Param("id")
		id, _ := strconv.Atoi(idparam)
		var newuser User

		if err := c.ShouldBindJSON(&newuser) ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid JSON",
			})
			return
		}
		newuser.Id = id
		for i, u := range Users {
			if u.Id == id {
				Users[i] = newuser
				c.JSON(http.StatusAccepted, newuser)
				return
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})  
	})

	r.PATCH("/users/:id", func(c *gin.Context){
		idparam := c.Param("id")
		id, _ :=strconv.Atoi(idparam)

		var user map[string]interface{}

		if err := c.ShouldBindJSON(&user) ; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}
		for i, u := range Users{
			if u.Id == id {
				if name, ok := user["name"].(string) ; ok {
					Users[i].Name = name
				}
				if age, ok := user["nagee"].(int) ; ok {
					Users[i].Age = age
				}
				c.JSON(http.StatusOK, Users[i])
				return
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found.",
		})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		idparam := c.Param("id")
		id, _ :=strconv.Atoi(idparam)

		for i, u := range Users {
			if u.Id == id {
				Users = append(Users[:i], Users[i+1:]... )
				c.JSON(http.StatusOK, gin.H{
					"message": "User removed.",
				})
				return
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
	})

	r.Run(":8080")
}