package helper

import (
	"github.com/danilopolani/gocialite"
	"github.com/jinzhu/gorm"
)

var GormMasterDB *gorm.DB

var (
	Gocial = gocialite.NewDispatcher()
)
