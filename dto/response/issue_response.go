package response

import (
	"time"
)

type IssueKanbanResponse struct {
	Status StatusResponse  `json:"status"`
	Issues []IssueResponse `json:"issues"`
}
type IssueResponse struct {
	ID              uint64               `json:"id"`
	Title           string               `json:"title"`
	Description     string               `json:"description"`
	IssueForeignId  string               `json:"issueForeignID"`
	TargetTime      uint32               `json:"targetTime"`
	StatusID        uint32               `json:"statusID"`
	Status          StatusResponse       `gorm:"-" json:"status"`
	SubjectID       uint64               `json:"subjectID"`
	ReporterID      uint64               `json:"reporterID"`
	AssignieID      *uint64              `json:"assignieID"`
	ParentIssueID   *uint64              `json:"parentIssueID"`
	ChildIssues     []*LeafIssueResponse `gorm:"foreignkey:ParentIssueID;" json:"issues"`
	DependentIssues []*LeafIssueResponse `gorm:"many2many:DependentIssues;foreignkey:ID;joinForeignKey:issueID;References:ID;joinReferences:dependentIssueID" json:"dependentIssues"`
	CreatedAt       time.Time            `json:"createdAt"`
	UpdatedAt       time.Time            `json:"updatedAt"`
}

type LeafIssueResponse struct {
	ID             uint64         `json:"id"`
	Title          string         `json:"title"`
	Description    string         `json:"description"`
	IssueForeignId string         `json:"issueForeignID"`
	TargetTime     uint32         `json:"targetTime"`
	StatusID       uint32         `json:"statusID"`
	Status         StatusResponse `gorm:"-" json:"status"`
	ParentIssueID  *uint64        `json:"parentIssueID"`
	SubjectID      uint64         `json:"subjectID"`
	ReporterID     uint64         `json:"reporterID"`
	AssignieID     *uint64        `json:"assignieID"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
}

// TableName overrides the table name for smart select
func (LeafIssueResponse) TableName() string {
	return "issues"
}

type StatusResponse struct {
	ID      uint32
	Title   string
	HexCode string
}
