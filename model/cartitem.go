package model

type CartItem struct {
	CartItemID int64 //购物项id
	Book       *Book
	Count      int64   //某个图书数量
	Amount     float64 //某个图书总价
	CartID     string  //购物车id
}

//获取图书金额小记
func (cartItem *CartItem) GetAmount() float64 {
	price := cartItem.Book.Price
	return float64(cartItem.Count) * price
}
