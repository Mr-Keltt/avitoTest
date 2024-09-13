package comment_models

// CommentModel represents a comment returned to the user
type CommentModel struct {
	ID                int
	UserID            int
	OrganizationID    int
	CompanyName       string
	TenderName        string
	TenderDescription string
	BidDescription    string
	ServiceType       string
	Content           string
	CreatedAt         string
}
