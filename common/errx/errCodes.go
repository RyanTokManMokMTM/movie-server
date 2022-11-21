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
	USER_NOT_EXIST                 uint32 = 110001
	EMAIL_HAS_BEEN_REGISTERED      uint32 = 110002
	USER_PASSWORD_INCORRECT        uint32 = 110003
	USER_UPLOAD_USER_AVATAR_FAILED uint32 = 110004
	USER_UPLOAD_USER_COVER_FAILED  uint32 = 110005
)

//Movie Service 120xxx
const (
	MOVIE_NOT_EXIST     uint32 = 120001
	MOVIE_ALREADY_LIKED uint32 = 120002
	MOVIE_NOT_YET_LIKED uint32 = 120003
)

//List Service 130xxx
const (
	LIST_NOT_EXIST             uint32 = 130001
	LIST_MOVIE_ALREADY_IN_LIST uint32 = 130002
	LIST_MOVIE_NOT_IN_LIST     uint32 = 130003
)

//Post Service 140xxx
const (
	POST_NOT_EXIST uint32 = 140001
)

//Comment Service 150xxx
const (
	POST_COMMENT_NOT_EXIST uint32 = 150001
)

//Friend Service 160xxx
const (
	FOLLOW_FRIEND_ERROR   uint32 = 160001
	UNFOLLOW_FRIEND_ERROR uint32 = 160002
)

//PostsLiked Service
const (
	NOT_LIKE_POST_YET uint32 = 1700001
)

//CommentLiked Service
const (
	NOT_LIKE_COMMENT_YET uint32 = 180001
)

//WebSocket error
const (
	WEBSOCKET_CONNECTION_ERROR    uint32 = 200001
	WEBSOCKET_READ_MESSAG_ERROR   uint32 = 200002
	WEBSOCKET_WRITE_MESSAGE_ERROR uint32 = 20003
)
