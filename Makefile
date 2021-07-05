APP_NAME = golang-fiber-base-project
APP_VERSION = 1.0
BUILD_DIR = ./build
DOCKER_LOCAL = localhost:5000
DOCKER_SERVER = 68.183.224.212:5000

config:
	cp config.example.yml config.yml

clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf *.out

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

swag:
	swag init

build: swag clean
	CGO_ENABLED=0  go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)

sonar:
	sonar-scanner.bat -Dproject.settings=./sonar-project.properties

docker.mariadb:
	docker run --rm -d \
		--name demo_mariadb \
		-e MYSQL_ROOT_PASSWORD=123456aA@ \
		-v data_mariadb:/var/lib/mysql \
		-p 3306:3306 \
		-d mariadb:10 \
		--character-set-server=utf8mb4 \
		--collation-server=utf8mb4_unicode_ci

docker.mysql:
	docker run --rm -d \
		--name demo_mysql \
		-e MYSQL_ROOT_PASSWORD=123456aA@ \
		-v data_mariadb:/var/lib/mysql \
		-p 3306:3306 \
		-d mysql:8 \
		--character-set-server=utf8mb4 \
		--collation-server=utf8mb4_unicode_ci

docker.sqlserver:
	docker run --rm -d \
		--name demo_sqlserver \
		-e ACCEPT_EULA=Y \
		-e SA_PASSWORD=123456aA@ \
		-p 1433:1433 \
		-d mcr.microsoft.com/mssql/server:2019-latest

docker.redis:
	docker run --rm -d \
	    --name demo_redis \
	    -p 6379:6379 \
	    -d redis:alpine

docker.sonar:
	docker run --rm -d \
    	--name demo_sonarqube \
    	-e SONAR_ES_BOOTSTRAP_CHECKS_DISABLE=true \
    	-p 9000:9000 \
    	-d sonarqube:8-community

docker.build:
	docker build -t $(DOCKER_LOCAL)/$(APP_NAME):$(APP_VERSION) .

docker.push:
	docker push $(DOCKER_LOCAL)/$(APP_NAME):$(APP_VERSION)

docker.run:
	docker run --name $(APP_NAME) -d -p 8000:8000 $(DOCKER_LOCAL)/$(APP_NAME):$(APP_VERSION) -e SERVER_PORT=8000

docker.deploy: docker.build docker.run

docker.local: docker.build docker.push

docker.server:
	docker build -t $(DOCKER_LOCAL)/$(APP_NAME):$(APP_VERSION) .
	docker push $(DOCKER_LOCAL)/$(APP_NAME):$(APP_VERSION)