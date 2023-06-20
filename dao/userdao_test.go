package dao

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试bookdao的方法")
	m.Run()
}

// func TestUser(t *testing.T) {
// 	// fmt.Println("开始测试userdao中的函数")
// 	t.Run("测试查询所有图书", testGetBooks)

// }

func TestBook(t *testing.T) {
	fmt.Println("测试")
	// t.Run("测试添加session", testGetOrder)
	// t.Run("测试更新图书", testGetBookByPrice)
}

// //测试页面
// func testGetPages(t *testing.T) {

// 	pages, _ := GetPageBooksByPrice("1", "12", "80")
// 	for _, book := range pages.Books {
// 		fmt.Printf("book: %v\n", book)
// 	}
// 	fmt.Printf("pages.PageNum: %v\n", pages.PageNum)
// 	fmt.Printf("pages.PageSize: %v\n", pages.PageSize)
// 	fmt.Printf("pages.TotalPageNum: %v\n", pages.TotalPageNum)
// 	fmt.Printf("pages.TotalRecord: %v\n", pages.TotalRecord)

// 	// fmt.Printf("books: %v\n", pages.Books)
// }

// //测试根据价格获取图书

// func testGetBookByPrice(t *testing.T) {
// 	pages, err := GetPageBooksByPrice("1", "12", "80")
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	for _, book := range pages.Books {
// 		fmt.Printf("book: %v\n", book)
// 	}

// }
// func testSession(t *testing.T) {
// 	sess := &model.Session{
// 		SessionID: "qwqeqweqqweewe",
// 		UserName:  "闫俊伟",
// 		UserID:    2,
// 	}
// 	AddSession(sess)
// }
// func testDeleteSession(t *testing.T) {
// 	sessionID := "qwqeqweqqweewe"
// 	DeleteSession(sessionID)
// }
// func testGetSession(t *testing.T) {
// 	sessionID := "5f8846c9-73bc-49a0-7bb8-9d4bf40578f5"
// 	session, _ := GetSessionByID(sessionID)
// 	fmt.Printf("session: %v\n", session)
// }

// //测试添加购物项
// func testAddCart(t *testing.T) {
// 	//设置图书
// 	book := &model.Book{
// 		ID:    12,
// 		Price: 33,
// 	}
// 	book2 := &model.Book{
// 		ID:    13,
// 		Price: 40,
// 	}
// 	//购物项
// 	cartItem := &model.CartItem{
// 		Book:   book,
// 		Count:  2,
// 		CartID: "666888",
// 	}
// 	cartItem2 := &model.CartItem{
// 		Book:   book2,
// 		Count:  4,
// 		CartID: "666888",
// 	}
// 	cartItems := []*model.CartItem{cartItem, cartItem2}
// 	//创建购物车
// 	cart := &model.Cart{
// 		CartID:    "666888",
// 		CartItems: cartItems,
// 		UserID:    9,
// 	}
// 	AddCart(cart)

// }

// func testGetCartItemByBookID(t *testing.T) {
// 	cartItem, _ := GetCartItemByBookID("12")
// 	fmt.Printf("cartItem: %v\n", cartItem)

// }

// func testGetCartItemByCartID(t *testing.T) {
// 	fmt.Println("1")
// 	cartItems, _ := GetCartItemsByCartID("666888")
// 	fmt.Println("2")
// 	for _, cartitem := range cartItems {
// 		fmt.Printf("cartitem: %v\n", cartitem)
// 	}
// }

// func testGetCartByUserID(t *testing.T) {
// 	cart, _ := GetCartByUserID(3)
// 	fmt.Printf("cart: %v\n", cart)

// }

// func testDeleteCart(t *testing.T) {
// 	cartID := "4ef1acec-f735-463c-458b-df41a40db269"
// 	DeleteCartByCartID(cartID)
// }

// //测试根据图书id删除购物项
// func testDeteteCartItemByBookID(t *testing.T) {
// 	cartItemID := "76"

// 	DeleteCartItemByID(cartItemID)
// }

//测试添加订单项
// func testAddOrderItem(t *testing.T) {

// 	//获取图书信息
// 	book, err := GetBookById("4")
// 	if err != nil {
// 		fmt.Printf("获取图书信息出错: %v\n", err)
// 	}

// 	orderItem := &model.OrderItem{
// 		Count:   3,
// 		Amount:  43.2,
// 		Titlt:   book.Title,
// 		Author:  book.Author,
// 		Price:   book.Price,
// 		ImgPath: book.ImgPath,
// 		OrderID: "dasdasdasdasda",
// 	}
// 	//要先创建order才可以创建orderitem
// 	order := &model.Order{
// 		OrderID:     "dasdasdasdasda",
// 		CreateTime:  time.Now(),
// 		TotalCount:  11,
// 		TotalAmount: 342.2,
// 		State:       1,
// 		UserID:      9,
// 	}
// 	AddOrder(order)
// 	AddOrderItem(orderItem)
// }

//测试获取订单
// func testGetOrder(t *testing.T) {

// 	userID := 8
// 	order, err := GetOrderByUserID(userID)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// 	fmt.Printf("order: %v\n", order)

// 	fmt.Println("****************************")
// 	orderitems, err2 := GetOrderItemByOrderID(order.OrderID)
// 	if err2 != nil {
// 		fmt.Printf("err2: %v\n", err2)

// 	}
// 	for _, orderitem := range orderitems {
// 		fmt.Printf("orderitem: %v\n", orderitem)
// 	}
// }
