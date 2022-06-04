package response

type UserResponse struct {
	Title           string              `json:"title" form:"title" binding:"required"`
	Description     string              `json:"description" form:"description" binding:"required"`
	IssueForeignId  string              `json:"issueForeignId"`
	TargetTime      uint32              `json:"targetTime"`
	Status          uint8               `json:"status"`
	SubjectID       uint64              `json:"subject_id" form:"subject_id" binding:"required"`
	CreatorID       uint64              `json:"creator_id,omitempty"  form:"creator_id,omitempty`
	AssignieID      *uint64             `json:"assignie_id,omitempty"  form:"assignie_id,omitempty`
	ChildIssues     []LeafIssueResponse `gorm:"foreignkey:ParentIssueID;" json:"issue"`
	DependentIssues []LeafIssueResponse `gorm:"many2many:DependentIssues;" json:"dependentIssues"`
}
