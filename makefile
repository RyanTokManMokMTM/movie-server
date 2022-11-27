run:
	go run movieservice.go

api:
	goctl api go -api ./apis/movie-service.api --dir=./ --home=./tool