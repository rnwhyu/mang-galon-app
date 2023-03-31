package order

type OrderMakeRequest struct {
	//UserID     int `json:"user_id" binding:"required"`
	GalonID    int `json:"galon_id" binding:"required"`
	TotalOrder int `json:"total_order" binding:"required"`
}

type OrderFindReq struct {
	ID string `uri:"id" binding:"required,numeric"`
}
type OrderUpdateReq struct {
	OrderFindReq
	Status string `json:"status" binding:"required"`
}
