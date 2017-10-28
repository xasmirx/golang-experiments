package main

import (
	"fmt"
	"math"
)

type pagination struct {
	total_records 	int
	total_pages 	int
	per_page        int
	current_page 	int
	next_page		int
	previous_page 	int
	last_page 		int
	first_page 		int
}

func (p *pagination) setTotalPages() {
	c := float64(p.total_records) / float64(p.per_page)
	p.total_pages = int(math.Floor(c + 0.5))
}

func (p *pagination) setCurrentPage(current_page int) {
	p.current_page = current_page
} 

func (p *pagination) setNextPage() {
	if p.current_page + 1 <= p.total_pages {
		p.next_page++
	}
}

func (p *pagination) setPreviousPage() {
	if p.current_page - 1 > 0 {
		p.previous_page--
	} else {
		p.previous_page = 1
	}
}

func (p *pagination) setLastPage() {
	p.last_page = p.total_pages
}

func (p *pagination) init() {
	p.setTotalPages()
	p.setCurrentPage(1)
	p.setNextPage()
	p.setPreviousPage()
	p.setLastPage()
}

func main() {
	pager := pagination{total_records: 55, per_page: 15}
	pager.init()

	fmt.Printf("Total pages: %v Current page: %v Next page: %v Previous page: %v", pager.total_pages, pager.current_page, pager.next_page, pager.previous_page)
}