package utils

import (
	"go-template/helpers/str"
	"math/rand"
)

const (
	letterBytes  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	defaultLimit = 10
	maxLimit     = 50
	defaultSort  = "asc"
)

var (
	sortWhitelist = []string{"asc", "desc"}
)

func GenerateProcessID() string {
	b := make([]byte, 20)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func SetPaginationParameter(page, limit int, orderBy, sort string, orderByWhiteLists, orderByStringWhiteLists []string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}

	orderBy = checkWhiteList(orderBy, orderByWhiteLists)
	if str.Contains(orderByStringWhiteLists, orderBy) {
		orderBy = `LOWER(` + orderBy + `)`
	}

	if !str.Contains(sortWhitelist, sort) {
		sort = defaultSort
	}
	offset := (page - 1) * limit

	return offset, limit, page, orderBy, sort
}

func checkWhiteList(orderBy string, whiteLists []string) string {
	for _, whiteList := range whiteLists {
		if orderBy == whiteList {
			return orderBy
		}
	}

	return "def.updated_at"
}
