package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Username string
	Password string
}

var users []User

func main() {
	fmt.Println("欢迎使用用户管理系统")

	for {
		fmt.Println("请选择操作：")
		fmt.Println("1. 注册新用户")
		fmt.Println("2. 用户登录")
		fmt.Println("3. 修改密码")
		fmt.Println("4. 退出")

		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			registerUser()
		case "2":
			login()
		case "3":
			changePassword()
		case "4":
			fmt.Println("再见！")
			return
		default:
			fmt.Println("无效的选择")
		}
	}
}

func registerUser() {
	fmt.Println("注册新用户")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入用户名: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	if isUsernameTaken(username) {
		fmt.Println("用户名已存在")
		return
	}

	fmt.Print("请输入密码: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	user := User{
		Username: username,
		Password: password,
	}

	users = append(users, user)
	fmt.Println("注册成功")
}

func login() {
	fmt.Println("用户登录")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入用户名: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("请输入密码: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	for _, user := range users {
		if user.Username == username && user.Password == password {
			fmt.Println("登录成功")
			return
		}
	}

	fmt.Println("用户名或密码错误")
}

func changePassword() {
	fmt.Println("修改密码")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("请输入用户名: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	fmt.Print("请输入旧密码: ")
	oldPassword, _ := reader.ReadString('\n')
	oldPassword = strings.TrimSpace(oldPassword)

	fmt.Print("请输入新密码: ")
	newPassword, _ := reader.ReadString('\n')
	newPassword = strings.TrimSpace(newPassword)

	for i, user := range users {
		if user.Username == username && user.Password == oldPassword {
			users[i].Password = newPassword
			fmt.Println("密码修改成功")
			return
		}
	}

	fmt.Println("用户名或密码错误")
}

func isUsernameTaken(username string) bool {
	for _, user := range users {
		if user.Username == username {
			return true
		}
	}
	return false
}
