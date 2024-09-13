package comment_models

// CommentCreateModel represents the data needed to create a comment
type CommentCreateModel struct {
	UserID            int    `json:"user_id"`
	OrganizationID    int    `json:"organization_id"`
	CompanyName       string `json:"company_name"`
	TenderName        string `json:"tender_name"`
	TenderDescription string `json:"tender_description"`
	BidDescription    string `json:"bid_description"`
	ServiceType       string `json:"service_type"`
	Content           string `json:"content"`
}
