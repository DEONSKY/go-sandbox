package request

type ProjectCreateRequest struct {
	Title           string `json:"title" form:"title" binding:"required"`
	Description     string `json:"description" form:"description" binding:"required"`
	ProjectLeaderID uint64 `json:"-" form:"projectID" binding:"required"`
}
