//订单项
package model

type OrderItem struct {
	OrderItemID int64
	Count       int64
	Amount      float64
	Title       string //图书书名
	Author      string //作者
	Price       float64
	ImgPath     string
	OrderID     string
}
