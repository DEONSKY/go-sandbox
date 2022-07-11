package request

type IssueCommentCreateRequest struct {
	Context   string `json:"context"`
	IssueID   uint64 `json:"issueID"`
	CreatorID uint64 `json:"-"`
}
