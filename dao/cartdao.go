package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

//插入购物车
func AddCart(cart *model.Cart) error {
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	//获取购物车中的所有购物项
	cartItems := cart.CartItems
	//遍历的到每一个购物项
	for _, cart := range cartItems {
		//将购物项插入数据库
		AddCartItem(cart)
	}

	return nil
}

// 根据用户id获取购物车内容
func GetCartByUserID(userID int) (*model.Cart, error) {
	sqlStr := "select * from carts where user_id = ?"
	row := utils.Db.QueryRow(sqlStr, userID)
	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}
	//查购物项
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	cart.CartItems = cartItems

	return cart, nil

}

//更新购物车中总数量总金额
func UpdateCart(cart *model.Cart) error {
	sqlStr := "update carts set total_count= ? ,total_amount = ?  where id = ?"
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}

//删除购物车
func DeleteCartByCartID(cartID string) error {
	DeleteCartItemByCartID(cartID)
	sqlStr := "delete from carts where  id = ? "
	_, err := utils.Db.Exec(sqlStr, cartID)
	if err != nil {
		fmt.Printf("删除出错cart: %v\n", err)
	}
	return nil

}
