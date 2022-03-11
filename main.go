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
	"github.com/ada-social-network/api/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
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

	commentRepository := repository.NewCommentRepository(db)
	commentHandler := handler.NewCommentHandler(commentRepository)

	bdaPostRepository := repository.NewBdaPostRepository(db)
	bdaPostHandler := handler.NewBdaPostHandler(bdaPostRepository)

	postRepository := repository.NewPostRepository(db)
	postHandler := handler.NewPostHandler(postRepository)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryHandler := handler.NewCategoryHandler(categoryRepository)

	protected.
		GET("/me", handler.Me(db)).
		PATCH("/me/password", handler.UpdatePassword(db)).
		GET("/topics/:id/posts", postHandler.ListPost).
		GET("/topics/:id/posts/:postId", postHandler.GetPost).
		POST("/topics/:id/posts", postHandler.CreatePost).
		PATCH("/topics/:id/posts/:postId", postHandler.UpdatePost).
		DELETE("/topics/:id/posts/:postId", postHandler.DeletePost).
		GET("/posts/:id/likes", handler.ListPostLikes(db)).
		POST("/posts/:id/likes", handler.CreatePostLike(db)).
		DELETE("/posts/:id/likes/:likeId", handler.DeletePostLike(db)).
		GET("/users", handler.ListUser(db)).
		GET("/users/:id", handler.GetUser(db)).
		POST("/users", handler.CreateUser(db)).
		PATCH("/users/:id", handler.UpdateUser(db)).
		DELETE("/users/:id", handler.DeleteUser(db)).
		GET("/bdaposts", bdaPostHandler.ListBdaPost).
		GET("/bdaposts/:id", bdaPostHandler.GetBdaPost).
		POST("/bdaposts", bdaPostHandler.CreateBdaPost).
		PATCH("/bdaposts/:id", bdaPostHandler.UpdateBdaPost).
		DELETE("/bdaposts/:id", bdaPostHandler.DeleteBdaPost).
		GET("/bdaposts/:id/likes", handler.ListBdaPostLikes(db)).
		POST("/bdaposts/:id/likes", handler.CreateBdaPostLike(db)).
		DELETE("/bdaposts/:id/likes/:likeId", handler.DeleteBdaPostLike(db)).
		GET("/bdaposts/:id/comments", commentHandler.ListBdaPostComments).
		GET("/bdaposts/:id/comments/:commentId", commentHandler.GetBdaPostComment).
		POST("/bdaposts/:id/comments", commentHandler.CreateBdaPostComment).
		PATCH("/bdaposts/:id/comments/:commentId", commentHandler.UpdateBdaPostComment).
		DELETE("bdaposts/:id/comments/:commentId", commentHandler.DeleteBdaPostComment).
		GET("/comments/:id/likes", handler.ListCommentLikes(db)).
		POST("/comments/:id/likes", handler.CreateCommentLike(db)).
		DELETE("/comments/:id/likes/:likeId", handler.DeleteCommentLike(db)).
		GET("/promos", handler.ListPromo(db)).
		POST("/promos", handler.CreatePromo(db)).
		GET("/promos/:id/users", handler.ListPromoUsers(db)).
		PATCH("/promos/:id", handler.UpdatePromo(db)).
		DELETE("/promos/:id", handler.DeletePromo(db)).
		GET("/categories", categoryHandler.ListCategories).
		POST("/categories", categoryHandler.CreateCategory).
		DELETE("/categories/:id", categoryHandler.DeleteCategory).
		GET("/categories/:id/topics", handler.ListCategoryTopics(db)).
		GET("/topics", handler.ListTopics(db)).
		POST("/categories/:id/topics", handler.CreateTopic(db)).
		PATCH("/topics/:id", handler.UpdateTopic(db)).
		DELETE("/topics/:id", handler.DeleteTopic(db)).
		GET("/topics/:id", handler.GetTopic(db))

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
