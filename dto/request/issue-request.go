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
	Title          string  `json:"title" form:"title" binding:"required"`
	Description    string  `json:"description" form:"description" binding:"required"`
	IssueForeignId string  `json:"issueForeignId"`
	TargetTime     uint32  `json:"targetTime"`
	Status         uint8   `json:"status"`
	SubjectID      uint64  `json:"subjectID" form:"subjectID" binding:"required"`
	CreatorID      uint64  `json:"creatorID,omitempty"  form:"creatorID,omitempty`
	AssignieID     *uint64 `json:"assignieID,omitempty"  form:"assignieID,omitempty`
}

type IssueGetQuery struct {
	UserID        *uint64 `query:"userID"`
	SubjectID     *uint64 `query:"subjectID"`
	ProjectID     *uint64 `query:"projectID"`
	CreatorID     *uint64 `query:"creatorID"`
	AssignieID    *uint64 `query:"assignieID"`
	Status        *uint8  `query:"status"`
	ParentIssueID *uint64 `query:"parentIssueID"`
}
