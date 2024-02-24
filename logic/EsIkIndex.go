package logic

import (
	"context"
	"github.com/As1433223/jiyuan/es"
	"github.com/As1433223/jiyuan/proto"
)

func (s *ServerRpc) EsIkIndex(ctx context.Context, in *proto.EsIkIndexReq) (*proto.EsIkIndexRes, error) {
	var rr proto.EsIkIndexRes
	err := es.EsIkIndex(in.Index)
	if err != nil {
		return &rr, err
	}
	return &rr, nil
}
