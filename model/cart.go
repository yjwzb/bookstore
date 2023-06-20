package model

type Cart struct {
	CartID      string      //uuid
	CartItems   []*CartItem //购物车中所有的购物项
	TotalCount  int64       //所有图书总数量
	TotalAmount float64     //总金额
	UserID      int         //购物车要绑定用户
	UserName    string
}

//获取总数量
func (cart *Cart) GetTotalCount() int64 {
	//遍历购物项切片
	var totalCount int64
	for _, v := range cart.CartItems {
		totalCount = totalCount + v.Count
	}
	return totalCount

}

//获取总价
func (cart *Cart) GetTotalAmount() float64 {

	//遍历购物项切片
	var totalAmount float64
	for _, v := range cart.CartItems {
		totalAmount = totalAmount + v.GetAmount()
	}
	return totalAmount
}
