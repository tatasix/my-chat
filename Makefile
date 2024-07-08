# 定义变量
API_DIR = chat/service/chat/api
MODEL_DIR = chat/service/chat/model

# 定义目标
.PHONY: generate_api

# 定义生成API的规则
generate_api:
	cd $(API_DIR) && goctl api go -api chat.api -dir .

# 定义目标
.PHONY: generate_model

generate_model:
	cd $(MODEL_DIR) && goctl model mysql ddl -src user_portrait.sql -dir . -c

# 定义目标
.PHONY: build

# 定义生成API的规则
build:
	@git pull && docker-compose build chat && docker-compose stop chat && docker-compose up -d chat

# 定义目标
.PHONY: restart

restart:
	@docker-compose up -d && docker-compose restart chat && docker-compose restart server



# 定义目标
.PHONY: build_admin_web

build_admin_web:
	@cd WhiteMeow-admin && git pull && cd .. && docker-compose build web && docker-compose stop web && docker-compose up -d web



