package resp

// ------------------------------------------------------
// Created by fei wo at 2020/11/16
// ------------------------------------------------------
// Copyright©2020-2030
// ------------------------------------------------------
// blog: http://www.feiwo.xyz
// ------------------------------------------------------
// email: zhuyongluck@qq.com
// ------------------------------------------------------
//  分页对象
// ------------------------------------------------------

type Page struct {
	Index   int `json:"index"`
	Size    int `json:"size"`
	Total   int64 `json:"total"`
	Records interface{} `json:"records"`
}

func buildPage(index int, size int, total int64, data interface{}) *Page {
	return &Page{
		Index:   index,
		Size:    size,
		Total:   total,
		Records: data,
	}
}
