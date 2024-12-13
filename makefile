POSTGRES_CONTAINER = todo_postgres
DEV_MIGRATIONS_DB = todo_db_migrations_dev
DB_MIGRATIONS_URL = postgres://postgres:postgres@localhost:5432/$(DEV_MIGRATIONS_DB)?search_path=public&sslmode=disable
DEV_DB = todo_db
DB_URL = postgres://postgres:postgres@localhost:5432/$(DEV_DB)?search_path=public&sslmode=disable
REVISIONS_SCHEMA = atlas_schema_revisions


.PHONY: new-migration
new-migration:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make new-migration name=<migration_name>"; \
		exit 1; \
	fi
	# 임시 DB 준비 (깨끗한 상태)
	docker exec -i $(POSTGRES_CONTAINER) dropdb -U postgres $(DEV_MIGRATIONS_DB) 2>/dev/null || true
	docker exec -i $(POSTGRES_CONTAINER) createdb -U postgres $(DEV_MIGRATIONS_DB)
	
	# 깨끗한 상태에서 새로운 마이그레이션 생성
	atlas migrate diff $(name) \
		--dir "file://migrations" \
		--to "ent://internal/infrastructure/persistence/ent/schema" \
		--dev-url "$(DB_MIGRATIONS_URL)"
	
	# 생성된 마이그레이션 파일 테스트를 위해 적용
	atlas migrate apply \
		--dir "file://migrations" \
		--url "$(DB_MIGRATIONS_URL)"
	
	# 임시 DB 정리
	docker exec -i $(POSTGRES_CONTAINER) dropdb -U postgres $(DEV_MIGRATIONS_DB)

.PHONY: migrate
migrate:
	atlas migrate apply \
		--dir "file://migrations" \
		--url "$(DB_URL)" \
		--revisions-schema "$(REVISIONS_SCHEMA)"

.PHONY: migrate-status
migrate-status:
	atlas migrate status \
		--dir "file://migrations" \
		--url "$(DB_URL)" \
		--revisions-schema "$(REVISIONS_SCHEMA)"