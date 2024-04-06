package sqlite

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/pleuvoir/gamine/helper/helper_os"
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
	normalizePath, err := helper_os.NormalizePath(i.Path)
	if err != nil {
		color.Redln(fmt.Sprintf("sqlite打开失败，路径：%s", i.Path))
		return err
	}
	db, err = gorm.Open(sqlite.Open(normalizePath), &gorm.Config{})
	if err != nil {
		color.Redln(fmt.Sprintf("sqlite打开失败，路径：%s", normalizePath))
		return err
	}
	return nil
}

func Get() *gorm.DB {
	return db
}
