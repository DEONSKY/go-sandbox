package request

type SubjectCreateRequest struct {
	Title        string `json:"title" form:"title" binding:"required"`
	Description  string `json:"description" form:"description" binding:"required"`
	ProjectID    uint64 `json:"projectID" form:"projectID" binding:"required"`
	TeamLeaderID uint64 `json:"-" form:"teamLeaderID" binding:"required"`
}
