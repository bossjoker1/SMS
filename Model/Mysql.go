package Model

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"io/ioutil"
)

type Mysql struct{
	Dbname string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var Db *gorm.DB

func (m *Mysql)load(path string) bool{
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil{
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &m)
	if err != nil {
		fmt.Println(err.Error())
	}
	return true
}

func init()  {
	var (
		m Mysql
		err error
	)
	m.load("D:\\go\\src\\SMS\\config.yaml")
	// 连接数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=10ms",
		m.Username, m.Password, m.Host, m.Port, m.Dbname)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用负数
		},
	})
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
	if Db.Error != nil{
		fmt.Printf("database error: %v\n", Db.Error)
	}
}
