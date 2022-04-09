package service

import (
	"context"
	pb "emo_ana/proto"
	"emo_ana/server/data/domain"
	"emo_ana/server/model"
	"emo_ana/server/utils"
	"errors"
	"log"
)

func UserLogin(ctx context.Context, req *pb.UserRequest) (*pb.UserDetailResponse, error) {

	user := req.UserName
	rsp := pb.UserDetailResponse{}
	userInfo, err := model.SelectUser(user)
	if err != nil {
		log.Fatalf("select user error:%v", err)
		rsp.Code = -1
		rsp.Msg = "用户验证错误"
		return &rsp, err
	}
	if userInfo == nil {
		err := errors.New("用户不存在")
		rsp.Code = -1
		rsp.Msg = "用户不存在"
		return &rsp, err
	}

	if !utils.CheckPassword(req.Password, userInfo.Password) {
		err := errors.New("密码错误")
		rsp.Code = -1
		rsp.Msg = "密码错误"
		return &rsp, err
	}
	rsp.Code = 0
	rsp.Msg = "登录成功"
	return &rsp, nil
}

func UserRegister(ctx context.Context, req *pb.UserRequest) (*pb.UserDetailResponse, error) {

	user := req.UserName
	rsp := pb.UserDetailResponse{}
	if req.Password != req.PasswordConfirm {
		rsp.Code = -1
		rsp.Msg = "两次密码输入不一致"
		err := errors.New("两次密码输入不一致")
		return &rsp, err
	}
	userDetail, err := model.SelectUser(user)
	if err != nil {
		log.Fatalf("select user error:%v", err)
		rsp.Code = -1
		rsp.Msg = "用户验证错误"
		return &rsp, err
	}
	if userDetail != nil {
		err := errors.New("用户已存在")
		rsp.Code = -1
		rsp.Msg = "用户已存在"
		return &rsp, err
	}

	password, err := utils.SetPassword(req.Password)
	if err != nil {
		log.Fatalf("ser password error:%v", err)
		rsp.Code = -1
		rsp.Msg = "创建密码错误"
		return &rsp, err
	}

	userInfo := domain.TblUserInfo{
		UserName: user,
		Password: password,
	}

	err = model.InsertUserInfo(userInfo)
	if err != nil {
		log.Fatalf("ser user error:%v", err)
		rsp.Code = -1
		rsp.Msg = "创建用户错误"
		return &rsp, err
	}
	rsp.Code = 0
	rsp.Msg = "创建用户成功"
	return &rsp, nil
}
