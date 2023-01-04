run:
	go run movieservice.go

api:
	goctl api go -api ./apis/movie-service.api --dir=./ --home=./tool

mock:
	mockgen -destination ./internal/dao/mock/store.go -package mockdb github.com/ryantokmanmokmtm/movie-server/internal/dao Store

swagger:
	swagger serve -F=swagger ./data/rest.swagger.json