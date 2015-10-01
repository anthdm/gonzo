package main

import (
	"github.com/dropon/gonzi/database"
	"github.com/dropon/gonzi/model"
)

func main() {
	repo, err := database.Connect()
	if err != nil {
		panic(err)
	}
	repo.LogMode(true)

	for _, user := range users {
		user.SetPassword(user.Password)
		repo.Create(&user)
	}

}

var users = []model.User{
	{Email: "anthonydm@bubblepost.be", Password: "bubble07", Role: "admin"},
	{Email: "andrew@bubblepost.be", Password: "bubble07", Role: "admin"},
	{Email: "gertjan@bubblepost.be", Password: "bubble07", Role: "admin"},
	{Email: "michael@bubblepost.be", Password: "bubble07", Role: "admin"},
	{Email: "thony@bubblepost.be", Password: "bubble07", Role: "admin"},
}
