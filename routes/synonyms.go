package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oezg/keymatch-api/models"
)

func getSynonyms(context *gin.Context) {
	synonyms, err := models.GetSynonyms()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
	}

	context.JSON(http.StatusOK, synonyms)
}

func createSynonym(context *gin.Context) {
	var synonym models.Synonym
	err := context.ShouldBindJSON(&synonym)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the request data"})
		return
	}

	err = synonym.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create synonym group. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Synonym group created.", "synonym": synonym})
}

func getSynonym(context *gin.Context) {
	synonym_id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse synonym id."})
		return
	}

	synonym, err := models.GetSynonym(synonym_id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Could not find synonym", "id": synonym_id})
		return
	}

	context.JSON(http.StatusOK, synonym)
}
