package model

type Page struct {
	Books        []*Book //当前页的图书切片
	PageNum      int64   //当前页
	PageSize     int64   //每页的图书数量
	TotalPageNum int64   //图书
	TotalRecord  int64   //总记录数
	MinPrice     string
	MaxPrice     string
	IsLogin      bool
	Username     string
}

//如果是首页就不显示上一页之类的方法
func (page *Page) IsHasPrev() bool {
	return page.PageNum > 1

}

// 判断是否有下一页
func (page *Page) IsHasNext() bool {
	return page.PageNum < page.TotalPageNum
}

//获取上一页获取下一页
func (page *Page) GetPrevPageNo() int64 {
	if page.PageNum == 1 {
		return 1
	} else {
		return page.PageNum - 1

	}
}

//获取下一页
func (page *Page) GetNextPageNo() int64 {
	if page.IsHasNext() {
		return page.PageNum + 1
	} else {
		return page.TotalPageNum

	}
}
