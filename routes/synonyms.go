package routes

import "github.com/gin-gonic/gin"

func GetSynonyms(context *gin.Context) {
	query := `
	SELECT 
		synonyms.domain
		synonyms.url
		synonyms.caption
		keymatches.keyword
	FROM synonyms
	JOIN keymatches ON synonyms.id = keymatches.synonym_id 
	`
	rows, err := 
}