package request

type SubjectCreateRequest struct {
	Title        string `json:"title" form:"title" validate:"required,max=32"`
	Description  string `json:"description" form:"description" validate:"required,max=255"`
	ProjectID    uint64 `json:"projectID" form:"projectID" binding:"required"`
	TeamLeaderID uint64 `json:"-" form:"teamLeaderID" binding:"required"`
}
