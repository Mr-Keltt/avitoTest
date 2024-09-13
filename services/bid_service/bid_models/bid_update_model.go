package bid_models

type BidUpdateModel struct {
	ID          int    `json:"id" validate:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
