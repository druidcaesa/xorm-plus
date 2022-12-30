package xormplus

type Page[T any] struct {
	CurrentPage int64
	PageSize    int64
	Total       int64
	Pages       int64
	Data        []T
}

func (page *Page[T]) SelectPage(wrapper QueryWrapper[T]) (e error) {
	e = nil
	clone := wrapper.DB.Clone()
	coun, err := clone.Count()
	if err != nil {
		return err
	}
	page.Total = coun
	if page.Total == 0 {
		page.Data = []T{}
		return
	}
	paginate, i := Paginate(page)
	err = wrapper.DB.Limit(i, paginate).Find(&page.Data)
	if err != nil {
		return err
	}
	return nil
}

func Paginate[T any](page *Page[T]) (int, int) {
	if page.CurrentPage <= 0 {
		page.CurrentPage = 0
	}
	switch {
	case page.PageSize > 100:
		page.PageSize = 100
	case page.PageSize <= 0:
		page.PageSize = 10
	}
	page.Pages = page.Total / page.PageSize
	if page.Total%page.PageSize != 0 {
		page.Pages++
	}
	p := page.CurrentPage
	if page.CurrentPage > page.Pages {
		p = page.Pages
	}
	size := page.PageSize
	offset := int((p - 1) * size)
	return offset, int(size)
}
