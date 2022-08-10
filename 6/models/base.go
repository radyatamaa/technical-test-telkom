package models

import (
	"math"
	"net/http"
)

type Count struct {
	Count int `json:"count"`
}
type ResponseCommand struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}
type ResponseCommandResult struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}
type MetaPagination struct {
	Page          int `json:"page"`
	Total         int `json:"total_pages"`
	TotalRecords  int `json:"total_records"`
	Prev          int `json:"prev"`
	Next          int `json:"next"`
	RecordPerPage int `json:"record_per_page"`
}


type Response struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status_desc"`
	Msg        string      `json:"message"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"errors"`
}

type Paginator struct {
	CurrentPage int `json:"current_page"`
	PerPage int `json:"limit_per_page"`
	PreviousPage int `json:"back_page"`
	NextPage int `json:"next_page"`
	TotalRecords int `json:"total_records"`
	TotalPages int `json:"total_pages"`
}

func (r *Response) MappingResponseSuccess(message string, data interface{}) {
	r.StatusCode = http.StatusOK
	r.Status = http.StatusText(r.StatusCode)
	r.Msg = message
	r.Data = data
	r.Errors = nil
}

func (r *Response) MappingResponseError(statusCode int,message string, error interface{}) {
	r.StatusCode = statusCode
	r.Status = http.StatusText(r.StatusCode)
	r.Msg = message
	r.Data = nil
	r.Errors = error
}
func (p Paginator)MappingPaginator(page,limit,totalAllRecords,countData int) Paginator {
	totalPage := int(math.Ceil(float64(totalAllRecords) / float64(limit)))
	prev := page
	next := page
	if page != 1 {
		prev = page - 1
	}

	if page != totalPage {
		next = page + 1
	}

	p = Paginator{
		CurrentPage:  page,
		PerPage:      countData,
		PreviousPage: prev,
		NextPage:     next,
		TotalRecords: totalAllRecords,
		TotalPages:   totalPage,
	}

	return p
}