package service

type CountTagRequest struct {
	Name  string `from:"name" binding:"max=100"`
	State uint8  `from:"state,default=1" binding:"oneof=01"`
}

type TagListRequest struct {
	Name  string `from:"name" binding:"max=100"`
	State uint8  `from:"state,default=1" binding:"oneof=01"`
}

type CreateTagRequset struct {
	Name      string `from:"name" binding:"required,min=3,max=100"`
	CreatedBy string `from:"create_by" binding:"required,min=3,max=100"`
	State     uint8  `from:"state,default=1 binding:"oneof=01"`
}

type UpdateTagRequest struct {
	ID         uint32 `from:"id" binding:"required,gte=1"`
	Name       string `from:"name" binding:"min=3,max=100"`
	State      uint8  `from:"state" binding:"required,oneof=01"`
	ModifiedBy string `from:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `from:"id" binding:"required,gte=1"`
}
