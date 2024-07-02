package repository

import (
	"errors"
	"follow/internal/service"
	"follow/pkg/util"
)

type Follow struct {
	FollowID    int64 `gorm:"primarykey"`
	UserId      int64
	FollowingId int64
	Followed    int64
}

func (dao *Follow) Create(req *service.FollowRequest) error {
	var follow Follow
	var count int64
	DB.Where("user_id=? && following_id = ?", req.UserID, req.FollowingID).Count(&count)
	if count != 0 {
		return errors.New("Follow relation Exist")
	}
	follow = Follow{
		FollowID:    req.FollowID,
		UserId:      req.UserID,
		FollowingId: req.FollowingID,
		Followed:    req.Followed,
	}
	if err := DB.Create(&follow).Error; err != nil {
		util.LogrusObj.Error("Insert User Error:" + err.Error())
		return err
	}
	return nil
}
