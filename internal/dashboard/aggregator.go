package dashboard

import "fmt"

func Aggregate(user *User, todos *TodosResponse) DashboardResponse {
	result := DashboardResponse{
		ID:       user.ID,
		FullName: fmt.Sprintf("%s %s", user.FirstName, user.LastName),
	}

	if user.Age > 50 {
		result.Status = "Veteran"
	} else {
		result.Status = "Rookie"
	}

	for _, todo := range todos.Todos {
		if !todo.Completed {
			result.PendingTaskCount++
			if result.NextUrgentTask == nil {
				t := todo.Title
				result.NextUrgentTask = &t
			}
		}
	}

	return result
}
