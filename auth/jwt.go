package auth

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte(os.Getenv("Secret_key"))

func generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Token required"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) == 2 && strings.ToLower(tokenParts[0]) == "bearer" {
			tokenString = tokenParts[1] // Extract the actual token
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, _ := token.Claims.(jwt.MapClaims)

		if claims["username"]!= "Admin"{
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Set("username", claims["username"])
		c.Next()
	}
}

func Authenticate(c *gin.Context){
	var user struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}
	
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON request body"})
		return
	}

	if user.Username != "Admin"{
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	hashedpwd,err :=hashPassword(user.Password)
	if err!=nil{
		c.JSON(500,gin.H{"error":"internal server error"})
		c.Abort()
		return
	}
	if !comparePasswords(hashedpwd,"NIsh@9741") {
		c.JSON(401, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	token,err := generateToken(user.Username)
	if err!=nil{
		c.JSON(500,gin.H{"error":"internal server error"})
		c.Abort()
		return
	}

	c.JSON(200,gin.H{"token":token})



}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func comparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}