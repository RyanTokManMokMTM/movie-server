# FYP Project - Rebuilt with Go-ZERO

### TODO LIST


### Work in progress
| Service | Method |      Api       |        Desc         |
|:-------:|:------:|:--------------:|:-------------------:|
|  User   |  POST  | /users/signup  |     Create User     |
|  User   |  POST  | /users/signin  |        Login        |
|  User   |  GET   | /users/profile |  GET USER PROFILE   |
|  User   |  ———   |      ———       | Update User Profile |
|  User   |  ———   |      ———       |  Update User Token  |

| Service | Method |      Api       |                   Desc                   |
|:-------:|:------:|:--------------:|:----------------------------------------:|
|  POST   |  POST  |     /posts     |               Create Post                |
|  POST   |  GET   |   /posts/all   | Get All Post(10 posts by recent created) |
|  POST   | PATCH  |     /posts     |               UPDATE POST                |
|  POST   | DELETE |     /posts     |               DELETE POST                |
|  POST   |  GET   | /posts/:userID |             GET USER PROFILE             |
|  POST   |  GET   |     /posts     |           GET USER RECENT POST           |

| Service | Method |    Api     |                   Desc                   |
|:-------:|:------:|:----------:|:----------------------------------------:|
|  MOVIE  |  POST  |   /posts   |               Create Post                |
|  MOVIE  |  GET   | /posts/all | Get All Post(10 posts by recent created) |
|  MOVIE  |  ———   |    ———     |              ADD MOVIE INFO              |
|  MOVIE  |  ———   |    ———     |            UPDATE MOVIE INFO             |
|  MOVIE  |  ———   |    ———     |          GET upcoming etc MOVIE          |
|  MOVIE  |  ———   |    ———     |              Movie Trailer               |


| Service | Method |      Api       |                   Desc                   |
|:-------:|:------:|:--------------:|:----------------------------------------:|
|  LIST   |  POST  |     /lists     |               Create List                |
|  LIST   | PATCH  |     /lists     |               UPDATE POST                |
|  LIST   | DELETE |     /lists     |               DELETE POST                |
|  LIST   |  GET   | /lists/:userID |             GET USER PROFILE             |
|  LIST   |  GET   |     /lists     |           GET USER RECENT POST           |