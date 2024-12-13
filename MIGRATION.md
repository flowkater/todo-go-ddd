# 데이터베이스 마이그레이션 가이드

이 프로젝트는 데이터베이스 마이그레이션을 위해 Atlas를 사용합니다. 아래는 데이터베이스 마이그레이션을 관리하기 위한 일반적인 명령어와 절차입니다.

## 새 마이그레이션 생성하기

새로운 마이그레이션 파일을 생성하려면:

```bash
atlas migrate diff 마이그레이션_이름 \
  --dir "file://migrations" \
  --to "ent://internal/infrastructure/persistence/ent/schema" \
  --dev-url "postgres://postgres:postgres@localhost:5432/todo_db?sslmode=disable"
```

이 명령어는 `migrations` 디렉토리에 새로운 마이그레이션 파일을 생성합니다.

## 마이그레이션 적용하기 (Up)

모든 대기 중인 마이그레이션을 적용하려면:

```bash
atlas migrate apply \
  --dir "file://migrations" \
  --url "postgres://postgres:postgres@localhost:5432/todo_db?sslmode=disable"
```

특정 수의 마이그레이션만 적용하려면:

```bash
atlas migrate apply --num 1 \
  --dir "file://migrations" \
  --url "postgres://postgres:postgres@localhost:5432/todo_db?sslmode=disable"
```

## 마이그레이션 롤백하기 (Down)

마지막 마이그레이션을 롤백하려면:

```bash
atlas migrate down \
  --dir "file://migrations" \
  --url "postgres://postgres:postgres@localhost:5432/todo_db?sslmode=disable"
```

특정 수의 마이그레이션을 롤백하려면:

```bash
atlas migrate down --num 2 \
  --dir "file://migrations" \
  --url "postgres://postgres:postgres@localhost:5432/todo_db?sslmode=disable"
```

## 마이그레이션 상태 확인

마이그레이션 상태를 확인하려면:

```bash
atlas migrate status \
  --dir "file://migrations" \
  --url "postgres://postgres:postgres@localhost:5432/todo_db?sslmode=disable"
```

## 마이그레이션 초기화

모든 마이그레이션을 초기화하고 새로 시작해야 하는 경우:

1. `migrations` 디렉토리의 모든 마이그레이션 파일 삭제:
```bash
rm -rf migrations/*
```

2. 데이터베이스를 삭제하고 재생성하세요 (PostgreSQL):
```bash
psql -U postgres -c "DROP DATABASE todo_db;"
psql -U postgres -c "CREATE DATABASE todo_db;"
```

3. 새로운 초기 마이그레이션 생성:
```bash
atlas migrate diff init \
  --dir "file://migrations" \
  --to "ent://internal/infrastructure/persistence/ent/schema" \
  --dev-url "postgres://postgres:postgres@localhost:5432/todo_db?sslmode=disable"
```

## 모범 사례

1. 프로덕션에 적용하기 전에 항상 마이그레이션 파일을 검토하세요
2. 마이그레이션은 작고 구체적인 변경사항에 집중하세요
3. 개발 환경에서 먼저 마이그레이션을 테스트하세요
4. 프로덕션에서 마이그레이션을 적용하기 전에 데이터베이스를 백업하세요
5. 마이그레이션 파일의 이름은 변경 내용을 설명하는 의미 있는 이름을 사용하세요

## 일반적인 문제 해결

마이그레이션 상태가 "dirty" 상태가 된 경우:
1. 마이그레이션 상태를 확인하세요
2. 데이터베이스의 오류를 수정하세요
3. 필요한 경우 `atlas migrate fix` 명령어를 사용하세요

프로덕션 환경에서 마이그레이션을 실행하기 전에 항상 데이터베이스를 백업하는 것을 잊지 마세요!
