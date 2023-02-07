.PHONY : build build-darwin coverage

build: main.go
	GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o go-ap-incentive-linux

build-darwin: main.go
	GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o go-ap-incentive-darwin

run:
	go run *.go -config_name=config.local

mocks:
	mockgen -source=controllers/health/c_health.go -destination=mocks/health/health_controller_mock.go -package=mocks
	mockgen -source=services/health/s_health.go -destination=mocks/health/health_service_mock.go -package=mocks
	
	mockgen -source=controllers/setting/c_setting.go -destination=mocks/setting/setting_controller_mock.go -package=mocks
	mockgen -source=services/setting/s_setting.go -destination=mocks/setting/setting_service_mock.go -package=mocks
	
	mockgen -source=controllers/scheduler/c_scheduler.go -destination=mocks/scheduler/scheduler_controller_mock.go -package=mocks
	mockgen -source=services/scheduler/s_scheduler.go -destination=mocks/scheduler/scheduler_service_mock.go -package=mocks

	mockgen -source=client/redis/redis.go -destination=mocks/redis/redis.go -package=mocks
test:
	GOPRIVATE=github.com/luqmanulhakimdev go test ./... --count=1 -v -race -short

coverage:
	go test -race -coverprofile=cover.out ./...
	go tool cover -html cover.out

coverage-all:
	go test -v -coverpkg=./... -coverprofile=profile.cov ./...
	go tool cover -func profile.cov

generate-swagger:
	swag init -g main.go --output docs --parseDependency
