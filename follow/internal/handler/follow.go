package handler

import (
	"context"
	"follow/internal/service"
)

type FollowService struct {
	service.UnimplementedFollowServiceServer
}

func NewFollowService() *FollowService {
	return &FollowService{}
}

func (*FollowService) FollowAction(ctx context.Context, req *service.FollowRequest) (resp *service.FollowDetailResponse, err error) {

}

func (*FollowService) FollowList(ctx context.Context, req *service.FollowRequest) (resp *service.FollowDetailResponse, err error) {

}

func (*FollowService) FollowerList(ctx context.Context, req *service.FollowRequest) (resp *service.FollowDetailResponse, err error) {

}
