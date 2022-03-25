Design
======

## Ping

Ping indicates if the server is working.

| Name | Path    | Method | Content-Type       | Description                         |
|------|---------|--------|--------------------|-------------------------------------|
| Ping | `/ping` | `GET`  | `application/json` | Get Pong with status code 200 if OK |

## Authentication

- Base path: `/auth`
- Content-Type: `application/json`

| Name     | Resource       | Response | Code | Path        | Method | Description             |     
|----------|----------------|----------|------|-------------|--------|-------------------------|
| Register | `UserRegister` | `User`   | 200  | `/register` | `POST` | Register a new user     |
| Login    | `UserLogin`    | `Token`  | 200  | `/login`    | `POST` | Log in and create token |
| Refresh  | `TokenRefresh` | `Token`  | 200  | `/refresh`  | `GET`  | Refresh existing token  |

### How to register

You can register to the API

```shell
curl --location --request POST 'http://localhost:8080/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
        "lastName": "Fanny",
        "firstName": "Armand",
        "email": "fannyarmand2@gmail.com",
        "password": "secretpassword"
    
}'
```

**Sample:**

```json
{
  "id": "80a08d36-cfea-4898-aee3-6902fa562f1d",
  "createdAt": "2021-11-19T15:59:58.407451298+01:00",
  "updatedAt": "2021-11-19T15:59:58.407451298+01:00",
  "deletedAt": null,
  "lastName": "Baba",
  "firstName": "Ali",
  "email": "ali@gmail.com",
  "dateOfBirth": "",
  "apprenticeAt": "",
  "profilPic": "",
  "privateMail": "",
  "instagram": "",
  "facebook": "",
  "github": "",
  "linkedin": "",
  "mbti": "",
  "isAdmin": false,
  "promoId": "80a08d36-cfea-4898-aee3-6902fa562f0a",
  "bdaPosts": null,
  "posts": null
}
```

In this example, localhost:8080 is the address of your API.

### How to login

You can login :

```shell
curl --location --request POST 'http://localhost:8080/auth/login' \
--header 'Content-Type: application/json' \
--data-raw '{
        "email": "ali@gmail.com",
        "password": "alibabaalibaba"
    
}'
```

**Sample:**

```json
{
  "code": 200,
  "expire": "2021-11-19T17:06:58+01:00",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzczMzgwMTgsImZpcnN0bmFtZSI6IkFsaSIsImlkIjoiYWxpQGdtYWlsLmNvbSIsImxhc3RuYW1lIjoiQmFiYSIsIm9yaWdfaWF0IjoxNjM3MzM0NDE4fQ.YUicgImgZI1fUK6XRh6DlD3k8H3XDk6opNSTM63kfw8"
}
```

### How to refresh a token

You can renew an expired token:

```shell
curl --location --request GET 'http://localhost:8080/auth/refresh' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mzc5MjU2NzMsImZpcnN0bmFtZSI6IkFsaSIsImlkIjoiYWxpQGdtYWlsLmNvbSIsImxhc3RuYW1lIjoiQmFiYSIsIm9yaWdfaWF0IjoxNjM3OTIyMDczfQ.GKpWcKhI_DbVOlByFfxwyR99il7XgCHJzNmEfp-l7zw'
```

### How to refresh a token using axios

A token have to be renew if we have a forbidden (403) on a request because the token is expiring after one hour. But it
can be done easily with an interceptor axios. You can use and adapt the following article for example:
[Using Axios interceptors for refreshing your API token](https://thedutchlab.com/blog/using-axios-interceptors-for-refreshing-your-api-token)

### How to use API authenticated endpoint?

(add in header token, add in cookie token, add in query token and put some curl for example).

## Rest Api

- Base path: `/api/rest/v1`
- Content-Type: `application/json`
- Authentication: `true`
- Rights: `anyone`

| Name                   | Resource   | Response               | Code | Path                                | Method   | Description                                            |     
|------------------------|------------|------------------------|------|-------------------------------------|----------|--------------------------------------------------------|
| Get Current User       | `User`     | `User`                 | 200  | `/me`                               | `GET`    | Get the current user                                   |
| Update User password   | `User`     | `<empty>`              | 204  | `/me/password`                      | `PATCH`  | Update password of current user                        |
| List Posts             | `Post`     | `Collection<Post>`     | 200  | `/topics/:id/posts`                 | `GET`    | Retrieve a collection of post                          |
| Get Post               | `Post`     | `Post`                 | 200  | `/topics/:id/posts/:postId`         | `GET`    | Get a specific post                                    |
| Create Post            | `Post`     | `Post`                 | 200  | `/topics/:id/posts`                 | `POST`   | Create a new post                                      |
| Update Post            | `Post`     | `Post`                 | 200  | `/topics/:id/posts/:postId`         | `PATCH`  | Update a post                                          |
| Delete Post            | `Post`     | `<empty>`              | 204  | `/topics/:id/posts/:postId`         | `DELETE` | Delete a post                                          |
| List Post Likes        | `Like`     | `LikeCollection`       | 200  | `/posts/:id/likes`                  | `GET`    | Retrieve a collection of likes with a count and a bool |
| Create Post Like       | `Like`     | `LikePostResponse`     | 200  | `/posts/:id/likes`                  | `POST`   | Create a new like                                      |
| Delete Post Like       | `Like`     | `<empty>`              | 204  | `/posts/:id/likes/:likeId`          | `DELETE` | Delete a like                                          |
| List Users             | `User`     | `Collection<User>`     | 200  | `/users`                            | `GET`    | Retrieve a collection of user                          |
| Get User               | `User`     | `User`                 | 200  | `/users/:id`                        | `GET`    | Get a specific user                                    |
| Create User            | `User`     | `User`                 | 200  | `/users`                            | `POST`   | Create a new user                                      |
| Update User            | `User`     | `User`                 | 200  | `/users/:id`                        | `PATCH`  | Update a user                                          |
| Delete User            | `User`     | `<empty>`              | 204  | `/users/:id`                        | `DELETE` | Delete a user                                          |
| List  BdaPosts         | `BdaPost`  | `Collection<BdaPost>`  | 200  | `/bdaposts`                         | `GET`    | Retrieve a collection of bda post                      |
| Get BdaPost            | `BdaPost`  | `BdaPost`              | 200  | `/bdaposts/:id`                     | `GET`    | Get a specific bda post                                |
| Create  BdaPost        | `BdaPost`  | `BdaPost`              | 200  | `/bdaposts`                         | `POST`   | Create a new bda post                                  |
| Update  BdaPost        | `BdaPost`  | `BdaPost`              | 200  | `/bdaposts/:id`                     | `PATCH`  | Update a bda post                                      |
| Delete  BdaPost        | `BdaPost`  | `<empty>`              | 204  | `/bdaposts/:id`                     | `DELETE` | Delete a bda post                                      |
| List BdaPost Likes     | `Like`     | `LikeCollection`       | 200  | `/bdaposts/:id/likes`               | `GET`    | Retrieve a collection of likes with a count and a bool |
| Create BdaPost Like    | `Like`     | `LikeBdaPostResponse`  | 200  | `/bdaposts/:id/likes`               | `POST`   | Create a new like                                      |
| Delete BdaPost Like    | `Like`     | `<empty>`              | 204  | `/bdaposts/:id/likes/:likeId`       | `DELETE` | Delete a like                                          |   
| Create BdaPost Comment | `Comment`  | `Comment`              | 200  | `/bdaposts/:id/comments`            | `POST`   | Create a new comment                                   |
| Update BdaPost Comment | `Comment`  | `Comment`              | 200  | `/bdaposts/:id/comments/:commentId` | `PATCH`  | Update a comment                                       |
| Delete BdaPost Comment | `Comment`  | `<empty>`              | 204  | `/bdaposts/:id/comments/:commentId` | `DELETE` | Delete a comment                                       |
| List BdaPost Comments  | `Comment`  | `Collection<Comment>`  | 200  | `/bdaposts/:id/comments`            | `GET`    | Retrieve a collection of comment                       |
| Get BdaPost Comment    | `Comment`  | `Comment`              | 200  | `/bdaposts/:id/comments/:commentId` | `GET`    | Retrieve a specific comment                            |
| List Comment Likes     | `Like`     | `LikeCollection`       | 200  | `/comments/:id/likes`               | `GET`    | Retrieve a collection of likes with a count and a bool |
| Create Comment Like    | `Like`     | `LikeCommentResponse`  | 200  | `/comments/:id/likes`               | `POST`   | Create a new like                                      |
| Delete Comment Like    | `Like`     | `<empty>`              | 204  | `/comments/:id/likes/:likeId`       | `DELETE` | Delete a like                                          |
| List Promos            | `Promo`    | `Collection<Promo>`    | 200  | `/promos`                           | `GET`    | Retrieve a collection of promo                         |
| Create Promo           | `Promo`    | `Promo`                | 200  | `/promos`                           | `POST`   | Create a new promo                                     |
| Update Promo           | `Promo`    | `Promo`                | 200  | `/promos/:id`                       | `PATCH`  | Update a promo                                         |
| Delete Promo           | `Promo`    | `<empty>`              | 204  | `/promos/:id`                       | `DELETE` | Delete a promo                                         |
| Get Users Promo        | `Promo`    | `Users`                | 204  | `/promos/:id/users`                 | `GET`    | Get users of a promo                                   |
| Create Category        | `Category` | `Category`             | 200  | `/categories`                       | `POST`   | Create a category                                      |
| List Categories        | `Category` | `Collection<Category>` | 200  | `/categories`                       | `GET`    | List all categories                                    |
| Get Category           | `Category` | `Category `            | 200  | `/categories/:id`                   | `GET`    | Get a specific category                                |
| Update Category        | `Category` | `Category `            | 200  | `/categories/:id`                   | `PATCH`  | Update a category                                      |
| Delete Category        | `Category` | `<empty>`              | 204  | `/categories/:id`                   | `DELETE` | Delete a category                                      |
| Create Topic           | `Topic`    | `Topic`                | 200  | `/categories/:id/topics`            | `POST`   | Create a topic                                         |
| List Category Topics   | `Topic`    | `Collection<Topic>`    | 200  | `/categories/:id/topics`            | `GET`    | Get all the topics of a category                       |
| List Topics            | `Topic`    | `Collection<Topic>`    | 200  | `/topics`                           | `GET`    | Get all the topics                                     |
| Get Topic              | `Topic`    | `Topic`                | 200  | `/topics/:id`                       | `GET`    | Get a specific topic                                   |
| Update Topic           | `Topic`    | `Topic`                | 200  | `/topics/:id`                       | `PATCH`  | Update a topic                                         |
| Delete Topic           | `Topic`    | `<empty>`              | 204  | `/topics/:id`                       | `DELETE` | Delete a topic                                         |

### Resource

All resources will be represented with the following fields:

| Key         | Type     | Description                         |
|-------------|----------|-------------------------------------|
| `createdAt` | `string` | Date of creation in RFC 3339 format |
| `updatedAt` | `string` | Date of updation in RFC 3339 format |
| `deletedAt` | `string` | Date of deletion in RFC 3339 format |

### Collection

A collection represent a list of resources. Any resources will be represented in a common way.

For example:

```json
[
  {
    "id": "80a08d36-cfea-4898-aee3-6902fa562f1d",
    "createdAt": "2021-10-13T10:52:11.50932133+02:00",
    "updatedAt": "2021-10-13T10:52:11.50932133+02:00",
    "content": "lorem ipsum sit dolor set amet..."
  },
  {
    "id": "80a08d36-cfea-4898-aee3-6902fa562f2c",
    "createdAt": "2021-10-13T12:52:11.50932133+02:00",
    "updatedAt": "2021-10-13T12:52:11.50932133+02:00",
    "content": "foo bar..."
  }
]
```

### Errors

When a request can not be fulfilled an error will be returned with a status code >= 400 and the
following `application/json` content:

| Key       | Type     | Description   |
|-----------|----------|---------------|
| `message` | `string` | Error message |

The following errors are supported:

- `400`: The request is not valid
- `404`: The resource is not found
- `409`: The resource is in conflict (e.g. already exist)
- `500`: An internal error happened

## Resources

### Ping

A Ping represents a dummy response .

| Key       | Type     | Description   |
|-----------|----------|---------------|
| `message` | `string` | Dummy message |

**Sample:**

```json
{
  "message": "pong"
}
```

### Post

A Post represents a message .

| Key         | Type     | Creatable | Mutable | Required | Validation                | Description                             |
|-------------|----------|-----------|---------|----------|---------------------------|-----------------------------------------|
| `id`        | `string` | no        | no      | no       | no                        | Unique identifier for a `Post` resource |
| `title`     | `string` | yes       | yes     | yes      | `required,min=4,max=1024` | Title of a `Post` resource              |
| `content`   | `string` | yes       | yes     | yes      | `required,min=4,max=1024` | Content of a `Post` resource            |
| `userId`    | `string` | no        | no      | no       | no                        | User id of a `Post` resource            |
| `createdAt` | `string` | no        | no      | no       | no                        | Date of creation in RFC 3339 format     |
| `updatedAt` | `string` | no        | no      | no       | no                        | Date of updation in RFC 3339 format     |
| `deletedAt` | `string` | no        | no      | no       | no                        | Date of deletion in RFC 3339 format     |

**Sample:**

```json
{
  "id": "80a08d36-cfea-4898-aee3-6902fa562f1d",
  "createdAt": "2021-11-05T16:54:49.182599198+01:00",
  "updatedAt": "2021-11-05T16:54:49.182599198+01:00",
  "deletedAt": null,
  "title": " Titre",
  "content": "lorem ipsum sit dolor set amet...",
  "userId": "80a08d36-cfea-4898-aee3-6902fa562f8d"
}
```

### BdaPost

A Bda Post represents a message.

| Key         | Type     | Creatable | Mutable | Required | Validation                | Description                                |
|-------------|----------|-----------|---------|----------|---------------------------|--------------------------------------------|
| `id`        | `string` | no        | no      | no       | no                        | Unique identifier for a `BdaPost` resource |
| `title`     | `string` | yes       | no      | yes      | `required,min=4,max=1024` | Title of a `BdaPost` resource              |
| `content`   | `string` | yes       | no      | yes      | `required,min=4,max=1024` | Content of a `BdaPost` resource            |
| `userId`    | `string` | no        | no      | no       | no                        | User id of a `BdaPost` resource            |
| `createdAt` | `string` | no        | no      | no       | no                        | Date of creation in RFC 3339 format        |
| `updatedAt` | `string` | no        | no      | no       | no                        | Date of updation in RFC 3339 format        |
| `deletedAt` | `string` | no        | no      | no       | no                        | Date of deletion in RFC 3339 format        |

**Sample:**

```json
{
  "id": "80a08d36-cfea-4898-aee3-6902fa562f1d",
  "createdAt": "2021-11-05T17:04:13.475674216+01:00",
  "updatedAt": "2021-11-05T17:04:13.475674216+01:00",
  "deletedAt": null,
  "title": " Titre",
  "content": "lorem ipsum sit dolor set amet...",
  "userId": "80a08d36-cfea-4898-aee3-6902fa562f9k"
}
```

### User

A User represents informations about a user.

| Key            | Type                  | Creatable | Mutable | Required | Validation               | Description                             |                           
|----------------|-----------------------|-----------|---------|----------|--------------------------|-----------------------------------------|
| `id`           | `string`              | no        | no      | no       | no                       | Unique identifier for a `User` resource |
| `lastName`     | `string`              | yes       | no      | yes      | `required,min=2,max=20 ` | Last name of a `User` resource          |
| `firstName`    | `string`              | yes       | no      | yes      | `required,min=2,max=20`  | First name of a `User` resource         |
| `email`        | `string`              | yes       | no      | yes      | `required,email`         | Email of a `User` resource              |
| `Password`     | `string`              | no        | yes     | yes      | no                       | Hashed password of a `User`resource     |
| `dateOfBirth`  | `string`              | yes       | no      | yes      | no                       | Date of birth of a `User` resource      |
| `apprenticeAt` | `string`              | yes       | yes     | no       | no                       | Enterprise of a `User` resource         |
| `profilPic`    | `string`              | yes       | yes     | no       | no                       | Profil pic of a `User` resource         | 
| `privateMail`  | `string`              | yes       | yes     | no       | no                       | Private email of a `User` resource      |                         
| `instagram`    | `string`              | yes       | yes     | no       | no                       | Instagram Page of a `User` resource     | 
| `facebook`     | `string`              | yes       | yes     | no       | no                       | Facebook Page of a `User` resource      | 
| `github`       | `string`              | yes       | yes     | no       | no                       | Github Page of a `User` resource        |    
| `linkedin`     | `string`              | yes       | yes     | no       | no                       | Linkedin Page of a `User` resource      |
| `mbti`         | `string`              | yes       | no      | no       | no                       | Profil mbti of a `User` resource        |
| `isAdmin`      | `bool`                | no        | no      | no       | no                       | Profil admin of a `User` resource       | 
| `promoId`      | `string`              | yes       | no      | no       | no                       | Promo id of a `User` resource           |                                  
| `bdaPosts`     | `Collection<BdaPost>` | no        | no      | no       | no                       | Bda Posts of a `User` resource          |              
| `posts`        | `Collection<Post>`    | no        | no      | no       | no                       | Posts of a `User` resource              |        
| `createdAt`    | `string`              | no        | no      | no       | no                       | Date of creation in RFC 3339 format     |
| `updatedAt`    | `string`              | no        | no      | no       | no                       | Date of updation in RFC 3339 format     | 
| `deletedAt`    | `string`              | no        | no      | no       | no                       | Date of deletion in RFC 3339 format     |

**Sample:**

```json
{
  "id": "80a08d36-cfea-4898-aee3-6902fa562f1d",
  "createdAt": "2021-11-05T16:16:26.259246323+01:00",
  "updatedAt": "2021-11-05T16:16:26.259246323+01:00",
  "deletedAt": null,
  "lastName": "Lovelace",
  "firstName": "Ada",
  "email": "lovelace@gmail.com",
  "password": "$2a$10$cO3VM1aifnDImyhXqs1/xu9Oz1/NTjufIwyVivo2uuFHC2iI2DUCy",
  "dateOfBirth": "01/01/2020",
  "apprenticeAt": "Ada Tech School",
  "profilPic": "https://www.seekpng.com/png/detail/506-5061704_cool-profile-avatar-picture-cool-picture-for-profile.png",
  "privateMail": "ada@google.com",
  "instagram": "https://www.instagram.com/adatechschool/",
  "facebook": "https://www.facebook.com/AdaTechSchool",
  "github": "https://github.com/ada-social-network/",
  "linkedin": "https://www.linkedin.com/",
  "mbti": "INFP",
  "isAdmin": true,
  "promoId": "80a08d36-cfea-4898-aee3-6902fa562f0e",
  "bdaPost": null
}
```

### Promo

A promo represents informations about a promo.

| Key           | Type               | Creatable | Mutable | Required | Validation | Description                              |                           
|---------------|--------------------|-----------|---------|----------|------------|------------------------------------------|
| `id`          | `string`           | no        | no      | no       | no         | Unique identifier for a `Promo` resource |
| `promoName`   | `string`           | yes       | no      | no       | no         | Promo name of a `Promo` resource         |
| `dateOfStart` | `string`           | yes       | no      | no       | no         | Date of start of a `Promo` resource      |
| `dateOfEnd`   | `string`           | yes       | no      | no       | no         | Date of end of a `Promo` resource        |
| `biography`   | `string`           | yes       | no      | no       | no         | Biography of a `Promo` resource          |
| `createdAt`   | `string`           | no        | no      | no       | no         | Date of creation in RFC 3339 format      |
| `updatedAt`   | `string`           | no        | no      | no       | no         | Date of updation in RFC 3339 format      |
| `deletedAt`   | `string`           | no        | no      | no       | no         | Date of deletion in RFC 3339 format      |
| `users`       | `Collection<User>` | no        | no      | no       | no         | Multiple `User` of a `Promo`             |            |                                          |

**Sample:**

```json
{
  "id": "80a08d36-cfea-4898-aee3-6902fa562f1d",
  "createdAt": "2021-11-19T15:15:16.218234962+01:00",
  "updatedAt": "2021-11-19T15:15:16.218234962+01:00",
  "deletedAt": null,
  "promoName": "Béatrice Worsley",
  "dateOfStart": "05/10/2020",
  "dateOfEnd": "30/06/2021",
  "biography": "La seconde promo qui a vu le jour à l'école Ada Tech School",
  "users": "['80a08d36-cfea-4898-aee3-6902fa562f1d','c20ccc44-7ac6-11ec-90d6-0242ac120003']"
}
```

### Category

A Category represents informations about a category.

| Key         | Type                | Creatable | Mutable | Required | Validation | Description                                 |
|-------------|---------------------|-----------|---------|----------|------------|---------------------------------------------|
| `id`        | `string`            | no        | no      | no       | no         | Unique identifier for a `Category` resource |
| `name`      | `string`            | yes       | no      | yes      | no         | Name of a `Category` resource               |
| `topics`    | `Collection<Topic>` | no        | no      | no       | no         | Multiple `Topic` of a `Category`            |
| `createdAt` | `string`            | no        | no      | no       | no         | Date of creation in RFC 3339 format         |
| `updatedAt` | `string`            | no        | no      | no       | no         | Date of updation in RFC 3339 format         |
| `deletedAt` | `string`            | no        | no      | no       | no         | Date of deletion in RFC 3339 format         |

**Sample:**

```json
{
  "id": "7907465b-7507-4fa4-a649-b9d90a17bb58",
  "createdAt": "2022-01-26T22:00:54.12918234+01:00",
  "updatedAt": "2022-01-26T22:00:54.12918234+01:00",
  "deletedAt": null,
  "name": "first category",
  "topics": null
}
```

### Topic

A Topic represents informations about a topic.

| Key          | Type               | Creatable | Mutable | Required | Validation                | Description                              |
|--------------|--------------------|-----------|---------|----------|---------------------------|------------------------------------------|
| `id`         | `string`           | no        | no      | no       | no                        | Unique identifier for a `Topic` resource |
| `name`       | `string`           | yes       | no      | yes      | no                        | Name of a `Topic` resource               |
| `content`    | `string`           | yes       | no      | yes      | `required,min=4,max=1024` | Content of a `Topic` resource            |
| `userId`     | `string`           | no        | no      | no       | no                        | User id of a `Topic` resource            |
| `categoryId` | `string`           | no        | no      | yes      | no                        | Category id of a `Topic` resource        |
| `posts`      | `Collection<Post>` | no        | no      | no       | no                        | Multiple `Post` of a `Topic`             |
| `createdAt`  | `string`           | no        | no      | no       | no                        | Date of creation in RFC 3339 format      |
| `updatedAt`  | `string`           | no        | no      | no       | no                        | Date of updation in RFC 3339 format      |
| `deletedAt`  | `string`           | no        | no      | no       | no                        | Date of deletion in RFC 3339 format      |

**Sample:**

```json
{
  "id": "91b61685-d73a-468d-a98e-5984218bec87",
  "createdAt": "2022-01-28T11:33:37.280422692+01:00",
  "updatedAt": "2022-01-28T11:33:37.280422692+01:00",
  "deletedAt": null,
  "name": "ceci est un topic",
  "content": "lorem ipsum",
  "userId": "412c0459-9dad-4720-9438-70db28e32ae3",
  "categoryId": "7907465b-7507-4fa4-a649-b9d90a17bb58",
  "posts": null
}
```

### Comment

A Comment represents a comment under a Post.

| Key         | Type     | Creatable | Mutable | Required | Validation                | Description                                |
|-------------|----------|-----------|---------|----------|---------------------------|--------------------------------------------|
| `id`        | `string` | no        | no      | no       | no                        | Unique identifier for a `Comment` resource |
| `content`   | `string` | yes       | yes     | yes      | `required,min=4,max=1024` | Content of a `Comment` resource            |
| `userId`    | `string` | no        | no      | no       | no                        | User id of a `Comment` resource            |
| `bdapostId` | `string` | no        | no      | yes      | no                        | Bdapost id of a `Comment` resource         |
| `createdAt` | `string` | no        | no      | no       | no                        | Date of creation in RFC 3339 format        |
| `updatedAt` | `string` | no        | no      | no       | no                        | Date of updation in RFC 3339 format        |
| `deletedAt` | `string` | no        | no      | no       | no                        | Date of deletion in RFC 3339 format        |

**Sample:**

```json
{
  "id": "80a08d36-cfea-4898-aee3-6902fa562f1d",
  "createdAt": "2021-11-05T16:54:49.182599198+01:00",
  "updatedAt": "2021-11-05T16:54:49.182599198+01:00",
  "deletedAt": null,
  "content": "lorem ipsum sit dolor set amet...",
  "userId": "80a08d36-cfea-4898-aee3-6902fa562f8e",
  "bdapostId": "80a08d36-cfea-4898-aee3-6902fa562f9c"
}
```

### Like

A like represents a like on a resource(can be Post, BdaPost, Comment)

| Key         | Type     | Creatable | Mutable | Required | Validation                | Description                             |
|-------------|----------|-----------|---------|----------|---------------------------|-----------------------------------------|
| `id`        | `string` | no        | no      | no       | no                        | Unique identifier for a `Like` resource |
| `userId`    | `string` | no        | no      | no       | no                        | User id of a `Like` resource            |
| `bdapostId` | `string` | no        | no      | no       | no                        | Bdapost id of a `Like` resource         |
| `postId`    | `string` | no        | no      | no       | no                        | Post id of a `Like` resource            |
| `commentId` | `string` | no        | no      | no       | no                        | Comment id of a `Like` resource         |
| `createdAt` | `string` | no        | no      | no       | no                        | Date of creation in RFC 3339 format     |
| `updatedAt` | `string` | no        | no      | no       | no                        | Date of updation in RFC 3339 format     |
| `deletedAt` | `string` | no        | no      | no       | no                        | Date of deletion in RFC 3339 format     |

**Sample:**

```json
{
  "id": "05ad6bdf-da72-42fa-867d-427d9d10a0d6",
  "createdAt": "2022-01-14T18:16:59.469363507+01:00",
  "updatedAt": "2022-01-14T18:16:59.469363507+01:00",
  "deletedAt": null,
  "userId": "622977e4-0097-44ef-9089-29debe93058a",
  "bdapostId": "5ad258c5-4db6-4cc2-8798-f731196f32de",
  "postId": "00000000-0000-0000-0000-000000000000",
  "commentId": "00000000-0000-0000-0000-000000000000"
}
```