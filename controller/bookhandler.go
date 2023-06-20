package controller

import (
	"bookstore/dao"
	"bookstore/model"
	"html/template"
	"net/http"
	"strconv"
)

//indexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	pageNum := r.FormValue("pageNo")
	if pageNum == "" {
		pageNum = "1"
	}
	//调用获取bookdao中带分页的函数
	pages, _ := dao.GetPageBooks(pageNum)
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, pages)
}

// func GetBooks(w http.ResponseWriter, r *http.Request) {
// 	//调用获取bookdao中获取所有图书的函数
// 	books, _ := dao.GetBooks()
// 	//解析模板文件
// 	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))

// 	//执行
// 	t.Execute(w, books)
// }

// //添加图书
// func AddBook(w http.ResponseWriter, r *http.Request) {
// 	//获取图书信息
// 	title := r.PostFormValue("title")
// 	author := r.PostFormValue("author")
// 	price := r.PostFormValue("price")
// 	sales := r.PostFormValue("sales")
// 	stock := r.PostFormValue("stock")
// 	//将价格销量库存转化类型
// 	fprice, _ := strconv.ParseFloat(price, 64)
// 	fsales, _ := strconv.ParseInt(sales, 10, 0)
// 	fstock, _ := strconv.ParseInt(stock, 10, 0)
// 	book := &model.Book{
// 		Title:   title,
// 		Author:  author,
// 		Price:   fprice,
// 		Sales:   int(fsales),
// 		Stock:   int(fstock),
// 		ImgPath: "static/img/default1.jpg",
// 	}
// 	//调用bookdao中的添加图书方法AddBook
// 	dao.AddBook(book)
// 	//提交之后调用getbook函数，再次查询数据库
// 	GetBooks(w, r)
// }

//删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// 获取id
	bookId := r.FormValue("bookId")
	//调用删除方法
	dao.DeleteBook(bookId)
	//删除之后返回查询页面
	GetPageBooks(w, r)
}

//去更新或者添加图书的页面
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	//获取id
	bookId := r.FormValue("bookId")
	//调用根据id查询图书的函数
	book, _ := dao.GetBookById(bookId)
	if book.ID > 0 {
		//更新图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, book)
	} else {
		//没有这个书，就添加图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, "")
	}

}

// //UpdateBook更新图书
// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	//获取bookid
// 	bookId := r.PostFormValue("bookId")
// 	//获取书的信息
// 	title := r.PostFormValue("title")
// 	author := r.PostFormValue("author")
// 	price := r.PostFormValue("price")
// 	sales := r.PostFormValue("sales")
// 	stock := r.PostFormValue("stock")
// 	//将价格销量库存转化类型
// 	fprice, _ := strconv.ParseFloat(price, 64)
// 	fsales, _ := strconv.ParseInt(sales, 10, 0)
// 	fstock, _ := strconv.ParseInt(stock, 10, 0)
// 	iid, _ := strconv.ParseInt(bookId, 10, 0)
// 	id := int(iid)
// 	book := &model.Book{
// 		ID:      id,
// 		Title:   title,
// 		Author:  author,
// 		Price:   fprice,
// 		Sales:   int(fsales),
// 		Stock:   int(fstock),
// 		ImgPath: "static/img/default1.jpg",
// 	}
// 	//调用更新id的方法
// 	dao.UpdateBook(book)

// 	//更新完之后需要再次查询所有图书信息
// 	GetBooks(w, r)
// }

//更新或者添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	//获取bookid
	bookId := r.PostFormValue("bookId")
	//获取书的信息
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	//将价格销量库存转化类型
	fprice, _ := strconv.ParseFloat(price, 64)
	fsales, _ := strconv.ParseInt(sales, 10, 0)
	fstock, _ := strconv.ParseInt(stock, 10, 0)
	iid, _ := strconv.ParseInt(bookId, 10, 0)
	id := int(iid)
	book := &model.Book{
		ID:      id,
		Title:   title,
		Author:  author,
		Price:   fprice,
		Sales:   int(fsales),
		Stock:   int(fstock),
		ImgPath: "static/img/default1.jpg",
	}

	if book.ID > 0 {
		//更新
		//调用更新id的方法
		dao.UpdateBook(book)
	} else {
		//调用bookdao中的添加图书方法AddBook
		dao.AddBook(book)
	}

	//更新完之后需要再次查询所有图书信息
	GetPageBooks(w, r)
}

//获取带分页的图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//首先要获取页码
	pageNum := r.FormValue("pageNo")
	if pageNum == "" {
		pageNum = "1"
	}
	//调用获取bookdao中带分页的函数
	pages, _ := dao.GetPageBooks(pageNum)
	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))

	//执行
	t.Execute(w, pages)
}

//根据价格删选图书
func GetPageBooksByPrice(w http.ResponseWriter, r *http.Request) {
	//获取获取传入的min和max
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")

	//首先要获取页码
	pageNum := r.FormValue("pageNo")
	if pageNum == "" {
		pageNum = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		page, _ = dao.GetPageBooks(pageNum)

	} else {
		//调用获取bookdao中带分页的函数
		page, _ = dao.GetPageBooksByPrice(pageNum, minPrice, maxPrice)
		//传入价格范围
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice

	}
	//调用IsLogin方法
	flag, session := dao.IsLogin(r)
	if flag {
		page.IsLogin = true
		page.Username = session.UserName
	}

	//解析模板文件
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, page)
}
