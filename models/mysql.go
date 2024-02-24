package models

func GetUserInfo(user User) (User, error) {
	var NewUser User
	res := db.Where(user).First(&NewUser)
	return NewUser, res.Error
}
func CreateGoods(goods Goods) (Goods, error) {
	res := db.Create(&goods)
	return goods, res.Error
}
