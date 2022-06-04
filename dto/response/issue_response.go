package response

import "time"

type IssueResponse struct {
	Title           string              `json:"title"`
	Description     string              `json:"description"`
	IssueForeignId  string              `json:"issueForeignId"`
	TargetTime      uint32              `json:"targetTime"`
	Status          uint8               `json:"status"`
	SubjectID       uint64              `json:"subjectID"`
	CreatorID       uint64              `json:"creatorID"`
	AssignieID      *uint64             `json:"assignieID"`
	ChildIssues     []LeafIssueResponse `json:"childIssues"`
	DependentIssues []LeafIssueResponse `json:"dependentIssues"`
	CreatedAt       time.Time           `json:"createdAt"`
	UpdatedAt       time.Time           `json:"updatedAt"`
}

type LeafIssueResponse struct {
	Title           string              `json:"title"`
	Description     string              `json:"description"`
	IssueForeignId  string              `json:"issueForeignId"`
	TargetTime      uint32              `json:"targetTime"`
	Status          uint8               `json:"status"`
	SubjectID       uint64              `json:"subjectID"`
	CreatorID       uint64              `json:"creatorID"`
	AssignieID      *uint64             `json:"assignieID"`
	ChildIssues     []LeafIssueResponse `json:"childIssues"`
	DependentIssues []LeafIssueResponse `json:"dependentIssues"`
	CreatedAt       time.Time           `json:"createdAt"`
	UpdatedAt       time.Time           `json:"updatedAt"`
}
