package response

type SubjectNavTreeResponse struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ProjectID   uint64 `json:"project_id"`
}
