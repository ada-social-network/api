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

| Name     | Resource         | Response              | Code | Path            | Method   | Description               |     
|----------|------------------|--------------------   |------|--------------   |----------|---------------------------|
| Register | `UserRegister`   |    `User`             | 200  | `/register`     | `POST`   |  Register a new user      |
| Login    | `UserLogin`      |    `Token`            | 200  | `/login`        | `POST`   |  Log in and create token  |                   |
| Refresh  | `TokenRefresh`   |    `Token`            | 200  | `/refresh`      | `GET`    |  Refresh existing token   |                  |

### How to register

You can register to the API 
```shell
curl --location --request POST 'http://localhost:8080/auth/register' \
--header 'Content-Type: application/json' \
--data-raw '{
        "last_name": "Fanny",
        "first_name": "Armand",
        "email": "fannyarmand2@gmail.com",
        "password": "secretpassword"
    
}'
```

**Sample:**

```json
{
  "ID": 6,
  "CreatedAt": "2021-11-19T15:59:58.407451298+01:00",
  "UpdatedAt": "2021-11-19T15:59:58.407451298+01:00",
  "DeletedAt": null,
  "last_name": "Baba",
  "first_name": "Ali",
  "email": "ali@gmail.com",
  "date_of_birth": "",
  "apprentice_at": "",
  "profil_pic": "",
  "private_mail": "",
  "instagram": "",
  "facebook": "",
  "github": "",
  "linkedin": "",
  "mbti": "",
  "is_admin": false,
  "promo_id": 0,
  "bda_posts": null,
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

### How to renew a token

You can renew a token which is already existing:


### How to use API authenticated endpoint?

(add in header token, add in cookie token, add in query token and put some curl for example).

## Rest Api

- Base path: `/api/rest/v1`
- Content-Type: `application/json`
- Authentication: `true`
- Rights: `anyone`

| Name            | Resource    | Response              | Code | Path            | Method   | Description |     
|-------------    |----------   |--------------------   |------|--------------   |----------|--------|
| List Posts      | `Post`      | `Collection<Post>`    | 200  | `/posts`        | `GET`    | Retrieve a collection of post |
| Get Post        | `Post`      | `Post`                | 200  | `/posts/:id`    | `GET`    | Get a specific post |
| Create Post     | `Post`      | `Post`                | 200  | `/posts`        | `POST`   | Create a new post |
| Update Post     | `Post`      | `Post`                | 200  | `/posts/:id`    | `PATCH`  | Update a post |
| Delete Post     | `Post`      | `<empty>`             | 204  | `/posts/:id`    | `DELETE` | Delete a post |
| List Users      | `User`      | `Collection<User>`    | 200  | `/users`        | `GET`    | Retrieve a collection of user |
| Get User        | `User`      | `User`                | 200  | `/users/:id`    | `GET`    | Get a specific user |
| Create User     | `User`      | `User`                | 200  | `/users`        | `POST`   | Create a new user |
| Update User     | `User`      | `User`                | 200  | `/users/:id`    | `PATCH`  | Update a user |
| Delete User     | `User`      | `<empty>`             | 204  | `/users/:id`    | `DELETE` | Delete a user |
| List  BdaPosts  | `BdaPost`   | `Collection<BdaPost>` | 200  | `/bdaposts`     | `GET`    | Retrieve a collection of bda post |
| Get BdaPost     | `BdaPost`   | `BdaPost`             | 200  | `/bdaposts/:id` | `GET`    | Get a specific bda post |
| Create  BdaPost | `BdaPost`   | `BdaPost`             | 200  | `/bdaposts`     | `POST`   | Create a new bda post |
| Update  BdaPost | `BdaPost`   | `BdaPost`             | 200  | `/bdaposts/:id` | `PATCH`  | Update a bda post |
| Delete  BdaPost | `BdaPost`   | `<empty>`             | 204  | `/bdaposts/:id` | `DELETE` | Delete a bda post |
| List Promos     | `Promo`     | `Collection<Promo>`   | 200  | `/promos`        | `GET`    | Retrieve a collection of promo |
| Create Promo    | `Promo`     | `Promo`               | 200  | `/promos`        | `POST`   | Create a new promo |
| Update Promo    | `Promo`     | `Promo`               | 200  | `/promos/:id`    | `PATCH`  | Update a promo|
| Delete Promo    | `Promo`     | `<empty>`             | 204  | `/promos/:id`    | `DELETE` | Delete a promo |

### Resource

All resources will be represented with the following fields:

| Key               | Type     | Description                             |
|-------------------|----------|-----------------------------------------|
| `created_at`      | `string` | Date of creation in RFC 3339 format |
| `updated_at`      | `string` | Date of updation in RFC 3339 format |
| `deleted_at`      | `string` | Date of deletion in RFC 3339 format |

### Collection

A collection represent a list of resources. Any resources will be represented
in a common way.

For example:

```json
[
  {
    "id": 1,
    "created_at": "2021-10-13T10:52:11.50932133+02:00",
    "updated_at": "2021-10-13T10:52:11.50932133+02:00",
    "content": "lorem ipsum sit dolor set amet..."
  },
  {
    "id": 2,
    "created_at": "2021-10-13T12:52:11.50932133+02:00",
    "updated_at": "2021-10-13T12:52:11.50932133+02:00",
    "content": "foo bar..."
  }
]
```

### Errors

When a request can not be fulfilled an error will be returned with a status code >= 400
and the following `application/json` content:

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

| Key       | Type     | Description                             |
|-----------|----------|-----------------------------------------|
| `message` | `string` | Dummy message |

**Sample:**

```json
{
  "message": "pong"
}
```

### Post

A Post represents an information.

| Key          | Type     | Creatable | Mutable | Required | Validation                | Description                            |
|-----------   |----------|-----------|---------|----------|---------------------------|---------------------------------------|
| `id`         | `uint`   | no        | no      | no       | no                        | Unique identifier for a `Post` resource |
| `content`    | `string` | yes       | yes     | yes      | `required,min=4,max=1024` | Content of a `Post` resource        |
| `user_id`    | `uint`   | no        | no      | yes      | no                        | User id of a `Post` resource        |
| `created_at` | `string` | no        | no      | no       | no                        | Date of creation in RFC 3339 format |
| `updated_at` | `string` | no        | no      | no       | no                        | Date of updation in RFC 3339 format |
| `deleted_at` | `string` | no        | no      | no       | no                        | Date of deletion in RFC 3339 format |

**Sample:**   

```json
{
  "ID": 1,
  "CreatedAt": "2021-11-05T16:54:49.182599198+01:00",
  "UpdatedAt": "2021-11-05T16:54:49.182599198+01:00",
  "DeletedAt": null,
  "content": "lorem ipsum sit dolor set amet...",
  "user_id": 1
}
```
### BdaPost

A Bda Post represents an information.

| Key          | Type     | Creatable | Mutable | Required | Validation                | Description                            |
|-----------   |----------|-----------|---------|----------|---------------------------|---------------------------------------|
| `id`         | `uint`   | no        | no      | no       | no                        | Unique identifier for a `BdaPost` resource |
| `title`      | `string` | yes       | no      | yes      | `required,min=4,max=1024` | Unique identifier for a `BdaPost` resource |
| `content`    | `string` | yes       | no      | yes      | `required,min=4,max=1024` | Content of a `BdaPost` resource        |
| `user_id`    | `uint`   | no        | no      | yes      | no                        | User id of a `BdaPost` resource        |
| `created_at` | `string` | no        | no      | no       | no                        | Date of creation in RFC 3339 format |
| `updated_at` | `string` | no        | no      | no       | no                        | Date of updation in RFC 3339 format |
| `deleted_at` | `string` | no        | no      | no       | no                        | Date of deletion in RFC 3339 format |

**Sample:**   

```json
{
  "ID": 1,
  "CreatedAt": "2021-11-05T17:04:13.475674216+01:00",
  "UpdatedAt": "2021-11-05T17:04:13.475674216+01:00",
  "DeletedAt": null,
  "title": " Titre",
  "content": "lorem ipsum sit dolor set amet...",
  "user_id": 1
}
```

### User

A User represents informations about a user.

| Key             | Type     | Creatable | Mutable | Required | Validation               | Description     |                           
|--------------   |----------|-----------|---------|----------|--------------------------|------------------|
| `id`            | `uint`   | no        | no      | no       | no                       |Unique identifier for a `User` resource |
| `last_name`     | `string` | yes       | no      | yes      | `required,min=2,max=20 ` | Last name of a `User` resource |
| `first_name`    | `string` | yes       | no      | yes      | `required,min=2,max=20`  | First name of a `User` resource |
| `email`         | `string` | yes       | no      | yes      | `required,email`         | Email of a `User` resource |
| `date_of_birth` | `string` | yes       | no      | yes      | no                       | Date of birth of a `User` resource |
| `apprentice_at` | `string` | yes       | yes     | no       | no                       | Enterprise of a `User` resource |
| `profil_pic`    | `string` | yes       | yes     | no       | no                       | Profil pic of a `User` resource |
| `private_mail`  | `string` | yes       | yes     | no       | no                       | Private email of a `User` resource |
| `instagram`     | `string` | yes       | yes     | no       | no                       | Instagram Page of a `User` resource |
| `facebook`      | `string` | yes       | yes     | no       | no                       | Facebook Page of a `User` resource |
| `github`        | `string` | yes       | yes     | no       | no                       | Github Page of a `User` resource |
| `linkedin`      | `string` | yes       | yes     | no       | no                       | Linkedin Page of a `User` resource |
| `mbti`          | `string` | yes       | no      | no       | no                       | Profil mbti of a `User` resource |
| `is_admin`      | `bool`   | no        | no      | no       | no                       | Profil admin of a `User` resource |
| `promo_id`      | `uint`   | yes       | no      | no       | no                       | Promo id of a `User` resource |
| `bda_posts`      | `Collection<BdaPost>`| no       | no      | no       | no                       | Bda Posts of a `User` resource |
| `posts`          | `Collection<Post>`   | no       | no      | no       | no                       | Posts of a `User` resource |
| `created_at`    | `string` | no        | no      | no       | no                       | Date of creation in RFC 3339 format |
| `updated_at`    | `string` | no        | no      | no       | no                       | Date of updation in RFC 3339 format |
| `deleted_at`    | `string` | no        | no      | no       | no                       | Date of deletion in RFC 3339 format |


**Sample:**

```json
{
  "ID": 1,
  "CreatedAt": "2021-11-05T16:16:26.259246323+01:00",
  "UpdatedAt": "2021-11-05T16:16:26.259246323+01:00",
  "DeletedAt": null,
  "last_name": "Lovelace",
  "first_name": "Ada",
  "email": "lovelace@gmail.com",
  "date_of_birth": "01/01/2020",
  "apprentice_at": "Ada Tech School",
  "profil_pic": "https://www.seekpng.com/png/detail/506-5061704_cool-profile-avatar-picture-cool-picture-for-profile.png",
  "private_mail": "ada@google.com",
  "instagram": "https://www.instagram.com/adatechschool/",
  "facebook": "https://www.facebook.com/AdaTechSchool",
  "github": "https://github.com/ada-social-network/",
  "linkedin": "https://www.linkedin.com/",
  "mbti": "INFP",
  "is_admin": true,
  "promo_id": 1,
  "BdaPost": null
}
```

### Promo

A promo represents informations about a promo.

| Key             | Type     | Creatable | Mutable | Required | Validation               | Description     |                           
|--------------   |----------|-----------|---------|----------|--------------------------|------------------|
| `id`            | `uint`   | no        | no      | no       | no                       | Unique identifier for a `Promo` resource |
| `promo_name`    | `string` | yes       | no      | no       | no                       | Promo name of a `Promo` resource |
| `date_of_start` | `string` | yes       | no      | no       | no                       | Date of start of a `Promo` resource |
| `date_of_end`   | `string` | yes       | no      | no       | no                       | Date of end of a `Promo` resource |
| `biography`     | `string` | yes       | no      | no       | no                       | Biography of a `Promo` resource |
| `created_at`    | `string` | no        | no      | no       | no                       | Date of creation in RFC 3339 format |
| `updated_at`    | `string` | no        | no      | no       | no                       | Date of updation in RFC 3339 format |
| `deleted_at`    | `string` | no        | no      | no       | no                       | Date of deletion in RFC 3339 format |

**Sample:**

```json
{
    "ID": 1,
    "CreatedAt": "2021-11-19T15:15:16.218234962+01:00",
    "UpdatedAt": "2021-11-19T15:15:16.218234962+01:00",
    "DeletedAt": null,
    "promo_name": "Béatrice Worsley",
    "date_of_start": "05/10/2020",
    "date_of_end": "30/06/2021",
    "biography": "La seconde promo qui a vu le jour à l'école Ada Tech School"
}
```
