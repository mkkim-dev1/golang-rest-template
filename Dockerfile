# 1. Build Stage
FROM golang:1.20 AS builder

# 2. 작업 디렉토리 생성 및 설정
WORKDIR /app

# 3. Go 모듈 파일 복사
COPY go.mod go.sum ./

# 4. 모듈 다운로드
RUN go mod download

# 5. 소스 코드 복사
COPY . .

# 6. 애플리케이션 빌드
RUN CGO_ENABLED=1 GOOS=linux go build -o app -a -ldflags '-linkmode external -extldflags "-static"' .

# 프로덕션용
# FROM scratch

# 개발용 
FROM alpine:latest

# 필요한 패키지 설치
RUN apt-get update && apt-get install -y \
    sqlite3 libsqlite3-0 && \
    rm -rf /var/lib/apt/lists/*

# 8. 빌드된 애플리케이션 복사
COPY --from=builder /app/app .

# 9. 데이터베이스 파일 복사 (optional)
COPY --from=builder /app/test.db .

# 10. 환경 변수 설정 (optional)
ENV SERVER_PORT=8080
ENV DATABASE_PATH=test.db
ENV GIN_MODE=release

ENTRYPOINT "./app"

