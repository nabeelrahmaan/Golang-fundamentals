package main

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func LoginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)

		fmt.Printf("[Login-Log] %s %s | Status: %d | Duration: %v\n",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		if user == nil {
			c.Redirect(http.StatusFound, "/")
			c.Abort()
			return
		}
		c.Next()
	}
}

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

	r.POST("/login", LoginLogger(), func(c *gin.Context) {
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

	r.GET("/dashboard", AuthMiddleware(), func(c *gin.Context) {
		session := sessions.Default(c)
		user := session.Get("user")

		c.HTML(http.StatusOK, "dashboard.html", gin.H{
			"user": user,
		})
	})

	r.GET("/logout", LoginLogger(), func(c *gin.Context) {
		session := sessions.Default(c)

		session.Clear()
		session.Save()

		c.Redirect(http.StatusFound, "/")
	})
	r.Run(":8080")
}
