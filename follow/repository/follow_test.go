package repository

import (
	"fmt"
	"follow/internal/service"
	"testing"
)

func TestUser_Create(t *testing.T) {
	//config.InitConfig()
	InitDB()
	f := new(Follow)
	req := new(service.FollowRequest)
	req.UserID = 10
	req.FollowingID = 11
	req.Followed = 1
	req.FollowID = 1
	err := f.Create(req)
	fmt.Println(err)
}
