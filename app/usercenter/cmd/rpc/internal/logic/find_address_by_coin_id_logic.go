package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"mscoin/app/usercenter/model"
	"mscoin/common/xerr"

	"mscoin/app/usercenter/cmd/rpc/internal/svc"
	"mscoin/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAddressByCoinIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAddressByCoinIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAddressByCoinIdLogic {
	return &FindAddressByCoinIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindAddressByCoinIdLogic) FindAddressByCoinId(in *pb.FindAddressByCoinIdReq) (*pb.FindAddressByCoinIdResp, error) {
	// 获取MemberAddressList
	var list []*model.MemberAddress
	err := l.svcCtx.Db.Table(model.MemberAddressTableName).Select("remark", "address").Where("member_id = ? and coin_id = ?", in.UserId, in.CoinId).Find(&list).Error
	if err != nil {
		l.Logger.Errorf("get MemberAddressList err. err=%+v", err)
		return nil, xerr.NewCommonErrorByCode(xerr.ServerError)
	}

	// 复制内容到结果集中
	var addressList []*pb.AddressSimple
	err = copier.Copy(&addressList, list)
	if err != nil {
		l.Logger.Errorf("copy MemberAddressList (Address,Remark) to AddressList err. err=%+v", err)
		return nil, xerr.NewCommonErrorByCode(xerr.ServerError)
	}

	return &pb.FindAddressByCoinIdResp{List: addressList}, nil
}
