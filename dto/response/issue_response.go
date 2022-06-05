package response

import (
	"time"
)

type IssueResponse struct {
	ID             uint64 `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	IssueForeignId string `json:"issueForeignId"`
	TargetTime     uint32 `json:"targetTime"`
	Status         uint8  `json:"status"`
	//SubjectID       uint64              `json:"subjectID"`
	CreatorID     uint64              `json:"creatorID"`
	AssignieID    *uint64             `json:"assignieID"`
	ParentIssueID *uint64             `json:"parentIssueId"`
	ChildIssues   []LeafIssueResponse `gorm:"foreignkey:ParentIssueID;" json:"issues"`
	//DependentIssues []LeafIssueResponse `json:"dependentIssues"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type LeafIssueResponse struct {
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	IssueForeignId string  `json:"issueForeignId"`
	TargetTime     uint32  `json:"targetTime"`
	Status         uint8   `json:"status"`
	ParentIssueID  *uint64 `json:"parentIssueId"`
	//SubjectID       uint64              `json:"subjectID"`
	CreatorID  uint64    `json:"creatorID"`
	AssignieID *uint64   `json:"assignieID"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
