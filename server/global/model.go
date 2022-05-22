package global


type CommonModel struct {
	Id string `json:"id" gorm:"type:varchar(64);id;comment:id"`
}
