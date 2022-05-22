package models

import "github.com/Lmineor/mto/global"

type TangPoem struct {
	global.CommonModel
	Title      string `json:"title" gorm:"type:varchar(255);comment:诗名"`
	Paragraphs string `json:"paragraphs" gorm:"type:text;comment:诗内容"`
	Author     string `json:"author" gorm:"column:author;comment:诗作者"`
}

func (TangPoem) TableName() string {
	return "tang_poems"
}

type SongPoem struct {
	global.CommonModel
	Title      string `json:"title" gorm:"type:varchar(255);comment:诗名"`
	Paragraphs string `json:"paragraphs" gorm:"type:text;comment:诗内容"`
	Author     string `json:"author" gorm:"column:author;comment:诗作者"`
}

func (SongPoem) TableName() string {
	return "song_poems"
}
