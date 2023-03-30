package galon

type GalonAddReq struct{
	Brandname string `json:"brand_name" binding:"required"`
	Stock int `json:"stock" binding:"required"`
}
type GalonFindReq struct{
	ID string `uri:"id" binding:"required,numeric"`
}
type GalonUpdateReq struct{
	GalonFindReq
	Stock int `json:"stock" binding:"required"`
	//GalonAddReq
}