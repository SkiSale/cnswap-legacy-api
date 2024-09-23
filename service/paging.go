package service

import (
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type pagingArgs struct {
	PerPage     int
	CurrentPage int
	Offset      int
}

func parsePagingParameters(c *gin.Context) (*pagingArgs, error) {
	perPage, err := strconv.Atoi(c.DefaultQuery("perPage", "25"))
	if err != nil {
		return nil, err
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return nil, err
	}

	if page < 0 {
		return nil, errors.New("page out of range, must be greater than 1")
	}

	return &pagingArgs{
		PerPage:     perPage,
		CurrentPage: page,
		Offset:      perPage * (page - 1),
	}, nil
}

func pagingHeaders(c *gin.Context, total int, perPage int) {
	if perPage < 1 {
		return
	}
	tp := math.Ceil(float64(total) / float64(perPage))

	c.Header("X-Total", fmt.Sprintf("%d", total))
	c.Header("X-Total-Pages", fmt.Sprintf("%d", int(tp)))
}
