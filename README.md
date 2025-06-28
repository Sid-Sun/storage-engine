[# Storage Engine API

An API which can be easily integrated into applications to provide state/storage.

The cryptography is based on [this](https://medium.com/@sidharth.soni525/doing-secure-password-authentication-without-storing-passwords-part-1-7b6024843763) medium article by me - the main difference being here we do CRUD on notes instead of encryption and decryption on asymmetric private keys.

## Usage:

The codebase uses OpenAPI for documentation, consult swagger.json

## Clients

Clients can be generated with swagger codegen, a Go client was created before OpenAPI which should still work:

Golang Client:
- https://github.com/fitant/storage-engine-go
