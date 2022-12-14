// Code generated by goctl. DO NOT EDIT.
package types

type OrderCreateReq struct {
	UserId      int64  `json:"user_id"`
	ProductId   int64  `json:"product_id"`
	ProductName string `json:"product_name"`
	Nums        int64  `json:"nums"`
}

type OrderCreateReply struct {
	Result  bool   `json:"result"`
	OrderNo string `json:"order_no"`
	Message string `json:"message"`
}
