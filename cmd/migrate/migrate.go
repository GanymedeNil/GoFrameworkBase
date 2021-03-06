package main

import (
	"github.com/GanymedeNil/GoFrameworkBase/internal/core"
	"github.com/GanymedeNil/GoFrameworkBase/internal/global"
	"github.com/GanymedeNil/GoFrameworkBase/internal/model"

	"golang.org/x/crypto/bcrypt"
)

func init() {
	core.Viper()
	core.Gorm()
}

func main() {

	global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&model.User{}, &model.BlogPost{}, &model.BlogTag{}, &model.BlogCategory{}, &model.BlogPostCategory{}, &model.BlogPostTag{}, &model.BlogPostCategory{})
	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	global.DB.Where(model.User{Name: "demo"}).FirstOrCreate(&model.User{Password: string(password)})
}
