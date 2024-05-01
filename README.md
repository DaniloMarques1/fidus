# Fidus

It is a cli password manager.

## The Vision

We have a cli app that will communicate with web server to store and retrieve
passwords.

Commands that fidus will have:

- register
- auth
- retrieve
- store
- generate
- delete

register is going to request the user for some information and create a *master*
account, a master account is the one used to store passwords.

auth will be used for authorization, auth will request information about your
master account, after authorization is completed you will then be able to use
the other commands to generate, store and retrieve passwords.

retrieve will be used to retrieve a password based on its key. 

store will be used to store a password, you need provide a key and the password
you want to store.

generate you will be generating a random password for you and then storing it.

Fidus will be breaking down into two applications, the server that will be
responsible for securely storing the master's passwords, and a cli app that
will be able to communicate with the server and issue commands.

## Fidus Server

Fidus server will need to expose the following endpoints:

- /register
- /auth
- /retrieve
- /store
- /generate
- /delete

POST `/register` will have the following request body:

```json
{
    "name": "Master's Name",
    "email": "masters@email",
    "password": "masterspassword"
}
```

The email must be unique, so if the user tries to create an account with an
already used email it should throw a 400 error.

PUT `/auth` must generate a jwt token with a 5 minutes expiration time and will
receive the following request body:

```json
{
    "email": "masters@email",
    "password": "masterspassword"
}
```
