package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/synonyms", getSynonyms)    // Get all synonyms
	server.POST("/synonyms", createSynonym) // Create a new synonym group in a domain
	server.GET("/synonyms/:id", getSynonym) // Get a synonym
	// server.PUT("/synonyms/:id", updateSynonym)              // Update a synonym group's url and caption
	// server.DELETE("/synonyms/:id", deleteSynonym)           // Delete a synonym group from database
	// server.POST("/synonyms/:id/keymatch", addKeymatch)      // Add a keymatch to the synonym
	// server.DELETE("/synonyms/:id/keymatch", removeKeymatch) // Remove the keymatch from the synonym
}
