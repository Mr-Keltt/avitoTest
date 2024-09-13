package comment_models

// CommentModel represents a comment returned to the user
type CommentModel struct {
	ID                int    `json:"id"`
	UserID            int    `json:"user_id"`
	OrganizationID    int    `json:"organization_id"`
	CompanyName       string `json:"company_name"`
	TenderName        string `json:"tender_name"`
	TenderDescription string `json:"tender_description"`
	BidDescription    string `json:"bid_description"`
	ServiceType       string `json:"service_type"`
	Content           string `json:"content"`
	CreatedAt         string `json:"created_at"`
}
