package sqlite

import (
	"fmt"
	"github.com/gookit/color"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const name = "sqlite"

var db *gorm.DB

type Instance struct {
	Path string `yaml:"path"`
}

func (i *Instance) GetName() string {
	return name
}

func (i *Instance) Run() error {
	var err error
	db, err = gorm.Open(sqlite.Open(i.Path), &gorm.Config{})
	if err != nil {
		color.Redln(fmt.Sprintf("sqlite打开失败，路径：%s", i.Path))
		return err
	}
	return nil
}

func Get() *gorm.DB {
	return db
}
