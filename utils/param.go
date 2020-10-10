package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetOrderParam(c *gin.Context) []string {
	var orderBy []string
	orders := c.DefaultQuery("order", "course_id")
	for _, order := range strings.Split(orders, ",") {
		order = strings.TrimSpace(order)
		if strings.Contains(order, "-") {
			order = strings.Replace(order, "-", "", -1)
			orderBy = append(orderBy, fmt.Sprintf("%s DESC", order))
		} else {
			orderBy = append(orderBy, fmt.Sprintf("%s ASC", order))
		}
	}
	return orderBy
}

func GetPageParam(c *gin.Context) (int, int) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "25"))
	return page, limit
}

func GetIntParam(key string, c *gin.Context) uint {
	val := c.Param(key)
	fmt.Println("val from Param ", val)
	id, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println("err from strconv ", err)
		return 0
	}
	return uint(id)
}

func GetOrderParams(c *gin.Context) []string {
	var orderBy []string
	orders := c.DefaultQuery("order", "rating")
	for _, order := range strings.Split(orders, ",") {
		order = strings.TrimSpace(order)
		orderBy = append(orderBy, fmt.Sprintf("%s DESC", order))
	}
	return orderBy
}

func GetOrderParamsByDate(c *gin.Context) []string {
	var orderBy []string
	orders := c.DefaultQuery("order", "date")
	for _, order := range strings.Split(orders, ",") {
		order = strings.TrimSpace(order)
		orderBy = append(orderBy, fmt.Sprintf("%s DESC", order))
	}
	return orderBy
}

func GetOrderParamsByLowest(c *gin.Context) []string {
	var orderBy []string
	orders := c.DefaultQuery("order", "rating")
	for _, order := range strings.Split(orders, ",") {
		order = strings.TrimSpace(order)
		orderBy = append(orderBy, fmt.Sprintf("%s ASC", order))
	}
	return orderBy
}
