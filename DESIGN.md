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

| Name       | Resource | Collection | Path     | Method | Description |     
|------------|----------|------------|----------|--------|-------------|
| List Posts | `Post`   | `yes`      | `/posts` | `GET`  | Retrieve a collection of post |

### Collection

A collection represent a list of resources. Any resources will be represented
in a common way.

For example:

```json
[
  {
    "id": "123456",
    "content": "lorem ipsum sit dolor set amet..."
  },
  {
    "id": "654321",
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

| Key       | Type     | Description                             |
|-----------|----------|-----------------------------------------|
| `id`      | `string` | Unique identifier for a `Post` resource |
| `content` | `string` | Content of a `Post` resource |

**Sample:**

```json
{
  "id": "123456",
  "content": "lorem ipsum sit dolor set amet..."
}
```