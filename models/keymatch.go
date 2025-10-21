package models

type Keymatch struct {
	Id        int64
	Keyword   string `binding:"required"`
	Domain    string `binding:"required"`
	SynonymId int64  `binding:"required"`
}
