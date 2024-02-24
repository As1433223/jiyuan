package global

import (
	"github.com/hashicorp/consul/api"
	"github.com/olivere/elastic/v7"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

func init() {
	MysqlInit()
	ConsulInit()
	EsInit()
}
func MysqlInit() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/xin?charset=utf8mb4&parseTime=True&loc=Local"
	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println("mysql连接失败", err)
		return
	}
	log.Println("mysql连接成功")
}
func EsInit() {
	var err error
	dsn := "http://127.0.0.1:9277"
	EsClient, err = elastic.NewClient(elastic.SetURL(dsn), elastic.SetSniff(false))
	if err != nil {
		log.Println("es连接失败", err)
		return
	}
	log.Println("es连接成功")
}
func ConsulInit() {
	consulConf := api.DefaultConfig()
	consulClient, err := api.NewClient(consulConf)
	if err != nil {
		panic(err)
	}
	service, _, err := consulClient.Health().Service("shang", "", true, nil)
	if err != nil {
		return
	}
	if len(service) > 0 {
		for _, entry := range service {
			if entry.Service.ID == "shang" {
				return
			}
		}
	}
	register := &api.AgentServiceRegistration{
		ID:      "shang",
		Name:    "shang",
		Tags:    []string{"xin"},
		Port:    5566,
		Address: "127.0.0.1",
		Check: &api.AgentServiceCheck{
			Interval: "3s",
			GRPC:     "127.0.0.1:5566",
		},
	}
	consulClient.Agent().ServiceRegister(register)
}
