package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"strconv"
)

//获取数据库中所有的图书
func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books "
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//添加到books中
		books = append(books, book)

	}
	return books, nil
}

//添加图书
func AddBook(book *model.Book) error {
	//sql语句
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err := utils.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

//删除图书
func DeleteBook(bookID string) error {
	sqlStr := "delete from books where id = ?"
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}
	return nil

}

//getbook by id

func GetBookById(bookId string) (*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"

	row := utils.Db.QueryRow(sqlStr, bookId)
	book := &model.Book{}
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	return book, nil
}

// 更细图书 根据id

func UpdateBook(b *model.Book) error {
	//写sql语句
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		return err
	}
	return nil
}

// 分页
func GetPageBooks(pageNum string) (*model.Page, error) {

	//获取数据库中所有图书的数目
	sqlStr := "select count(*) from books"
	//总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int64 = 4
	//获取总页数
	var totalPageNum int64
	if totalRecord%pageSize == 0 {
		totalPageNum = totalRecord / pageSize
	} else {
		totalPageNum = totalRecord/pageSize + 1
	}
	//获取当前页中的图书
	sqlStr_1 := "select * from books limit ?,?"
	page, _ := strconv.ParseInt(pageNum, 10, 64)
	rows, err := utils.Db.Query(sqlStr_1, (page-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	//创建page
	pages := &model.Page{
		Books:        books,
		PageNum:      page,
		PageSize:     pageSize,
		TotalPageNum: totalPageNum,
		TotalRecord:  totalRecord,
	}
	return pages, nil
}

// 根据价格筛选图书
func GetPageBooksByPrice(pageNum, minPrice, maxPrice string) (*model.Page, error) {
	//获取数据库中所有图书的数目
	sqlStr := "select count(*) from books where price between ? and ?"
	//总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int64 = 4
	//获取总页数
	var totalPageNum int64
	if totalRecord%pageSize == 0 {
		totalPageNum = totalRecord / pageSize
	} else {
		totalPageNum = totalRecord/pageSize + 1
	}
	//获取当前页中的图书
	sqlStr_1 := "select * from books where price between ? and ? limit ?,?"
	page, _ := strconv.ParseInt(pageNum, 10, 64)
	rows, err := utils.Db.Query(sqlStr_1, minPrice, maxPrice, (page-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		books = append(books, book)
	}
	//创建page
	pages := &model.Page{
		Books:        books,
		PageNum:      page,
		PageSize:     pageSize,
		TotalPageNum: totalPageNum,
		TotalRecord:  totalRecord,
	}
	return pages, nil
}
