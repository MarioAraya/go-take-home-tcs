# ID Watchdog — Dashboard Aggregator API

BFF service that aggregates user profile + todo data from dummyjson.com

## Run

git clone https://github.com/MarioAraya/go-take-home-tcs
cd go-take-home-tcs
go mod tidy
go run ./cmd/api

Visit: http://localhost:8080

## Example

curl http://localhost:8080/dashboard/1

Response:
{
  "id": 1,
  "full_name": "Emily Johnson",
  "status": "Rookie",
  "pending_task_count": 5,
  "next_urgent_task": "Do something important",
  "error_warning": null
}

## Run tests

go test ./...
