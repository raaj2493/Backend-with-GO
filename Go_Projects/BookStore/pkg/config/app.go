package app

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var (
	DB *gorm.DB
)

func Connect(){
	database , error := gorm.Open("mysql" , "raajgiri:raajgiri@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if error != nil {
		panic(error)
	}
	DB = database
}

func GetDB() *gorm.DB {
	return DB
}