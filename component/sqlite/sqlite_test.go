package sqlite

import (
	"github.com/pleuvoir/gamine"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"testing"
)

type account struct {
	gorm.Model
	Name string
}

func TestInstallComponents(t *testing.T) {
	dbPath, _ := filepath.Abs("../../test/test_data/sqlite-yml.db")
	t.Log(dbPath)

	defer os.Remove(dbPath)

	gamine.SetWorkDir("../../test/")
	gamine.InstallComponents(&Instance{})
	err := Get().AutoMigrate(&account{})
	assert.NoError(t, err)
	acc := account{Name: "gamine"}
	Get().Create(&acc)
	acc2 := account{}
	Get().First(&acc2, "name=?", acc.Name)
	assert.Equal(t, acc.Name, acc2.Name)
}

func TestInstance_Run(t *testing.T) {
	dbPath, _ := filepath.Abs("../../test/test_data/sqlite-test.db")
	t.Log(dbPath)

	defer os.Remove(dbPath)
	gamine.RunComponents(&Instance{Path: dbPath})
	err := Get().AutoMigrate(&account{})
	assert.NoError(t, err)
	acc := account{Name: "gamine"}
	Get().Create(&acc)
	acc2 := account{}
	Get().First(&acc2, "name=?", acc.Name)
	assert.Equal(t, acc.Name, acc2.Name)
}
