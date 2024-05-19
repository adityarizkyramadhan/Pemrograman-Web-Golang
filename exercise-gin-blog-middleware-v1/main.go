package main

import (
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var Posts = []Post{
	{ID: 1, Title: "Judul Postingan Pertama", Content: "Ini adalah postingan pertama di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Title: "Judul Postingan Kedua", Content: "Ini adalah postingan kedua di blog ini.", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{Username: "user1", Password: "pass1"},
	{Username: "user2", Password: "pass2"},
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.JSON(401, gin.H{"error": "Token not found"})
			c.Abort()
			return
		}
		// decode base64
		// split by : " "
		// check if user exists
		// check if password is correct
		// if not, return 401
		// if yes, continue

		decodeData := strings.Split(c.GetHeader("Authorization"), " ")[1]
		// decode base64
		decoded, err := base64.StdEncoding.DecodeString(decodeData)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// split by :
		splitData := strings.Split(string(decoded), ":")
		if len(splitData) != 2 {
			c.JSON(401, gin.H{"message": "kombinasi username dan password tidak valid"})
			c.Abort()
			return
		}

		username := splitData[0]
		password := splitData[1]

		// check if user exists
		var user User
		for _, u := range users {
			if u.Username == username {
				user = u
				break
			}
		}

		if user.Username == "" {
			c.JSON(401, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		// check if password is correct
		if user.Password != password {
			c.JSON(401, gin.H{"message": "password salah"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//Set up authentication middleware here // TODO: replace this
	r.Use(authMiddleware())

	r.GET("/posts", func(c *gin.Context) {
		idQuery := c.Query("id")

		if idQuery != "" {
			idQueryInt, err := strconv.Atoi(idQuery)
			if err != nil {
				c.JSON(400, gin.H{"error": "ID harus berupa angka"})
				return
			}
			for _, post := range Posts {
				if idQueryInt == post.ID {
					c.JSON(200, gin.H{"post": post})
					return
				}
			}
			c.JSON(404, gin.H{"error": "Postingan tidak ditemukan"})
			return
		}
		c.JSON(200, gin.H{"posts": Posts})
	})
	r.POST("/posts", func(c *gin.Context) {
		var post Post
		err := c.BindJSON(&post)
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid request body"})
			return
		}
		post.ID = len(Posts) + 1
		post.CreatedAt = time.Now()
		post.UpdatedAt = time.Now()
		Posts = append(Posts, post)

		c.JSON(201, gin.H{
			"message": "Postingan berhasil dibuat",
			"post":    post,
		})
	})

	return r
}

func main() {
	r := SetupRouter()

	r.Run(":8080")
}
