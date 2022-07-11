package response

type UserResponse struct {
	Title           string              `json:"title" form:"title"`
	Description     string              `json:"description" form:"description""`
	IssueForeignId  string              `json:"issueForeignId"`
	TargetTime      uint32              `json:"targetTime"`
	Status          uint8               `json:"status"`
	SubjectID       uint64              `json:"subject_id" form:"subject_id"`
	CreatorID       uint64              `json:"creator_id,omitempty" `
	AssignieID      *uint64             `json:"assignie_id,omitempty" `
	ChildIssues     []LeafIssueResponse `gorm:"foreignkey:ParentIssueID;" json:"issue"`
	DependentIssues []LeafIssueResponse `gorm:"many2many:DependentIssues;" json:"dependentIssues"`
}

type UserOptionResponse struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}
