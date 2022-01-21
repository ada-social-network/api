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
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var version = "dev"

const (
	basePath     = "/api/rest/v1"
	basePathAuth = "/auth"
)

// CORS used for adding cors support
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	var wait time.Duration
	var port int
	var host string
	var mode string
	var dsn string
	var withAuth bool
	var showVersion bool

	flag.BoolVar(&withAuth, "auth", true, "Use api authentication")
	flag.BoolVar(&showVersion, "version", false, "Show application current version")
	flag.IntVar(&port, "http-port", 8080, "Default port")
	flag.StringVar(&host, "http-host", "0.0.0.0", "Default interface")
	flag.StringVar(&mode, "mode", gin.ReleaseMode, "Running mode, can be 'debug', 'release' or 'test'")
	flag.StringVar(&dsn, "sqlite-dsn", "gorm.db", "sqlite database file (dsn) that will store data")
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	if showVersion {
		fmt.Printf("Current version: %s\n", version)
		return
	}
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed", err)
	}

	err = db.AutoMigrate(&models.Post{}, &models.User{}, &models.BdaPost{}, &models.Promo{}, &models.Comment{}, &models.Category{}, &models.Topic{}, &models.Like{})

	if err != nil {
		log.Fatal("Automigration failed", err)
	}

	gin.SetMode(mode)

	r := gin.New()
	r.
		Use(CORS()).
		// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
		Use(gin.Logger()).
		// Recovery middleware recovers from any panics and writes a 500 if there was one.
		Use(gin.Recovery()).
		// Add in the response current version details
		Use(middleware.Version(version)).
		GET("/ping", handler.Ping)

	authMiddleware, err := middleware.CreateAuthMiddleware(db)
	if err != nil {
		log.Fatal(err)
	}

	r.Group(basePathAuth).
		POST("/register", handler.Register(db)).
		POST("/login", authMiddleware.LoginHandler).
		GET("/refresh", authMiddleware.RefreshHandler)

	protected := r.Group(basePath)

	if withAuth {
		protected.Use(authMiddleware.MiddlewareFunc())
	}

	protected.
		GET("/me", handler.MeHandler(db)).
		GET("/posts", handler.ListPostHandler(db)).
		GET("/posts/:id", handler.GetPostHandler(db)).
		PATCH("/posts/:id", handler.UpdatePostHandler(db)).
		DELETE("/posts/:id", handler.DeletePostHandler(db)).
		GET("/users", handler.ListUserHandler(db)).
		GET("/users/:id", handler.GetUserHandler(db)).
		POST("/users", handler.CreateUserHandler(db)).
		PATCH("/users/:id", handler.UpdateUserHandler(db)).
		DELETE("/users/:id", handler.DeleteUserHandler(db)).
		GET("/bdaposts", handler.ListBdaPost(db)).
		GET("/bdaposts/:id", handler.GetBdaPost(db)).
		POST("/bdaposts", handler.CreateBdaPost(db)).
		PATCH("/bdaposts/:id", handler.UpdateBdaPost(db)).
		DELETE("/bdaposts/:id", handler.DeleteBdaPost(db)).
		GET("/bdaposts/:id/comments", handler.ListBdaPostComments(db)).
		GET("/bdaposts/:id/comments/:commentId", handler.GetBdaPostComment(db)).
		POST("/bdaposts/:id/comments", handler.CreateBdaPostComment(db)).
		PATCH("/bdaposts/:id/comments/:commentId", handler.UpdateBdaPostComment(db)).
		DELETE("bdaposts/:id/comments/:commentId", handler.DeleteBdaPostComment(db)).
		GET("/bdaposts/:id/likes", handler.ListBdaPostLikes(db)).
		POST("/bdaposts/:id/likes", handler.CreateBdaPostLike(db)).
		GET("/promos", handler.ListPromo(db)).
		POST("/promos", handler.CreatePromo(db)).
		GET("/promos/:id/users", handler.ListPromoUsers(db)).
		PATCH("/promos/:id", handler.UpdatePromo(db)).
		DELETE("/promos/:id", handler.DeletePromo(db)).
		GET("/categories", handler.ListCategories(db)).
		POST("/categories", handler.CreateCategory(db)).
		DELETE("/categories/:id", handler.DeleteCategory(db)).
		GET("/categories/:id/topics", handler.ListCategoryTopics(db)).
		GET("/topics", handler.ListTopics(db)).
		POST("/categories/:id/topics", handler.CreateTopic(db)).
		PATCH("/topics/:id", handler.UpdateTopic(db)).
		DELETE("/topics/:id", handler.DeleteTopic(db)).
		GET("/topics/:id", handler.GetTopic(db)).
		GET("/topics/:id/posts", handler.ListTopicPosts(db)).
		POST("/topics/:id/posts", handler.CreatePost(db))

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
