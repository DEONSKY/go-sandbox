package response

import (
	"time"
)

type IssueKanbanResponse struct {
	Status StatusResponse  `json:"status"`
	Issues []IssueResponse `json:"issues"`
}
type IssueResponse struct {
	ID             uint64             `json:"id"`
	Title          string             `json:"title"`
	Description    string             `json:"description"`
	IssueForeignId string             `json:"issueForeignID"`
	TargetTime     uint32             `json:"targetTime"`
	SpendingTime   uint32             `json:"spendingTime"`
	StatusID       uint32             `json:"statusID"`
	Status         StatusResponse     `gorm:"-" json:"status"`
	SubjectID      uint64             `json:"subjectID"`
	ReporterID     *uint64            `json:"ReporterID"`
	Reporter       UserLabelResponse  `json:"reporter"`
	ParentIssueID  *uint64            `json:"parentIssueID"`
	AssignieID     *uint64            `json:"assignieID"`
	Assignie       *UserLabelResponse `json:"assignie"`
	//Comments        []*IssueCommentResponse `gorm:"foreignkey:IssueID;" json:"issueComments"`
	ChildIssues     []*LeafIssueResponse `gorm:"foreignkey:ParentIssueID;" json:"issues"`
	DependentIssues []*LeafIssueResponse `gorm:"many2many:DependentIssues;foreignkey:ID;joinForeignKey:issueID;References:ID;joinReferences:dependentIssueID" json:"dependentIssues"`
	CreatedAt       time.Time            `json:"createdAt"`
	UpdatedAt       time.Time            `json:"updatedAt"`
}

type LeafIssueResponse struct {
	ID             uint64             `json:"id"`
	Title          string             `json:"title"`
	Description    string             `json:"description"`
	IssueForeignId string             `json:"issueForeignID"`
	TargetTime     uint32             `json:"targetTime"`
	SpendingTime   uint32             `json:"spendingTime"`
	StatusID       uint32             `json:"statusID"`
	Status         StatusResponse     `gorm:"-" json:"status"`
	ParentIssueID  *uint64            `json:"parentIssueID"`
	SubjectID      uint64             `json:"subjectID"`
	ReporterID     *uint64            `json:"ReporterID"`
	Reporter       UserLabelResponse  `json:"reporter"`
	AssignieID     *uint64            `json:"assignieID"`
	Assignie       *UserLabelResponse `json:"assignie"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      time.Time          `json:"updatedAt"`
}

// TableName overrides the table name for smart select
func (LeafIssueResponse) TableName() string {
	return "issues"
}

type IssueCommentResponse struct {
	Context   string `json:"context"`
	IssueID   uint64 `json:"issueID"`
	CreatorID uint64 `json:"-"`
}

func (IssueCommentResponse) TableName() string {
	return "issue_comments"
}

type StatusResponse struct {
	ID      uint32
	Title   string
	HexCode string
}

func (UserLabelResponse) TableName() string {
	return "users"
}

type UserLabelResponse struct {
	ID                uint64 `json:"id"`
	Name              string `json:"name"`
	ProfilePictureURL string `json:"profilePictureURL"`
}
