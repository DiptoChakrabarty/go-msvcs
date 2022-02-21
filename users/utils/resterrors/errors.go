package resterrors

type RestErr struct {
	Message string `json:"msg"`
	Status  int    `json:"status"`
	Error   string `json:err"`
}
