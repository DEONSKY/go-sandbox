package request

type ProjectCreateRequest struct {
	Title           string `json:"title" form:"title" validate:"required,max=32"`
	Description     string `json:"description" form:"description" validate:"required,max=255"`
	ProjectLeaderID uint64 `json:"-" form:"projectID"`
}
