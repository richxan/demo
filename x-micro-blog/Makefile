SHELL := /bin/bash
.PHONY: all build clean run test proto docker-up docker-down docker-restart db-drop db-reset

# 默认目标
all: build

# 构建所有服务
build:
	@echo "Building all services..."
	go build -o bin/api-service cmd/api-service/main.go
	go build -o bin/user-service cmd/user-service/main.go
	go build -o bin/post-service cmd/post-service/main.go

# 清理构建产物
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f */*/tmp/

# 运行所有服务
run: build
	@echo "Running all services..."
	MICRO_REGISTRY=mdns MICRO_CONFIG_FILE=config/config.yml MICRO_ENV=development ./bin/api-service -c config/config.yml start &
	MICRO_REGISTRY=mdns MICRO_CONFIG_FILE=config/config.yml MICRO_ENV=development ./bin/user-service &
	MICRO_REGISTRY=mdns MICRO_CONFIG_FILE=config/config.yml MICRO_ENV=development ./bin/post-service

# 运行单个服务
run-api: build
	@echo "Running API service..."
	MYSQL_HOST=127.0.0.1 MYSQL_PORT=3314 MYSQL_USER=root MYSQL_PASSWORD=root123 MYSQL_DATABASE=x_micro_blog \
	MICRO_REGISTRY=mdns MICRO_CONFIG_FILE=config/config.yml MICRO_ENV=development ./bin/api-service -c config/config.yml start

run-user: build
	@echo "Running User service..."
	MICRO_REGISTRY=mdns MICRO_CONFIG_FILE=config/config.yml MICRO_ENV=development ./bin/user-service

run-post: build
	@echo "Running Post service..."
	MICRO_REGISTRY=mdns MICRO_CONFIG_FILE=config/config.yml MICRO_ENV=development ./bin/post-service

# 运行测试
test:
	@echo "Running tests..."
	go test -v ./...

# 生成 protobuf 代码
proto:
	@echo "Generating protobuf code..."
	protoc --proto_path=proto --go_out=. --micro_out=. proto/user/user.proto
	protoc --proto_path=proto --go_out=. --micro_out=. proto/post/post.proto

# Docker 相关命令
docker-up:
	@echo "Starting all containers..."
	docker-compose up -d

docker-down:
	@echo "Stopping all containers..."
	docker-compose down

docker-restart: docker-down docker-up

# 查看服务日志
logs:
	@echo "Showing service logs..."
	docker-compose logs -f

# 清理 Docker 资源
docker-clean: docker-down
	@echo "Cleaning Docker resources..."
	docker-compose down -v
	docker system prune -f

# 重新构建并启动服务
rebuild: clean build docker-restart

# 检查服务状态
status:
	@echo "Checking service status..."
	docker-compose ps

# 进入容器的便捷命令
mysql-cli:
	docker exec -it xblog-mysql mysql -uroot -proot123 x_micro_blog

redis-cli:
	docker exec -it xblog-redis redis-cli

# 初始化数据库
init-db:
	@echo "Initializing database..."
	chmod +x scripts/db_init.sh
	./scripts/db_init.sh

# 使用镜像加速拉取镜像
docker-pull:
	@echo "Pulling Docker images from mirror..."
	docker pull registry.cn-hangzhou.aliyuncs.com/library/mysql:8.0
	docker pull registry.cn-hangzhou.aliyuncs.com/library/redis:6.2
	docker pull registry.cn-hangzhou.aliyuncs.com/library/rabbitmq:3.9-management
	docker pull registry.cn-hangzhou.aliyuncs.com/library/prom/prometheus:latest
	docker pull registry.cn-hangzhou.aliyuncs.com/library/grafana/grafana:latest
	docker pull registry.cn-hangzhou.aliyuncs.com/library/jaegertracing/all-in-one:latest
	docker pull registry.cn-hangzhou.aliyuncs.com/library/oliver006/redis_exporter:latest
	docker pull registry.cn-hangzhou.aliyuncs.com/library/prom/node-exporter:latest
	docker pull registry.cn-hangzhou.aliyuncs.com/library/google/cadvisor:latest
	docker pull registry.cn-hangzhou.aliyuncs.com/library/prom/mysqld-exporter:latest

# 删除数据库
db-drop:
	@echo "Dropping database..."
	chmod +x scripts/db_drop.sh
	./scripts/db_drop.sh

# 重置数据库（删除并重新初始化）
db-reset: db-drop init-db

# 启动 Swagger UI
swagger-ui:
	@echo "Starting Swagger UI..."
	docker-compose up -d swagger-ui

# 停止 Swagger UI
swagger-ui-down:
	@echo "Stopping Swagger UI..."
	docker-compose stop swagger-ui

# 重启 Swagger UI
swagger-ui-restart: swagger-ui-down swagger-ui

# 帮助信息
help:
	@echo "Available commands:"
	@echo "  make build          - Build all services"
	@echo "  make clean          - Clean build artifacts"
	@echo "  make run           - Run all services"
	@echo "  make run-api       - Run API service only"
	@echo "  make run-user      - Run User service only"
	@echo "  make run-post      - Run Post service only"
	@echo "  make test           - Run tests"
	@echo "  make proto          - Generate protobuf code"
	@echo "  make docker-up      - Start all containers"
	@echo "  make docker-down    - Stop all containers"
	@echo "  make docker-restart - Restart all containers"
	@echo "  make logs           - View service logs"
	@echo "  make docker-clean   - Clean Docker resources"
	@echo "  make rebuild        - Rebuild and restart services"
	@echo "  make status         - Check service status"
	@echo "  make mysql-cli      - Connect to MySQL CLI"
	@echo "  make redis-cli      - Connect to Redis CLI"
	@echo "  make init-db        - Initialize database"
	@echo "  make db-drop        - Drop database and remove volume"
	@echo "  make db-reset       - Reset database (drop and reinitialize)"
	@echo "  make swagger-ui     - Start Swagger UI"
	@echo "  make swagger-ui-down - Stop Swagger UI"
	@echo "  make swagger-ui-restart - Restart Swagger UI"
	@echo "  make docker-pull    - Pull Docker images from mirror" 