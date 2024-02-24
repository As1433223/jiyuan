package es

import (
	"context"
	"github.com/As1433223/jiyuan/global"
	"github.com/As1433223/jiyuan/models"
	"github.com/olivere/elastic/v7"
	"strconv"
)

var EsClient = global.EsClient

func EsIkIndex(index string) error {
	_, err := global.EsClient.CreateIndex(index).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func EsAdd(index string, data models.Goods) error {
	_, err := EsClient.Index().Index(index).Id(strconv.Itoa(int(data.ID))).BodyJson(data).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func EsShow(index string, query elastic.Query, page, size int, high *elastic.Highlight) (*elastic.SearchResult, error) {
	res, err := EsClient.Search(index).Query(query).Highlight(high).From((page - 1) * size).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return res, nil
}
