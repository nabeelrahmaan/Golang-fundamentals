package main

import (
	"Login/middlewear"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("login.html", "dashboard.html")

	store := cookie.NewStore([]byte("secret-key"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})

	const originalUsername = "nabeel"
	const originalPassword = "nabeel123"

	r.POST("/login", middlewear.LoginLogger(), func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		if username == originalUsername && password == originalPassword {
			session := sessions.Default(c)

			session.Set("user", username)
			session.Save()

			c.Redirect(http.StatusFound, "/dashboard")
			return
		}

		c.String(http.StatusUnauthorized, "invalid credentials")
	})

	r.GET("/dashboard", middlewear.AuthMiddleware(), func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"user": user,
		})
	})

	r.GET("/logout", middlewear.LoginLogger(), func(c *gin.Context) {
		session := sessions.Default(c)

		session.Clear()
		session.Save()

		c.Redirect(http.StatusFound, "/")
	})
	r.Run(":8080")
}
