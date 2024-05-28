package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"mscoin/app/usercenter/cmd/rpc/internal/model"
	"mscoin/common/xerr"

	"mscoin/app/usercenter/cmd/rpc/internal/svc"
	"mscoin/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMemberByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMemberByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMemberByIdLogic {
	return &FindMemberByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindMemberByIdLogic) FindMemberById(in *pb.FindMemberByIdReq) (*pb.FindMemberByIdResp, error) {
	member := new(model.Member)
	err := l.svcCtx.Db.Table(model.MemberTableName).Where("id = ?", in.MemberId).Take(member).Error
	if err == gorm.ErrRecordNotFound {
		return nil, xerr.NewCommonErrorByCode(xerr.FindInfoError)
	} else if err != nil {
		return nil, errors.Wrapf(xerr.NewCommonErrorByCode(xerr.ServerError), "查找Member数据失败, err=%+v", err)
	}

	resp := &pb.FindMemberByIdResp{}
	copier.Copy(resp, member)

	return resp, nil
}
