package components

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"lianghj.top/admin/goadminv1/models"
)

func InitViper() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func InitGormMysql() *gorm.DB {
	username := viper.GetString("databases.username")
	password := viper.GetString("databases.password")
	host := viper.GetString("databases.host")
	port := viper.GetString("databases.port")
	database := viper.GetString("databases.database")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
	)
	fmt.Println("dsn:  ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database err: ", err)
		return nil
	}
	return db

}
func InitAll() {
	InitViper()
	models.DB = InitGormMysql()
}
