package token

import (
	"net/http"

	"github.com/getupandgo/snooper-wooper/validators"
	"github.com/gin-gonic/gin"
)

func ParseText(c *gin.Context) {
	jsonBody := validators.Token{}

	if err := c.ShouldBindJSON(&jsonBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//todo: log stuff
		return
	}
	// todo: we must find a nice way to convert from validator to model here
	//saved, err := db.SaveToken(jsonBody)
	//if saved.Count == newToken.Count {
	//	w.WriteHeader(http.StatusCreated)
	//} else {
	//	// todo: change to OK when we return body
	//	w.WriteHeader(http.StatusNoContent)
	//}
	////todo: return the created/updated entity in body
	//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func GetTopTokens(c *gin.Context) {
	//	// todo:
	//	// it's a common scenario to support pagination with limit & offset
	//	// also would be great to have total count in the response body
	//	// fixme:
	//	// 2. what if limit is missing?
	params := validators.GetTopTokens{}

	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//todo: log stuff
		return
	}

	//	tokens, err := dao.GetTokens(limitNum); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	} else {
	//		// todo:
	//		// would be great if we had serialize function(or method)
	//		// in order to hide the exact implementation
	//		// so it will be possible to change serialization strategies
	//		c.JSON(http.StatusBadRequest, gin.H{"tokens": tokens})
	//	}
}
