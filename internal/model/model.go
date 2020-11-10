package model

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreateBy   string `json:"created_by`
	ModifiedBy string `json:"modifyied_by`
	CreatedOn  uint32 `json:"created_on`
	ModifyedOn uint32 `json:"modifyed_on`
	DeletedOn  uint32 `json:"deleted_on`
	IsDel      uint8  `json:"is_del`
}
