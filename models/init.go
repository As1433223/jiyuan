package models

import (
	"github.com/As1433223/jiyuan/global"
	"gorm.io/gorm"
	"log"
)

var db = global.MysqlDB

func init() {
	err := db.AutoMigrate(&User{}, &FruitMes{}, &Goods{})
	if err != nil {
		log.Println("数据表创建失败", err)
		return
	}
}

// User todo: 用户表
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(15);not null;comment:账号" json:"username"`
	Password string `gorm:"type:varchar(15);not null;comment:密码" json:"password"`
	Status   int    `gorm:"type:tinyint(1);check:status IN (0,1);not null;comment:状态" json:"status"`
}

type FruitMes struct {
	gorm.Model
	Fid  int    `gorm:"type:int(11);not null;comment:父级id"`
	Name string `gorm:"type:varchar(30);not null;comment:名称"`
}
type Fruit struct {
	gorm.Model
	Fid  int
	Name string
	FAll []*Fruit
}
type Goods struct {
	gorm.Model
	GoodsName string `gorm:"type:varchar(255);not null;comment:商品名称" json:"goods_name"`
	Title     string `gorm:"type:varchar(255);not null;comment:商品介绍" json:"title"`
	Comment   string `gorm:"type:varchar(255);not null;comment:商品详情" json:"comment"`
}

//	type Goods struct {
//		gorm.Model
//		VideoName string `gorm:"type:varchar(30);not null;comment:视频名称" json:"username"`
//		VideoName string `gorm:"type:varchar(30);not null;comment:视频名称" json:"username"`
//		VideoName string `gorm:"type:varchar(30);not null;comment:视频名称" json:"username"`
//	}
func GetFruitInfo() []FruitMes {
	var rr []FruitMes
	db.Find(&rr)
	return rr
}
func GetFruit(fid int) ([]*Fruit, error) {
	//todo: 查询数据库
	var fruitAll []*FruitMes
	err := db.Where("fid", fid).Find(&fruitAll).Error
	if err != nil {
		return nil, err
	}
	//todo: 容器   复制给容器结构体  并且放入切片
	var fruitList []*Fruit
	for _, fruit := range fruitAll {
		//todo： 调用方法赋值
		fruits := F(fruit)
		//todo： 添加入切片
		fruitList = append(fruitList, fruits)
	}

	//todo： 循环容器切片得到单独容器结构体  开始递归
	//todo： 竖向单分支查询  每查一层，就先放入父级切片（赋值），再查他的子集  如果查到底  竖向逐层返回
	for _, fruits := range fruitList {
		cl, err := GetFruit(int(fruits.ID))
		if err != nil {
			return nil, err
		}
		fruits.FAll = cl
	}
	//todo： 全部查询完毕  返回
	return fruitList, nil
}
func F(f *FruitMes) *Fruit {
	return &Fruit{
		Model: gorm.Model{
			ID:        f.ID,
			CreatedAt: f.CreatedAt,
			UpdatedAt: f.UpdatedAt,
			DeletedAt: f.DeletedAt,
		},
		Name: f.Name,
		Fid:  f.Fid,
		FAll: make([]*Fruit, 0),
	}
}
