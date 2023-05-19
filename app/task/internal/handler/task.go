package handler

import (
	"context"
	"sync"

	"github.com/CocaineCong/grpc-todolist/app/task/internal/repository/db/dao"
	taskPb "github.com/CocaineCong/grpc-todolist/idl/task/pb"
	"github.com/CocaineCong/grpc-todolist/pkg/e"
)

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

type TaskSrv struct {
	taskPb.UnimplementedTaskServiceServer
}

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}
func (*TaskSrv) TaskCreate(ctx context.Context, req *taskPb.TaskRequest) (resp *taskPb.CommonResponse, err error) {
	resp = new(taskPb.CommonResponse)
	resp.Code = e.SUCCESS
	err = dao.NewTaskDao(ctx).CreateTask(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return
}

func (*TaskSrv) TaskShow(ctx context.Context, req *taskPb.TaskRequest) (resp *taskPb.TasksDetailResponse, err error) {
	r, err := dao.NewTaskDao(ctx).ListTaskByUserId(req.UserID)
	resp.Code = e.SUCCESS
	if err != nil {
		resp.Code = e.ERROR
		return
	}
	for i := range r {
		resp.TaskDetail = append(resp.TaskDetail, &taskPb.TaskModel{
			TaskID:    r[i].TaskID,
			UserID:    r[i].UserID,
			Status:    int64(r[i].Status),
			Title:     r[i].Title,
			Content:   r[i].Content,
			StartTime: r[i].StartTime,
			EndTime:   r[i].EndTime,
		})
	}
	resp.TaskDetail = repository.BuildTasks(tRep)
	return
}

func (*TaskService) TaskUpdate(ctx context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	var task repository.Task
	resp = new(service.CommonResponse)
	resp.Code = e.SUCCESS
	err = task.Update(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return resp, err
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return resp, nil
}

func (*TaskService) TaskDelete(ctx context.Context, req *service.TaskRequest) (resp *service.CommonResponse, err error) {
	var task repository.Task
	resp = new(service.CommonResponse)
	resp.Code = e.SUCCESS
	err = task.Delete(req)
	if err != nil {
		resp.Code = e.ERROR
		resp.Msg = e.GetMsg(e.ERROR)
		resp.Data = err.Error()
		return resp, err
	}
	resp.Msg = e.GetMsg(uint(resp.Code))
	return resp, nil
}
