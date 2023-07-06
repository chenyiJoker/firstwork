package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

var db *sql.DB

// User represents a user object
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	initDB()
	r := gin.Default()

	// Register endpoints
	r.POST("/register", register)
	r.POST("/login", login)
	r.PUT("/changePassword", changePassword)
	r.DELETE("/deleteAccount", deleteAccount)

	err := r.Run()
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func initDB() {
	var err error
	dsn := "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
}

func register(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sqlStr := "insert into user (username, password) values (?, ?)"

	_, err = db.Exec(sqlStr, user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func login(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStr := "select username, password from user where username = ? and password = ?"
	err = db.QueryRow(sqlStr, user.Username, user.Password).Scan(&user.Username, &user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func changePassword(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStr := "update user set password = ? where username = ?"

	_, err = db.Exec(sqlStr, user.Password, user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

func deleteAccount(c *gin.Context) {
	var user User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sqlStr := "delete from user where username = ? and password = ?"

	_, err = db.Exec(sqlStr, user.Username, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deleted successfully"})
}
