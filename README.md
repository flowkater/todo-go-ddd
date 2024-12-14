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
gh repo clone flowkater/todo-go-ddd
cd todo-go-ddd
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

#### EntGo (ORM)

EntGo는 Facebook에서 개발한 강력한 엔티티 프레임워크입니다. `go generate` 명령어를 실행하면 다음과 같은 작업이 수행됩니다:

1. **스키마 정의에서 코드 생성**: `internal/infrastructure/persistence/ent/schema` 디렉토리의 스키마 정의를 기반으로 엔티티 타입, CRUD 클라이언트, 쿼리 빌더를 생성합니다.
2. **타입 안전성**: 컴파일 타임에 타입 체크가 가능한 쿼리 API를 제공합니다.
3. **관계 처리**: 엔티티 간의 관계(1:1, 1:N, N:M)를 자동으로 처리합니다.
4. **마이그레이션**: Atlas와 통합되어 데이터베이스 마이그레이션을 자동으로 생성합니다.

#### Wire (의존성 주입)

Wire는 Google에서 개발한 컴파일 타임 의존성 주입 도구입니다:

1. **컴파일 타임 검증**: 런타임이 아닌 컴파일 타임에 의존성 그래프를 검증합니다.
2. **코드 생성**: `wire.go` 파일의 Provider 세트를 기반으로 의존성 주입 코드를 자동 생성합니다.
3. **명시적 의존성**: 모든 의존성이 코드에 명시적으로 선언되어 있어 디버깅이 용이합니다.
4. **테스트 용이성**: 모의 객체(mock)를 쉽게 주입할 수 있어 단위 테스트가 용이합니다.

```bash
# Wire 의존성 주입 코드 생성
go generate ./...
```

## 미들웨어 및 에러 처리

### 서버 미들웨어

서버는 다음과 같은 미들웨어를 사용합니다:

1. **로깅 미들웨어**: 
   - 요청/응답 로깅
   - 실행 시간 측정
   - 요청 ID 추적

2. **인증 미들웨어**:
   - JWT 토큰 검증
   - 사용자 인증 상태 확인

3. **에러 처리 미들웨어**:
   - 패닉 복구
   - 도메인 에러를 HTTP 응답으로 변환
   - 구조화된 에러 응답 제공

### 에러 처리

에러 처리는 다음 계층으로 구성됩니다:

1. **도메인 에러**:
   - 비즈니스 규칙 위반
   - 엔티티 유효성 검사 실패

2. **애플리케이션 에러**:
   - 유스케이스 실행 실패
   - 트랜잭션 실패

3. **인프라 에러**:
   - 데이터베이스 연결 실패
   - 외부 서비스 통신 실패

각 에러는 고유한 에러 코드와 메시지를 가지며, HTTP 응답 시 적절한 상태 코드로 매핑됩니다.

## 인터페이스 레이어

### DTO (Data Transfer Objects)

DTO는 클라이언트와 서버 간의 데이터 전송을 위한 객체입니다:

1. **유효성 검사**:
   - `pkg/validator` 패키지를 사용한 구조체 태그 기반 검증
   - 커스텀 유효성 검사 규칙 지원
   - 다국어 에러 메시지 지원

2. **요청 DTO**:
   - 클라이언트 입력 데이터 검증
   - 도메인 객체로의 변환

3. **응답 DTO**:
   - 도메인 객체를 클라이언트 응답 형식으로 변환
   - 민감한 정보 필터링
   - 일관된 응답 구조 제공

예시:
```go
type CreateTodoDTO struct {
    Title     string `json:"title" validate:"required,min=1,max=100"`
    Content   string `json:"content" validate:"required"`
    Priority  int    `json:"priority" validate:"min=1,max=5"`
}
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
