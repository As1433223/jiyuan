package logic

import (
	"context"
	"github.com/As1433223/jiyuan/es"
	"github.com/As1433223/jiyuan/models"
	"github.com/As1433223/jiyuan/proto"
)

func (s *ServerRpc) EsAdd(ctx context.Context, in *proto.EsAddReq) (*proto.EsAddRes, error) {
	var rr proto.EsAddRes
	goods, err := models.CreateGoods(models.Goods{
		GoodsName: in.Name,
		Title:     in.Title,
		Comment:   in.Comment,
	})
	if err != nil {
		return &rr, err
	}
	err = es.EsAdd(in.Index, goods)
	if err != nil {
		return &rr, err
	}
	return &rr, nil
}
