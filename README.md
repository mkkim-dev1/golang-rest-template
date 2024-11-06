# Golang Rest 템플릿 
 Go 언어를 이용한 Rest 서비스

 위의 Golang 코드는 Gin 웹 프레임워크와 GORM ORM 라이브러리를 사용하여 RESTful API 서버를 구현한 예제입니다. 이 서버는 SQLite 데이터베이스와 연결되며, 책(Book) 정보를 저장하고 관리할 수 있는 CRUD(생성, 조회, 수정, 삭제) 기능을 제공합니다. 각 부분을 차례로 설명하겠습니다.

 ## 프로젝트 구조
 프로젝트는 아래와 같은 폴더 구조로 구성되어 있습니다.

```
 go-rest-template/
├── pkg/
│   ├── api/
│   │   ├── books.go                # Rest API 엔드포인트 정의
│   │   ├── message.go              # Websocket 엔드포인트 정의
│   │   └── server.go               # API 서버 설정
│   ├── config/
│   │   └── config.go               # 설정 관리 (환경 변수)
│   ├── db/
│   │   └── db.go             # 데이터베이스 초기화 및 연결 설정
│   ├── models/
│   │   └── book.go                 # Book 모델 정의
│   ├── repository/
│   │   ├── book_repository.go      # CRUD 기능 구현
│       └── book_repository_test.go # CRUD 기능 테스트 코드
├── main.go                         # 메인 함수: 설정 초기화, DB 연결, 서버 시작
└── .env                            # 환경 변수 파일 (포트 및 데이터베이스 경로)
```

## 각 파일 설명

### main.go: 메인 함수
메인 함수는 서버를 초기화하고 실행하는 역할을 합니다.
- config.InitConfig()로 설정을 초기화하여 .env 파일에서 환경 변수를 불러옵니다.
- db.InitDB()로 데이터베이스를 연결하고 테이블을 생성합니다.
- api.StartServer()로 서버를 시작합니다.

### pkg/config/config.go: 설정 파일
config 패키지는 .env 파일에서 설정 값을 불러와 전역으로 관리합니다. 포트(SERVER_PORT)와 데이터베이스 경로(DATABASE_PATH)를 설정합니다.

### pkg/db/database.go: 데이터베이스 초기화
데이터베이스 연결과 초기화를 담당하는 파일입니다. GORM을 이용하여 SQLite 데이터베이스를 연결하고, models.Book 모델의 테이블을 자동 생성합니다.

### pkg/models/book.go: 모델 정의
Book 모델은 gorm.Model을 포함하고 있으며, 제목(Title)과 저자(Author) 필드를 포함합니다. 각 필드는 데이터베이스 컬럼으로 매핑됩니다.

### pkg/repository/book.go: CRUD 기능 구현
Book 모델에 대해 데이터베이스와 상호작용하는 CRUD 기능을 정의합니다. 여기에는 CreateBook, GetBookByID, GetAllBooks, UpdateBook, DeleteBook 함수가 포함됩니다.

### pkg/api/server.go: 웹 서버 구성
서버를 설정하고 Gin 라우터를 이용해 WEB Server를 구성합니다.

### pkg/api/books.go: REST API 서비스 구현
Book 데이터에 대한 CRUD를 REST API로 제공하는 서비스를 구현한 코드 입니다.

### pkg/api/message.go: Chat 서비스 구현
Websocket을 이용한 Chat 서비스를 구현한 코드 입니다.

## 코드 요약

이 코드는 책 정보를 저장, 조회, 수정, 삭제하는 RESTful API 서버를 구현합니다. 주요 흐름은 다음과 같습니다:

	1.	config에서 설정 초기화.
	2.	db에서 데이터베이스 연결 및 테이블 마이그레이션.
	3.	repository에서 Book 모델의 CRUD 기능을 구현.
	4.	api에서 서버를 시작하고, 각 CRUD 엔드포인트를 정의.

테스트는 REST Client 파일을 사용하여 직접 API를 호출하고 응답을 확인하여 검증할 수 있습니다.

## 테스트 및 실행

이 프로젝트를 실행하기 위한 단계는 다음과 같습니다.

### DB 선택
지원되는 Database 연결은 sqlite 와 mysql 이 있습니다.
환경 변수에서 DB_KIND를 변경 하실 수 있습니다.

- SQL Lite: sqlite
- Mysql(MariaDB): mysql

### 프로젝트 클론 및 설정 파일 생성
1. 프로젝트를 클론하거나 생성한 후, 프로젝트 폴더로 이동합니다.
2. 프로젝트 루트에 .env 파일을 생성하고 서버 포트와 데이터베이스 경로를 설정합니다.

.env 예시
``` bash
DB_KIND=sqlite
SERVER_PORT=8080
DSN=test.db
```

### 의존성 설치
최초 프로젝트 실행 혹은 패키지가 추가되었다면 모듈을 초기화한 후 필요한 패키지를 설치합니다.
``` bash
go mod tidy
```

### 서버 실행
메인 함수를 포함한 ./main.go 파일을 실행하여 서버를 시작합니다.
``` bash
go run ./main.go
```
서버가 성공적으로 실행되면 터미널에 Listening and serving HTTP on :8080과 같은 메시지가 출력됩니다. (8080은 설정한 포트 번호)

### API 테스트
서버가 실행 중이면, 아래 방법으로 API를 테스트할 수 있습니다.

- VSCode REST Client 확장 프로그램을 이용해 /test/test.rest 파일을 열고 각 요청 위에 나타나는 Send Request 버튼을 클릭합니다. 각 요청에 대한 응답이 VSCode 내에 표시됩니다.
- curl 명령어를 사용해 터미널에서 테스트할 수도 있습니다. 예를 들어, 책을 생성하는 API를 호출하려면 아래와 같이 입력합니다.
``` bash
curl -X POST http://localhost:8080/books -H "Content-Type: application/json" -d '{"title": "Golang Basics", "author": "Author A"}'
```
- Postman 또는 Insomnia와 같은 HTTP 클라이언트 툴을 사용할 수도 있습니다.

### Unit 테스트 실행
pkg/repository/book_test.go 파일에서 작성한 테스트 코드를 실행하여 CRUD 함수가 올바르게 작동하는지 확인할 수 있습니다.
``` bash
go test ./pkg/repository
```
이 명령어는 repository 패키지에 있는 모든 테스트 파일을 실행하고, 결과를 터미널에 출력합니다.

### 서버 종료
서버를 종료하려면, 실행 중인 터미널에서 Ctrl + C를 눌러 서버를 중지합니다.

