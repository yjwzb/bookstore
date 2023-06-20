package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"fmt"
)

func AddCartItem(cartitem *model.CartItem) error {
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, cartitem.Count, cartitem.GetAmount(), cartitem.Book.ID, cartitem.CartID)
	if err != nil {
		fmt.Printf("添加cartitem信息出错: %v\n", err)
		return err
	}
	return nil
}

//根据bookid查购物项
func GetCartItemByBookIDAndCartID(bookID, cartID string) (*model.CartItem, error) {
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)
	// var cartItem *model.CartItem
	// row.Scan(cartItem.CartItemID, cartItem.Book, cartItem.Count, cartItem.Amount, cartItem.CartID)
	cartItem := &model.CartItem{}

	row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	book, _ := GetBookById(bookID)
	cartItem.Book = book

	return cartItem, nil

}

//根据购物车id查询购物车内所有的购物项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	//cart_items中的项目可能有相同的curt_id
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	rows, _ := utils.Db.Query(sqlStr, cartID)
	// cartItems := []*model.CartItem{}
	// row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	var cartItems []*model.CartItem
	for rows.Next() {
		var bookID string
		cartItem := &model.CartItem{}
		rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		//根据bookid获取图书信息
		book, _ := GetBookById(bookID)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}

//更新数据库中count数目，根据cart_id
func UpdateCartItem(cartItem *model.CartItem) error {

	sqlStr := "update cart_items set count = ?,amount = ? where book_id = ? and cart_id = ?"
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

//根据购物车id删除购物项
func DeleteCartItemByCartID(cartID string) error {
	sqlStr := "delete from cart_items where cart_id = ?"
	_, err := utils.Db.Exec(sqlStr, cartID)
	if err != nil {
		fmt.Println("删除出错cartItem:", err)
		return err
	}
	return nil
}

// //根据图书id删除购物项
// func DeleteCartItemByBookID(bookID string, cartID string, userID int) error {
// 	// book, err := GetBookById(bookID)
// 	// if err != nil {
// 	// 	fmt.Printf("DeleteCartItemByBookID中查询图书出错: %v\n", err)
// 	// 	return err
// 	// }
// 	sqlStr := "delete from cart_items where book_id = ? and cart_id =  ?"
// 	utils.Db.Exec(sqlStr, bookID, cartID)

// 	//删除之后还要把cart中的cartitems切片更新了
// 	//首先根据userid获取cart
// 	cart, err2 := GetCartByUserID(userID)
// 	if err2 != nil {
// 		fmt.Printf("DeleteCartItemByBookID中获取cart出错: %v\n", err2)
// 		return err2
// 	}
// 	cartItems := cart.CartItems
// 	for key, value := range cartItems {
// 		if value.CartID == cartID {
// 			if key == 0 {
// 				cart.CartItems = append(cartItems[:0], cartItems[1:]...)
// 			} else if key == len(cartItems)-1 {
// 				cart.CartItems = cartItems[:len(cartItems)-1]
// 			} else {
// 				cart.CartItems = append(cartItems[:key], cartItems[key+1:]...)
// 			}

// 		}
// 	}
// 	err3 := UpdateCart(cart)
// 	if err3 != nil {
// 		fmt.Printf("DeleteCartItemByCartID更新cart出错: %v\n", err3)
// 		return err3
// 	}
// 	return nil

// }
func DeleteCartItemByID(cartItemID string) error {
	sqlStr := "delete from cart_items where id = ?"
	_, err := utils.Db.Exec(sqlStr, cartItemID)
	if err != nil {
		fmt.Printf("DeleteCartItemByID中删除出错: %v\n", err)
		return err
	}
	return nil

}
