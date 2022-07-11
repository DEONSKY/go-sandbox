package request

/*
//BookUpdateDTO is a model that client use when updating a book
type BookUpdateDTO struct {
	ID          uint64 `json:"id" form:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint64 `json:"user_id,omitempty"  form:"user_id,omitempty"`
}
*/

//BookCreateDTO is a model that clinet use when create a new book
type IssueCreateRequest struct {
	Title          string  `json:"title" form:"title" validate:"required,max=32"`
	Description    string  `json:"description" form:"description" validate:"required,max=255"`
	IssueForeignId string  `json:"issueForeignId"`
	TargetTime     uint32  `json:"targetTime"`
	Status         uint8   `json:"status" validate:"numeric"`
	SubjectID      uint64  `json:"subjectID" form:"subjectID" validate:"required"`
	ReporterID     uint64  `json:"reporterID,omitempty"  form:"reporterID,omitempty"`
	AssignieID     *uint64 `json:"assignieID,omitempty"  form:"assignieID,omitempty"`
	ParentIssueID  *uint64 `json:"parentIssueID,omitempty"  form:"parentIssueID,omitempty"`
}

type IssueGetQuery struct {
	SubjectID      *uint64 `query:"subjectID"`
	ProjectID      *uint64 `query:"projectID"`
	ReporterID     *uint64 `query:"reporterID"`
	AssignieID     *uint64 `query:"assignieID"`
	Status         *uint8  `query:"status"`
	ParentIssueID  *uint64 `query:"parentIssueID"`
	GetOnlyOrphans *bool   `query:"getOnlyOrphans"`
}
