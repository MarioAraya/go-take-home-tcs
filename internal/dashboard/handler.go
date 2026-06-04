package dashboard

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DashboardHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no id"})
		return
	}

	user, err := FetchUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
		return
	}

	todos, err := FetchTodos(id)
	if err != nil {
		result := Aggregate(user, &TodosResponse{})
		warning := "Todos Unavailable"
		result.ErrorWarning = &warning
		c.JSON(http.StatusOK, result)
		return
	}

	result := Aggregate(user, todos)
	c.JSON(http.StatusOK, result)
}
