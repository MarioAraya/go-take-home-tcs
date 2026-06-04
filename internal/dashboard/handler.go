package dashboard

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

func DashboardHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no id"})
		return
	}

	var user *User
	var todos *TodosResponse
	var userErr, todosErr error

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		user, userErr = FetchUser(id)
	}()

	go func() {
		defer wg.Done()
		todos, todosErr = FetchTodos(id)
	}()

	wg.Wait()

	if userErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user fetch failed"})
		return
	}

	if todosErr != nil {
		result := Aggregate(user, &TodosResponse{})
		warning := "Todos Unavailable"
		result.ErrorWarning = &warning
		c.JSON(http.StatusOK, result)
		return
	}

	result := Aggregate(user, todos)
	c.JSON(http.StatusOK, result)
}
