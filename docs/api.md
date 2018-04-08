# API docs

## Authorization

To access routes which require authorization you have to pass the token in the `Authorization` header. You alse have to use the `Bearer` prefix

**Example**

```http
Authorization: Bearer EXAMPLE_TOKEN
```

:closed_lock_with_key: - Authorization is required for this route

## Modules

### Auth

#### `POST /api/auth/users`

Creates a new user. Authorization is not required to use this route, as it will be used to register users.

Required body keys:

* `username` Alphanumerical (with underscores) string (3-20 characters)
* `email` A valid email address
* `password`

Example request body:

```json
{
  "username": "smroot",
  "email": "smroot@smroot.faith",
  "password": "ktoniepracujetensmroot"
}
```

Example response:

```json
{
  "id": "5ac7bf2b1f43a64efbcadd53",
  "username": "ktoniepracujetensmroot",
  "email": "smroot@smroot.faith"
}
```

Errors:

* `400` when email or username is invalid
* `500` (facepalm) when username is duplicate (TODO: fix)

---

#### `GET /api/auth/users/{id}`

Returns the user with `{id}`.

Example response:

```json
{
  "id": "5ac7bf2b1f43a64efbcadd53",
  "username": "ktoniepracujetensmroot",
  "email": "smroot@smroot.faith"
}
```

Errors

* `404` when user does not exist

---

#### `POST /api/auth/sessions/login`

Checks given credidentials and returns a session with a token. Use the `"token"` response key as your `Bearer` token.

Required body keys:

* `username` string
* `password` string

Example request body:

```json
{
  "username": "smroot",
  "password": "ktoniepracujetensmroot"
}
```

Example response:

```json
{
	"active": true,
	"createdOn": "2018-04-08T19:38:09.97023346+02:00",
	"expiresOn": "2018-04-15T19:38:09.970233514+02:00",
	"id": "5aca53811f43a65757f13485",
	"ipAddress": "127.0.0.1:36728",
	"token": "kTYzEaI5TT1IZH5ZEA8J",
	"user": {
		"id": "5ac7bf2b1f43a64efbcadd53",
		"username": "test",
		"email": "srabada@jgfgfjgf.pl"
	},
	"userID": "5ac7bf2b1f43a64efbcadd53"
}
```

Errors:

* `401` Invalid username or password

#### `POST /api/auth/sessions/logout` :closed_lock_with_key:

Destroys the current session. Requires authentication.

Example request body:

```json
{}
```

Example response:

```json
{}
```

Errors

* `401` Authorization required

#### `GET /api/auth/sessions/current-user`

Returns the currently logged in user.

Example response:

```json
{
  "id": "5ac7bf2b1f43a64efbcadd53",
  "username": "ktoniepracujetensmroot",
  "email": "smroot@smroot.faith"
}
```

Errors

* `401` when there is no session token or it is invalid

### Blog

#### `POST /api/blog/posts` :closed_lock_with_key:

Creates new post. Requires authentication.

Required body keys:

* `content` Content string (3-1000 characters)

Example request body:

```json
{
  "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
}
```

Example response:

```json
{
  "id": "5ac7bf2e1f43a64efbcadd54",
  "authorID": "5ac7bf2b1f43a64efbcadd53",
  "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
}
```

Errors

* `401` Authorization required
* `400` Content is invalid

#### `GET /api/blog/posts/new`

Returns 20 new blog posts from all users in chronological order (newest to oldest). The cursor controls from where to start returning posts, to achieve infinite scrolling.

Optional query parameters:

* `after` string - cursor recieved from previous request (`nextCursor`)

Example response:

```json
{
	"nextCursor": "1523123368121000000",
	"posts": [
		{
			"author": {
				"id": "5ac7bf2b1f43a64efbcadd53",
				"username": "test",
				"email": "smroot@jgfgfjgf.pl"
			},
			"authorID": "5ac7bf2b1f43a64efbcadd53",
			"content": "Nazywam się Edward Maxwell i popieram kandydaturę smroota na prezydenta tego kraju!",
			"createdOn": "2018-04-07T19:49:44.295+02:00",
			"id": "5ac904b81f43a64eeb9b2a92"
		},
		{
			"author": {
				"id": "5ac7bf2b1f43a64efbcadd53",
				"username": "test",
				"email": "smroot@jgfgfjgf.pl"
			},
			"authorID": "5ac7bf2b1f43a64efbcadd53",
			"content": "Nazywam się Trevor Paul i popieram kandydaturę smroota na prezydenta tego kraju!",
			"createdOn": "2018-04-07T19:49:42.977+02:00",
			"id": "5ac904b61f43a64eeb9b2a91"
		}
	]
}
```

#### `GET /api/blog/posts/{id}`

Returns the post with `{id}`.

Example response:

```json
{
  "author": {
    "id": "5ac7bf2b1f43a64efbcadd53",
    "username": "test",
    "email": "srabada@jgfgfjgf.pl"
  },
  "authorID": "5ac7bf2b1f43a64efbcadd53",
  "content":
    "Nazywam się Edward Maxwell i popieram kandydaturę smroota na prezydenta tego kraju!",
  "createdOn": "2018-04-07T19:49:44.295+02:00",
  "id": "5ac904b81f43a64eeb9b2a92"
}
```

#### `POST /api/blog/posts/{id}/comments` :closed_lock_with_key:

Creates comment for post with id `{id}`. Requires authentication.

Required body keys:

* `content` Content string (3-1000 characters)

Example request body:

```json
{
  "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
}
```

Example response:

```json
{
  "id": "5ac7bf2e1f43a64efbcadd54",
  "authorID": "5ac7bf2b1f43a64efbcadd53",
  "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
}
```

#### `POST /api/blog/posts/{id}/upvote` :closed_lock_with_key:

Upvotes post with id `{id}`. Requires authentication.

Body content does not matter.

### `POST /api/blog/posts/{id}/downvote` :closed_lock_with_key:

Downvotes post with id `{id}`. Requires authentication.

Body content does not matter.

