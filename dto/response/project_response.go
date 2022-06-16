package response

type ProjectNavTreeResponse struct {
	ID          uint64                   `json:"id"`
	Title       string                   `json:"title"`
	Description string                   `json:"description"`
	Subjects    []SubjectNavTreeResponse `gorm:"foreignkey:ProjectID;" json:"subjects"`
}
