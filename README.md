# 🌐 go-rest-api <img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=Go&logoColor=white"/> <img src="https://github.com/user-attachments/assets/255241de-50ff-46c3-962d-dfd96db004d4" height="20"> <img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" height="30"> <img src="https://private-user-images.githubusercontent.com/53367916/255967639-d92caabf-98e0-473e-bfbf-ab554ba435e5.png?jwt=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3MzQzMDg2MTIsIm5iZiI6MTczNDMwODMxMiwicGF0aCI6Ii81MzM2NzkxNi8yNTU5Njc2MzktZDkyY2FhYmYtOThlMC00NzNlLWJmYmYtYWI1NTRiYTQzNWU1LnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNDEyMTYlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjQxMjE2VDAwMTgzMlomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPTZkZjc2YTMzZjcwNDAxODliZTEwZTRmOWRhOTg3N2MwMjIxYTU2ZjIzOTY0MjYwMzIzNmE4M2ZlYjA3MTYwYmUmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.vk97Iu-QEARclZ9v69KbRChFYIBz1wG9MuXjaqVYT44" height="30">
## 👉 Go 언어 기반의 웹 프레임워크들을 활용하여 구현한 REST API (fiber, gin, mux 등)

### ✅ 프로젝트 소개
다양한 Go 웹프레임워크를 사용하여 REST API를 설계하는 방식을 알아보기 위한 용도로 작성한 코드입니다.
학생 관리용으로 REST API가 필요한 상황을 가정하여 만들었습니다.
최근 재밌게 읽은 책 중 하나인 ‘수학귀신’ 속 등장인물을 예시로 사용하여 학생 데이터를 구성했습니다.

### REST와 REST API
REST(Representational State Transfer)는 인터넷에서 데이터를 주고받는 방식입니다.  
이를 사용하면 자원을 이름으로 구분한 후, 해당 자원의 상태를 주고받을 수 있습니다.

'REST API'는 REST 원리를 따르는 API라고 볼 수 있습니다.  
POST(생성), GET(읽기), PUT(업데이트), DELETE(삭제)와 같은 HTTP 동사를 사용하여 클라이언트와 서버 간의 통신을 관리합니다.

어떤 가게에서 물건을 주문한다고 가정을 해보겠습니다.  
실제 작업과 API 내용이 아래처럼 1:1 대응한다고 이해하면 좋습니다.
* POST : 새로 물건을 주문하는 것
* GET : 가게에 있는 물건을 보는 것
* PUT : 주문한 물건을 다른 걸로 바꾸는 것
* DELETE : 주문을 취소하는 것

이 가게의 사장님(서버)은 HTTP라는 공통 언어를 사용하여 소통합니다.
그리고 물건(데이터)을 URL이라는 주소로 정리해 두고, 손님이 요청을 보내면 그에 맞는 데이터를 보내주거나 작업을 처리합니다.
물건을 보낼 때는 주로 JSON이라는 정리된 상자를 사용하고, "요청이 성공했는지, 실패했는지"를 알려주는 상태 코드도 같이 보내줍니다.

결론적으로 REST는 서버(가게 사장님)와 클라이언트(손님)가 약속된 규칙으로 대화하는 방식이라고 할 수 있습니다!

### 코드 동작 전에 해야 하는 작업
1. 'rest-go-fiber' 폴더가 위치에서 'go build' 명령어를 입력합니다.
2. 'rest-go-gin' 폴더가 위치에서 'go build' 명령어를 입력합니다.
3. 'rest-go-mux' 폴더가 위치에서 'go build' 명령어를 입력합니다.

### fiber API 코드 동작 방법 (Linux 기준 작성)
1. 'rest-go-fiber' 폴더 위치에서 './rest-go-fiber' 명령어를 입력하여 fiber API를 활성화시킵니다.
2. 웹 브라우저 또는 Postman에서 'http://
localhost:3000/students'가 성공적으로 잘 동작하는지 확인합니다.

### gin API 코드 동작 방법 (Linux 기준 작성)
1. 'rest-go-gin' 폴더 위치에서 './rest-go-gin' 명령어를 입력하여 gin API를 활성화시킵니다.
2. 웹 브라우저 또는 Postman에서 'http://
localhost:3000/students'가 성공적으로 잘 동작하는지 확인합니다.

### mux API 코드 동작 방법 (Linux 기준 작성)
1. 'rest-go-mux' 폴더 위치에서 './rest-go-mux' 명령어를 입력하여 mux API를 활성화시킵니다.
2. 웹 브라우저 또는 Postman에서 'http://
localhost:3000/students'가 성공적으로 잘 동작하는지 확인합니다.

### 코드 실행 결과
![image](https://github.com/user-attachments/assets/9b15795c-0ad7-4a0e-aeb2-17b5d126c6a5)


