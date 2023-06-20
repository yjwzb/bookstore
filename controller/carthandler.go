package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"bookstore/utils"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	//判断是否登陆
	flag, session := dao.IsLogin(r)
	if flag {
		// /获取图书id
		bookId := r.FormValue("bookId")
		//获取图书信息
		book, _ := dao.GetBookById(bookId)
		userID := session.UserID
		//判断数据库中是否有当前用户的购物车
		cart, _ := dao.GetCartByUserID(userID)

		if cart != nil {

			//说明用户已经有购物车,根据用户id获取购物车，然后添加购物项，如果购物项存在就count+1
			// 首先，根据用户id获取购物车
			cartID := cart.CartID
			//然后判断这个图书是否存在购物项
			cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookId, cartID)

			if cartItem.CartItemID != 0 {
				//说明存在这个图书,则count+1
				cartItems := cart.CartItems
				for _, item := range cartItems {
					// if item.CartItemID == cartItem.CartItemID {
					// 	cartItem.Count++
					// }
					if item.Book.ID == cartItem.Book.ID {
						item.Count++

						//将数据库中的count加1//更新数据库
						dao.UpdateCartItem(item)

					}
				}
				// cartItem.Count++
			} else {

				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cartID,
				}
				//把新的item插入到切片中
				cart.CartItems = append(cart.CartItems, cartItem)
				//将新创建的购物项插入到数据库里边
				dao.AddCartItem(cartItem)
				//数量会变
			}
			dao.UpdateCart(cart)

		} else {
			//说明用户没有购物车，那么就要创建购物车
			uuid := utils.CreateUUID()
			cart := &model.Cart{
				CartID: uuid,
				UserID: userID,
			}
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: uuid,
			}
			cartItems := []*model.CartItem{}
			cartItems = append(cartItems, cartItem)
			cart.CartItems = cartItems
			// dao.AddCartItem(cartItem)
			dao.AddCart(cart)
		}
		w.Write([]byte("您刚刚将" + book.Title + "加入到购物车"))

	} else {
		// 没有登陆,

		w.Write([]byte("请先登陆！"))

	}

}

//获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLogin(r)
	//获取用户的id
	userID := session.UserID
	//根据用户的id从数据库中获取对应的购物车
	cart, _ := dao.GetCartByUserID(userID)
	if cart != nil {
		//将购物车设置到session中
		session.Cart = cart
		//解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		t.Execute(w, session)
	} else {
		//该用户还没有购物车
		//解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		t.Execute(w, session)
	}
}

func DeleteCart(w http.ResponseWriter, r *http.Request) {
	// 获取购物车id
	cartID := r.FormValue("cartId")
	//调用方法
	err := dao.DeleteCartByCartID(cartID)
	if err != nil {
		fmt.Printf("DeleteCart错误: %v\n", err)
	}
	//再次查询购物车信息
	GetCartInfo(w, r)

}

//删除购物项

func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	//转换为int64
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	//调用方法
	dao.DeleteCartItemByID(cartItemID)
	//获取session
	_, session := dao.IsLogin(r)
	//获取用户id
	userID := session.UserID
	//获取该用户的购物车
	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		fmt.Printf("DeleteCartItem处理出错: %v\n", err)
	}
	//获取购物项切片
	cartItems := cart.CartItems
	for key, value := range cartItems {
		if value.CartItemID == iCartItemID {
			//要删除的购物项
			cartItems = append(cartItems[:key], cartItems[key+1:]...)
			cart.CartItems = cartItems //现在的go版本似乎不需要这一步操作
			//将当前的购物项从数据库中移除
			dao.DeleteCartItemByID(cartItemID)

		}
	}

	//更新cart
	dao.UpdateCart(cart)
	//删除之后还是返回查询购物车页面
	GetCartInfo(w, r)
}

//更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemId")
	//转换为int64
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	//获取用户输入的count
	bookCount := r.FormValue("bookCount")
	ibookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	//获取session
	_, session := dao.IsLogin(r)
	//获取用户id
	userID := session.UserID
	//获取该用户的购物车
	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		fmt.Printf("DeleteCartItem处理出错: %v\n", err)
	}
	//获取购物项切片
	cartItems := cart.CartItems
	for _, value := range cartItems {
		if value.CartItemID == iCartItemID {
			//要更新的购物项
			value.Count = ibookCount
			//更新数据库中的该购物项
			dao.UpdateCartItem(value)
		}
	}
	//更新cart
	dao.UpdateCart(cart)
	//还是返回查询购物车页面
	GetCartInfo(w, r)
}
