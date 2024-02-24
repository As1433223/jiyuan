package logic

import (
	"context"
	"encoding/json"
	"github.com/As1433223/jiyuan/es"
	"github.com/As1433223/jiyuan/proto"
	"github.com/olivere/elastic/v7"
)

func (s *ServerRpc) EsShow(ctx context.Context, in *proto.EsShowReq) (*proto.EsShowRes, error) {
	var rr proto.EsShowRes
	rr.Data = nil
	high := elastic.NewHighlight()
	text := elastic.NewHighlighterField("goods_name")
	text.PreTags("<i>").PostTags("</i>").NumOfFragments(1)
	high.Fields(text)
	query := elastic.NewMultiMatchQuery(in.Word, "goods_name", "title", "goods_comment")
	res, err := es.EsShow(in.Index, query, int(in.Page), int(in.Size), high)
	if err != nil {
		return &rr, err
	}
	data, _ := json.Marshal(res)
	rr.Data = data
	return &rr, nil
}
