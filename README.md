# DDD Todo 애플리케이션

도메인 주도 설계(DDD)와 명령 쿼리 책임 분리(CQRS) 패턴을 구현한 Go 기반의 Todo 애플리케이션입니다.

## 아키텍처

이 프로젝트는 DDD와 CQRS 패턴을 적용한 클린 아키텍처 접근 방식을 따릅니다:

### DDD 계층
```
├── cmd                 # 애플리케이션 진입점
├── internal           
│   ├── domain         # 도메인 계층: 엔티티, 값 객체, 도메인 서비스
│   ├── application    # 애플리케이션 계층: 유스케이스, 커맨드, 쿼리
│   ├── infrastructure # 인프라 계층: 리포지토리, 외부 서비스
│   └── interfaces     # 인터페이스 계층: HTTP 핸들러, 미들웨어
└── pkg                # 공유 패키지
```

### CQRS 패턴
이 애플리케이션은 읽기와 쓰기 작업을 분리하는 CQRS를 구현합니다:
- **Commands**: 쓰기 작업 처리 (생성, 수정, 삭제)
- **Queries**: 읽기 작업 처리 (조회, 목록)

각 작업 유형은 자체 모델, 리포지토리 및 유스케이스를 가집니다.

## 사용 기술

- **언어**: Go 1.23.0
- **웹 프레임워크**: Fiber
- **데이터베이스**: PostgreSQL
- **ORM**: Ent
- **의존성 주입**: Wire
- **설정 관리**: Viper

## 사전 요구사항

- Go 1.23.0 이상
- Docker 및 Docker Compose
- Make (선택사항)

## 시작하기

1. 저장소 복제
```bash
git clone https://github.com/yourusername/ddd-todo-app.git
cd ddd-todo-app
```

2. PostgreSQL 데이터베이스 시작
```bash
docker-compose up -d
```

3. 의존성 설치
```bash
go mod download
```

4. Wire 의존성 주입 코드 생성
```bash
go generate ./...
```

5. 애플리케이션 실행
```bash
go run cmd/server/main.go
```

## 데이터베이스 설정

이 애플리케이션은 PostgreSQL을 데이터베이스로 사용합니다. Docker를 사용하여 시작할 수 있습니다:

```bash
# PostgreSQL 컨테이너 시작
docker-compose up -d

# PostgreSQL 컨테이너 중지
docker-compose down

# 로그 보기
docker-compose logs -f
```

## 환경 변수

루트 디렉토리에 `.env` 파일을 생성하세요:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=todos
DB_SSL_MODE=disable

SERVER_PORT=8080
```

## API 엔드포인트

### Todo API

#### Todo 생성
```http
POST /todos
Content-Type: application/json

{
    "title": "할 일 제목",
    "description": "할 일 설명"
}
```

응답:
```json
{
    "message": "Todo created successfully",
    "id": 1
}
```

#### ID로 Todo 조회
```http
GET /todos/:id
```

응답:
```json
{
    "id": 1,
    "title": "할 일 제목",
    "description": "할 일 설명",
    "completed": false
}
```

## 개발

### 프로젝트 구조
```
.
├── cmd/
│   └── server/          # 애플리케이션 진입점
├── config/              # 설정
├── internal/
│   ├── domain/
│   │   ├── entity/     # 도메인 엔티티
│   │   ├── service/    # 도메인 서비스
│   │   └── repository/ # 리포지토리 인터페이스
│   ├── application/
│   │   ├── command/    # 커맨드 유스케이스
│   │   └── query/      # 쿼리 유스케이스
│   ├── infrastructure/
│   │   └── persistence/# 리포지토리 구현
│   └── interfaces/
│       └── http/       # HTTP 핸들러
└── pkg/                # 공유 패키지
```

### 코드 생성
```bash
# Wire 의존성 주입 코드 생성
go generate ./...
```

## 테스트

```bash
# 모든 테스트 실행
go test ./...

# 커버리지와 함께 테스트 실행
go test -cover ./...
```

## 기여하기

1. 저장소를 포크합니다
2. 기능 브랜치를 생성합니다 (`git checkout -b feature/amazing-feature`)
3. 변경사항을 커밋합니다 (`git commit -m '새로운 기능 추가'`)
4. 브랜치에 푸시합니다 (`git push origin feature/amazing-feature`)
5. Pull Request를 생성합니다

## 라이선스

이 프로젝트는 MIT 라이선스를 따릅니다 - 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.
