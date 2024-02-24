package global

import (
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

var (
	MysqlDB  *gorm.DB
	EsClient *elastic.Client
)
