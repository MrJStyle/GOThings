package models

type Video struct {
	Id          int    `uri:"id"`
	Title       string `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
}
