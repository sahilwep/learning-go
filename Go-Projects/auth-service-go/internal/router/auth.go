package router

import (
	"auth-service-go/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
@RegisterAuthRoutes
  - POST: '/signup
  - Handel incoming JSON data & bind it with the SignupRequest struct
  - Hash the password string
  - Send back data to the repository to create user into DB
  - If all successfully done, user-created.
*/
func RegisterAuthRoutes(r *gin.Engine, userRepo *user.Repository) {

	r.POST("/signup", func(c *gin.Context) {

		var req SignupRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		hash, err := user.HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
			return
		}

		err = userRepo.Create(c.Request.Context(), req.Email, hash)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "user created"})
	})
}
