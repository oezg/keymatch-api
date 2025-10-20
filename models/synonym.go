package models

type Synonym struct {
	Id      int64
	Domain  string `binding:"required"`
	Url     string `binding:"required"`
	Caption string `binding:"required"`
}
