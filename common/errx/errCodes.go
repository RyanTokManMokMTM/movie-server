package errx

//Prefix 3 number - which service
//suffix 3 number - which service error
//100xxx common server error
//jwt,db,server,param etc...
//110xxx user service error
//120xxx movie service error
const (
	SUCCESS                uint32 = 20
	SERVER_COMMON_ERROR    uint32 = 100001
	REQ_PARAM_ERROR        uint32 = 100002
	TOKEN_EXPIRED_ERROR    uint32 = 100003
	TOKEN_GENERATE_ERROR   uint32 = 100004
	TOKEN_INVALID_ERROR    uint32 = 100005
	DB_ERROR               uint32 = 100006
	DB_AFFECTED_ZERO_ERROR uint32 = 100007
)

//User Service
const (
	USER_NOT_EXIST            uint32 = 110001
	EMAIL_HAS_BEEN_REGISTERED uint32 = 110002
	USER_PASSWORD_INCORRECT   uint32 = 110003
)

//Movie Service 120xxx
const (
	MOVIE_NOT_EXIST     uint32 = 120001
	MOVIE_ALREADY_LIKED uint32 = 120002
)

//List Service 130xxx
const (
	LIST_NOT_EXIST uint32 = 130001
)

//Post Service 140xxx
const (
	POST_NOT_EXIST uint32 = 150001
)
