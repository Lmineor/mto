package global

import (
	"github.com/Lmineor/mto/config"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	Config *config.Server
)
