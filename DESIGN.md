Design
======

## Ping

Ping indicates if the server is working.

| Name | Path    | Method | Content-Type       | Description                         |
|------|---------|--------|--------------------|-------------------------------------|
| Ping | `/ping` | `GET`  | `application/json` | Get Pong with status code 200 if OK |

## Rest Api

- Base path: `/api/rest/v1`
- Content-Type: `application/json`

| Name        | Resource | Response           | Code | Path         | Method   | Description |     
|-------------|----------|--------------------|------|--------------|----------|--------|
| List Posts  | `Post`   | `Collection<Post>` | 200  | `/posts`     | `GET`    | Retrieve a collection of post |
| Get Post    | `Post`   | `Post`             | 200  | `/posts/:id` | `GET`    | Get a specific post |
| Create Post | `Post`   | `Post`             | 200  | `/posts`     | `POST`   | Create a new post |
| Update Post | `Post`   | `Post`             | 200  | `/posts/:id` | `PATCH`  | Update a post |
| Delete Post | `Post`   | `<empty>`          | 204  | `/posts/:id` | `DELETE` | Delete a post |
| List User   | `User`   | `Collection<User>` | 200  | `/users`     | `GET`    | Retrieve a collection of user |
| Get User    | `User`   | `User`             | 200  | `/users/:id` | `GET`    | Get a specific user |
| Create User | `User`   | `User`             | 200  | `/users`     | `POST`   | Create a new user |
| Update User | `User`   | `User`             | 200  | `/users/:id` | `PATCH`  | Update a user |
| Delete User | `User`   | `<empty>`          | 204  | `/users/:id` | `DELETE` | Delete a user |

### Resource

All resources will be represented with the following fields:

| Key               | Type     | Description                             |
|-------------------|----------|-----------------------------------------|
| `created_at`      | `string` | Date of creation in RFC 3339 format |
| `updated_at`      | `string` | Date of updation in RFC 3339 format |

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

- `404`: The resource is not found
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
| `created_at` | `string` | no        | no      | no       | no                        | Date of creation in RFC 3339 format |
| `updated_at` | `string` | no        | no      | no       | no                        | Date of updation in RFC 3339 format |

**Sample:**   

```json
{
  "id": 1,
  "created_at": "2021-10-13T10:52:11.50932133+02:00",
  "updated_at": "2021-10-13T10:52:11.50932133+02:00",
  "content": "lorem ipsum sit dolor set amet..."
}
```

### User

A User represents informations about a user.

| Key             | Type     | Creatable | Mutable | Required | Validation | Description     |                           
|--------------   |----------|-----------|---------|----------|-----------|------------------|
| `id`            | `uint`   | no        | no      | no       | no        |Unique identifier for a `User` resource |
| `last_name`     | `string` | yes       | no      | yes      | `required,min=2,max=20 `| Last name of a `User` resource |
| `first_name`    | `string` | yes       | no      | yes      | `required,min=2,max=20` | First name of a `User` resource |
| `email`         | `string` | yes       | no      | yes      | `required,email`        | Email of a `User` resource |
| `date_of_birth` | `string` | yes       | no      | yes      | no        | Date of birth of a `User` resource |
| `created_at`    | `string` | no        | no      | no       | no        | Date of creation in RFC 3339 format |
| `updated_at`    | `string` | no        | no      | no       | no        | Date of updation in RFC 3339 format |


**Sample:**

```json
{
  "id": 1,
  "created_at": "2021-10-13T16:40:23.591222637+02:00",
  "updated_at": "2021-10-13T16:40:23.591222637+02:00",
  "last_name": "Armand",
  "first_name": "Fanny",
  "email": "fanfantam_33@hotmail.com",
  "date_of_birth": "18/09/1986"
  
}
```