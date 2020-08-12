# NOTES API

A Notes API written in Go to demonstrate CRUD (Create Read Update Delete) on encrypted notes anonymously without storing password in any manner - As an example to using cryptography to facilitate authentication

A proof-of-concept to my Medium article on [Doing secure password authentication without storing passwords](https://medium.com/@sidharth.soni525/doing-secure-password-authentication-without-storing-passwords-part-1-7b6024843763) - the main difference being here we do CRUD on notes instead of encryption and decryption on asymmetric private keys

# Usage:

## To Be Greeted

We all like greetings - This API likes to greet, but it is not pushy about it, it only offers you its heartfelt time-appropriate greetings if you ask for them. 

### Request:

Method: `GET`

URL:
```
http://$host:$port/
```



## To Set A New Note

### Request:

Method: `POST`

URL:
```
http://$host:$port/set
```

BODY:
```json
{
    "id": "mySecretNote",
    "pass": "&now:we@pluto",
    "note": "I am a bufferfly, flying through the sky"
}
```

### Response Success:

HTTP Status: `200`

**If an ID is not supplied or a note already exists with the supplied id**, An 8 character long random ID is generated.


BODY:
```json
{
    "id": "mySecretNote"
}
```

### Response Failure:

A request may fail due to: reasons:

  1. Data was not supplied / improperly supplied
  2. Something unexpected happened on the server-side

If a request fails due to 1, HTTP Status: `400` is returned

If a request fails due to 1, HTTP Status: `500` is returned



## To Get A Note

### Request:

Method: `GET`

URL:
```
http://$host:$port/get
```

BODY:
```json
{
    "id": "mySecretNote",
    "note": "I am a bufferfly, flying through the sky"
}
```

### Response Success:

HTTP Status: `200`

BODY:
```json
{
    "id": "mySecretNote"
}
```

### Response Failure:

A request may fail due to: reasons:

  1. Data was not supplied / improperly supplied
  2. Supplied password was incorrect
  3. Note not found
  4. Something unexpected happened on the server-side

If a request fails due to 1, HTTP Status: `400` is returned

If a request fails due to 2, HTTP Status: `403` is returned

If a request fails due to 3, HTTP Status: `404` is returned

If a request fails due to 4, HTTP Status: `500` is returned



## To Delete A Note

### Request:

Method: `DELETE`

URL:
```
http://$host:$port/delete
```

BODY:
```json
{
    "id": "mySecretNote",
    "pass": "&now:we@pluto"
}
```

### Response Success:

HTTP Status: `200`

BODY:
```json
{
    "id": "mySecretNote"
}
```

### Response Failure:

A request may fail due to: reasons:

  1. Data was not supplied / improperly supplied
  2. Supplied password was incorrect
  3. Note not found
  4. Something unexpected happened on the server-side

If a request fails due to 1, HTTP Status: `400` is returned

If a request fails due to 2, HTTP Status: `403` is returned

If a request fails due to 3, HTTP Status: `404` is returned

If a request fails due to 4, HTTP Status: `500` is returned



## To Update A Note

### Request:

Method: `PUT`

URL:
```
http://$host:$port/update/note
```


BODY:
```json
{
    "id": "mySecretNote",
    "pass": "&now:we@pluto",
    "note": "I am a bufferfly, flying through the sky on Mars"
}
```

Optionally to change the password, a new password may also be supplied, like so:

BODY:
```json
{
    "id": "mySecretNote",
    "pass": "&now:we@pluto",
    "new_pass": "&now:we@moon",
    "note": "I am a bufferfly, flying through the sky on Mars"
}
```

### Response Success:

HTTP Status: `200`

BODY:
```json
{
    "id": "mySecretNote"
}
```

### Response Failure:

A request may fail due to: reasons:

  1. Data was not supplied / improperly supplied
  2. Supplied password was incorrect
  3. Note not found
  4. Something unexpected happened on the server-side

If a request fails due to 1, HTTP Status: `400` is returned

If a request fails due to 2, HTTP Status: `403` is returned

If a request fails due to 3, HTTP Status: `404` is returned

If a request fails due to 4, HTTP Status: `500` is returned



## To Change A Note's Pass

### Request:

Method: `PATCH`

URL:
```
http://$host:$port/update/pas
```


BODY:
```json
{
    "id": "mySecretNote",
    "pass": "&now:we@pluto",
    "newpass": "&now:we@moon"
}
```

### Response Success:

HTTP Status: `200`

BODY:
```json
{
    "id": "mySecretNote"
}
```

### Response Failure:

A request may fail due to: reasons:

  1. Data was not supplied / improperly supplied
  2. Supplied password was incorrect
  3. Note not found
  4. Something unexpected happened on the server-side

If a request fails due to 1, HTTP Status: `400` is returned

If a request fails due to 2, HTTP Status: `403` is returned

If a request fails due to 3, HTTP Status: `404` is returned

If a request fails due to 4, HTTP Status: `500` is returned

## Cheers :)