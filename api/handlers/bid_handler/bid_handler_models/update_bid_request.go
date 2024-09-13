package bid_handler_models

type UpdateBidRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
