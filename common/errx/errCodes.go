package errx

//Prefix 3 number - which service
//suffix 3 number - which service error
//100xxx common server error
//jwt,db,server,param etc...
//110xxx user service error
//120xxx movie service error
type InternalCode uint32

const (
	SUCCESS                InternalCode = 20
	SERVER_COMMON_ERROR    InternalCode = 100001
	REQ_PARAM_ERROR        InternalCode = 100002
	TOKEN_EXPIRED_ERROR    InternalCode = 100003
	TOKEN_GENERATE_ERROR   InternalCode = 100004
	TOKEN_INVALID_ERROR    InternalCode = 100005
	DB_ERROR               InternalCode = 100006
	DB_AFFECTED_ZERO_ERROR InternalCode = 100007
)

//User Service
const (
	USER_NOT_EXIST                 InternalCode = 110001
	EMAIL_HAS_BEEN_REGISTERED      InternalCode = 110002
	USER_PASSWORD_INCORRECT        InternalCode = 110003
	USER_UPLOAD_USER_AVATAR_FAILED InternalCode = 110004
	USER_UPLOAD_USER_COVER_FAILED  InternalCode = 110005
)

//Movie Service 120xxx
const (
	MOVIE_NOT_EXIST     InternalCode = 120001
	MOVIE_ALREADY_LIKED InternalCode = 120002
	MOVIE_NOT_YET_LIKED InternalCode = 120003
)

//List Service 130xxx
const (
	LIST_NOT_EXIST             InternalCode = 130001
	LIST_MOVIE_ALREADY_IN_LIST InternalCode = 130002
	LIST_MOVIE_NOT_IN_LIST     InternalCode = 130003
)

//Post Service 140xxx
const (
	POST_NOT_EXIST InternalCode = 140001
)

//Comment Service 150xxx
const (
	POST_COMMENT_NOT_EXIST InternalCode = 150001
)

//Friend Service 160xxx
const (
	FOLLOW_FRIEND_ERROR   InternalCode = 160001
	UNFOLLOW_FRIEND_ERROR InternalCode = 160002
)

//PostsLiked Service
const (
	NOT_LIKE_POST_YET InternalCode = 1700001
)

//CommentLiked Service
const (
	NOT_LIKE_COMMENT_YET InternalCode = 180001
)

//WebSocket error
const (
	WEBSOCKET_CONNECTION_ERROR    InternalCode = 200001
	WEBSOCKET_READ_MESSAG_ERROR   InternalCode = 200002
	WEBSOCKET_WRITE_MESSAGE_ERROR InternalCode = 20003
)
