package handler

import (
	"context"
	"follow/internal/service"
	followsrv "follow/internal/service/follow"
	"follow/pkg/e"
	"log"
)

type FollowService struct {
	service.UnimplementedFollowServiceServer
}

func NewFollowService() *FollowService {
	return &FollowService{}
}

func (*FollowService) FollowAction(ctx context.Context, req *service.FollowRequest) (resp *service.FollowDetailResponse, err error) {
	l := followsrv.GetFollowSrv()
	resp = new(service.FollowDetailResponse)
	switch {
	case 1 == req.Type:
		go func() {
			_, err := l.FollowAction(context.Background(), int64(req.UserID), int64(req.FollowingID))
			if err != nil {
				log.Println(err)
			}
		}()
	case 2 == req.Type:
		go func() {
			_, err := l.CancelFollowAction(context.Background(), int64(req.UserID), int64(req.FollowingID))
			if err != nil {
				log.Println(err)
			}
		}()
	}
	resp.UserDetail = &service.FollowModel{
		FollowID:    req.FollowID,
		UserID:      req.UserID,
		FollowingID: req.FollowingID,
		Followed:    req.Followed,
	}
	resp.Code = e.SUCCESS
	return
}
