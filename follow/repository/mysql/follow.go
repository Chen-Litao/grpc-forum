package mysql

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Follow struct {
	FollowID    int64 `gorm:"primarykey"`
	UserId      int64
	FollowingId int64
	Followed    int64
}

//func (dao *Follow) Create(req *service.FollowRequest) error {
//	var follow Follow
//	var count int64
//	DB.Where("user_id=? && following_id = ?", req.UserID, req.FollowingID).Count(&count)
//	if count != 0 {
//		return errors.New("Follow relation Exist")
//	}
//	follow = Follow{
//		FollowID:    req.FollowID,
//		UserId:      req.UserID,
//		FollowingId: req.FollowingID,
//		Followed:    req.Followed,
//	}
//	if err := DB.Create(&follow).Error; err != nil {
//		util.LogrusObj.Error("Insert User Error:" + err.Error())
//		return err
//	}
//	return nil
//}

type FollowDao struct {
	*gorm.DB
}

// 创建一个可被追踪链路的上下文
func NewFollowDao(ctx context.Context) *FollowDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &FollowDao{NewDBClient(ctx)}
}
func (dao *FollowDao) InsertFollowRelation(userId int64, targetId int64) error {
	// 生成需要插入的关系结构体。
	follow := Follow{
		UserId:      userId,
		FollowingId: targetId,
		Followed:    1,
	}
	err := dao.Model(&Follow{}).Create(&follow).Error
	if nil != err {
		log.Println(err.Error())
		return err
	}
	fmt.Println("写入成功", err)
	return nil
}

// UpdateFollowRelation 给定用户和目标用户的id，更新他们的关系为取消关注或再次关注。
func (dao *FollowDao) UpdateFollowRelation(userId int64, targetId int64, followed int8) error {
	followobj := new(Follow)
	// 更新用户与目标用户的关注记录（正在关注或者取消关注）
	err := dao.Model(&Follow{}).
		Where("user_id = ?", userId).
		Where("following_id = ?", targetId).
		Update("followed", followed).Error
	err = dao.Model(&Follow{}).
		Where("user_id = ?", userId).
		Where("following_id = ?", targetId).
		Find(&followobj).Error
	fmt.Println("follow", followobj)
	// 更新失败，返回错误。
	if nil != err {
		// 更新失败，打印错误日志。
		log.Println(err.Error())
		return err
	}
	// 更新成功。
	return nil
}

func (dao *FollowDao) FindEverFollowing(userId int64, targetId int64) (*Follow, error) {
	follow := Follow{}
	err := dao.Model(&Follow{}).
		Where("user_id = ? AND following_id = ?", userId, targetId).
		Take(&follow).Error
	if nil != err {
		// 当没查到记录报错时，不当做错误处理。
		if "record not found" == err.Error() {
			return nil, nil
		}
		log.Println(err.Error())
		return nil, err
	}
	// 正常情况，返回取到的关系和空err.
	return &follow, nil
}
func (dao *FollowDao) GetFollowersInfo(userId int64) ([]int64, int64, error) {

	var followerCnt int64
	var followerId []int64

	// following_id -> user_id
	result := dao.Model(&Follow{}).Where("following_id = ?", userId).Where("followed = ?", 1).Pluck("user_id", &followerId)
	followerCnt = result.RowsAffected

	if nil != result.Error {
		log.Println(result.Error.Error())
		return nil, 0, result.Error
	}

	return followerId, followerCnt, nil
}

// GetFollowingsInfo 返回当前用户正在关注的用户信息列表，包括当前用户正在关注的用户ID列表和正在关注的用户总数
func (dao *FollowDao) GetFollowingsInfo(userId int64) ([]int64, int64, error) {

	var followingCnt int64
	var followingId []int64

	// user_id -> following_id
	result := dao.Model(&Follow{}).Where("user_id = ? AND followed = ?", userId, 1).Find(&followingId)
	followingCnt = result.RowsAffected

	if nil != result.Error {
		log.Println(result.Error.Error())
		return nil, 0, result.Error
	}

	return followingId, followingCnt, nil

}
