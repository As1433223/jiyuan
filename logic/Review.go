package logic

import (
	"context"
	"encoding/json"
	"github.com/As1433223/jiyuan/models"
	"github.com/As1433223/jiyuan/proto"
	"log"
)

func (s *ServerRpc) Review(ctx context.Context, in *proto.ReviewReq) (*proto.ReviewRes, error) {
	var rr proto.ReviewRes
	rr.Data = nil
	res, err := models.GetFruit(0)
	if err != nil {
		return &rr, err
	}
	data, _ := json.Marshal(res)
	rr.Data = data
	log.Println(res, "++++++++++++++++")
	return &rr, nil
}
