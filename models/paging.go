package models

const PAGING_MODEL_NAME = "paging"

type PagingIncoming struct {
	Page         int `form:"page"`
	ItemsPerPage int `form:"itemsPerPage"`
}

type PagingOutGoing struct {
	Page         int
	ItemsPerPage int
	Total        int
}
