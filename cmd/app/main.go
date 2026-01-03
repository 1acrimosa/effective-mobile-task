package main

import (
	"effective-mobile-task/internal/db"
	handlers "effective-mobile-task/internal/handler"
	"effective-mobile-task/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func main() {
	r := gin.Default()
	dbInstance := db.InitDB()
	repo := &repository.ItemRepository{DB: dbInstance}
	validate := validator.New()
	handler := &handlers.ItemHandler{Repo: repo, Validate: validate}

	api := r.Group("/api")
	{
		api.GET("/items", handler.GetItems)
		api.GET("/items/:id", handler.GetItem)
		api.POST("/items", handler.CreateItem)
		api.PUT("/items/:id", handler.UpdateItem)
		api.DELETE("/items/:id", handler.DeleteItem)
	}

	r.POST("/login", func(c *gin.Context) {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": 1,
		})
		tokenString, _ := token.SignedString([]byte("supersecretkey"))
		c.JSON(200, gin.H{"token": tokenString})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	r.Run(":8080")
}
