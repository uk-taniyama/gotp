package main

import (
	"context"
	"database/sql"
	"fmt"
	"local/controllers"
	"local/models"
	"local/services"
	"local/utils"
	"strings"
	"time"
	"unsafe"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.uber.org/zap"
)

type ZapWriter struct {
	logger *zap.SugaredLogger
}

func (i ZapWriter) Write(p []byte) (n int, err error) {
	i.logger.Info(strings.TrimSpace(*(*string)(unsafe.Pointer(&p))))
	return len(p), nil
}

func main() {
	utils.SetupLogger("./main.log")
	controllers.Init()

	gin.DefaultWriter = ZapWriter{utils.NewLogger("gin")}

	boil.DebugMode = true
	boil.DebugWriter = ZapWriter{utils.NewLogger("boil")}

	services.UseDB(func(db *sql.DB) error {
		count, err := models.Users().Count(context.Background(), db)
		if err != nil {
			fmt.Println("Query:Error")
			return err
		}
		if count == 0 {
			user := models.User{
				Username: "test",
				Password: "test",
			}
			user.Insert(context.Background(), db, boil.Infer())
		}
		return nil
	})

	app := gin.New()

	store := memstore.NewStore([]byte("secret"))
	app.Use(sessions.Sessions("mysession", store))

	app.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	app.Use(ginzap.RecoveryWithZap(zap.L(), true))

	app.POST("/login", controllers.PostLogin)
	app.POST("/hello", controllers.AuthGuard, controllers.PostHello)
	app.GET("/logout", controllers.GetLogout)
	app.GET("/ws", controllers.GetWS)

	server := controllers.NewSocketIOServer()
	defer server.Close()
	app.GET("/socket.io/*any", gin.WrapH(server))
	app.POST("/socket.io/*any", gin.WrapH(server))

	app.Run("0.0.0.0:3000")
}
