package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ada-social-network/api/handler"
	"github.com/ada-social-network/api/middleware"
	"github.com/ada-social-network/api/models"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	version  = "dev"
	basePath = "/api/rest/v1"
)

func main() {
	var wait time.Duration
	var port int
	var host string
	var mode string
	var dsn string

	flag.IntVar(&port, "http-port", 8080, "Default port")
	flag.StringVar(&host, "http-host", "0.0.0.0", "Default interface")
	flag.StringVar(&mode, "mode", gin.ReleaseMode, "Running mode, can be 'debug', 'release' or 'test'")
	flag.StringVar(&dsn, "sqlite-dsn", "gorm.db", "sqlite database file (dsn) that will store data")
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed", err)
	}

	err = db.AutoMigrate(&models.Post{}, &models.User{}, &models.BdaPost{})
	if err != nil {
		log.Fatal("Automigration failed", err)
	}

	gin.SetMode(mode)

	r := gin.New()
	r.
		Use(cors.Default()).
		// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
		Use(gin.Logger()).
		// Recovery middleware recovers from any panics and writes a 500 if there was one.
		Use(gin.Recovery()).
		// Add in the response current version details
		Use(middleware.Version(version)).
		GET("/ping", handler.Ping).
		GET(basePath+"/posts", handler.ListPostHandler(db)).
		GET(basePath+"/posts/:id", handler.GetPostHandler(db)).
		POST(basePath+"/posts", handler.CreatePostHandler(db)).
		PATCH(basePath+"/posts/:id", handler.UpdatePostHandler(db)).
		DELETE(basePath+"/posts/:id", handler.DeletePostHandler(db)).
		GET(basePath+"/users", handler.ListUserHandler(db)).
		GET(basePath+"/users/:id", handler.GetUserHandler(db)).
		POST(basePath+"/users", handler.CreateUserHandler(db)).
		PATCH(basePath+"/users/:id", handler.UpdateUserHandler(db)).
		DELETE(basePath+"/users/:id", handler.DeleteUserHandler(db)).
		GET(basePath+"/bdaposts", handler.ListBdaPost(db)).
		GET(basePath+"/bdaposts/:id", handler.GetBdaPost(db)).
		POST(basePath+"/bdaposts", handler.CreateBdaPost(db)).
		PATCH(basePath+"/bdaposts/:id", handler.UpdateBdaPost(db)).
		DELETE(basePath+"/bdaposts/:id", handler.DeleteBdaPost(db))

	srv := &http.Server{
		Addr: fmt.Sprintf("%s:%d", host, port),
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Printf("Http server is started on interface %s:%d", host, port)

		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
	os.Exit(0)

}
