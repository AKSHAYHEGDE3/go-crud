package controllers

import (
	"log"
	"net/http"
	"os"

	"time"

	"github.com/AKSHAYHEGDE3/go-crud/initializers"
	"github.com/AKSHAYHEGDE3/go-crud/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var body struct {
		Email    string
		Password string
		Name     string
	}

	c.Bind(&body)

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID != 0 {
		c.JSON(400, gin.H{
			"error": "User already exists",
		})
		return
	}

	if len(body.Password) < 6 {
		c.JSON(400, gin.H{
			"error": "Password must be at least 6 characters",
		})
		return

	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(500, gin.H{
			"message": "problem hashing password",
		})
		return
	}

	newuser := models.User{Email: body.Email, Name: body.Name, Password: string(hash)}
	result := initializers.DB.Create(&newuser)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "Error while creating the user",
		})
		return
	}

	userDetails := gin.H{
		"id":    newuser.ID,
		"email": newuser.Email,
		"name":  newuser.Name,
	}

	log.Println("result")

	c.JSON(200, gin.H{
		"user": userDetails,
	})

}

func LoginUser(c *gin.Context) {
	var body struct {
		Email    string
		Password string
	}

	c.Bind(&body)

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(400, gin.H{
			"error": "User does not exist",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid Password or Email",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid": user.ID,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenString, 3600*24*30, "", "", false, true)

	userDetails := gin.H{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
	}

	c.JSON(200, gin.H{
		"user": userDetails,
	})

}

func Validate(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(500, gin.H{"error": "User not found"})
		return
	}

	// Check if the user has the expected type
	userObj, ok := user.(models.User)
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid user type"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":    userObj.ID,
		"email": userObj.Email,
		"name":  userObj.Name,
	})
}
