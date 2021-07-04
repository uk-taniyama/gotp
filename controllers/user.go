package controllers

import (
	"context"
	"database/sql"
	"local/models"
	"local/services"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const sessionKey = "user"

type LoginRequest struct {
	User     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func AuthGuard(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get(sessionKey) == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func GetUser(c *gin.Context) string {
	session := sessions.Default(c)
	return session.Get(sessionKey).(string)
}

func PostLogin(c *gin.Context) {
	username, _ := c.GetPostForm("username")
	password, _ := c.GetPostForm("password")
	logger.Info(username, password)
	var user *models.User = nil
	services.UseDB(func(db *sql.DB) error {
		// cols := &models.UserColumns
		where := &models.UserWhere
		user, _ = models.Users(
			where.Username.EQ(username),
			where.Password.EQ(password),
		).One(context.Background(), db)
		return nil
	})
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unkown user",
		})
		return
	}
	session := sessions.Default(c)
	session.Set(sessionKey, user.Username)
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
	})
}

func GetLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func PostHello(c *gin.Context) {
	message, _ := c.GetPostForm("message")
	user := GetUser(c)
	c.JSON(http.StatusUnauthorized, gin.H{
		"message": message + " " + user,
	})
}
