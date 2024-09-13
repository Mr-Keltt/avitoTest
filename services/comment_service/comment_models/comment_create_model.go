package comment_models

// CommentCreateModel represents the data needed to create a comment
type CommentCreateModel struct {
	UserID            int
	OrganizationID    int
	BidID             int
	CompanyName       string
	TenderName        string
	TenderDescription string
	BidDescription    string
	ServiceType       string
	Content           string
}
