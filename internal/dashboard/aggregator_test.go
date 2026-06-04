package dashboard

import "testing"

// table-driven test
func TestAggregate(t *testing.T) {
	tests := []struct {
		name     string
		user     *User
		todos    *TodosResponse
		expected DashboardResponse
	}{
		{
			name:     "30 must be Rookie",
			user:     &User{ID: 1, FirstName: "Mario", LastName: "Romero", Age: 30},
			todos:    &TodosResponse{},
			expected: DashboardResponse{Status: "Rookie", FullName: "Mario Romero"},
		},
		{
			name:     "55 must be Veteran",
			user:     &User{ID: 2, FirstName: "Carlos", LastName: "Santana", Age: 55},
			todos:    &TodosResponse{},
			expected: DashboardResponse{Status: "Veteran", FullName: "Carlos Santana"},
		},
		{
			name:     "Fullname concatenates OK",
			user:     &User{ID: 3, FirstName: "John", LastName: "Carmack", Age: 30},
			todos:    &TodosResponse{},
			expected: DashboardResponse{Status: "Rookie", FullName: "John Carmack"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Aggregate(tt.user, tt.todos)
			if result.Status != tt.expected.Status {
				t.Errorf("status: got %s, want %s", result.Status, tt.expected.Status)
			}
			if result.FullName != tt.expected.FullName {
				t.Errorf("full_name: got %s, want %s", result.FullName, tt.expected.FullName)
			}
		})
	}
}
