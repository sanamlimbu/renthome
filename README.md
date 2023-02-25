# `renthome.com`

A platform for rent seeker, property owners, and property managers which facilitates their use cases. The functionality of app is similar to ```realstate.com.au```. 

### _Dev dependencies_

- [Go](https://go.dev/dl/)
- [Makefile](https://www.gnu.org/software/make/)
- [Docker](https://docs.docker.com/engine/install/)

### Envars
```
    RENTHOME_MAIL_HOST=
    RENTHOME_MAIL_PORT=
    RENTHOME_MAIL_USERNAME=
    RENTHOME_MAIL_PASSWORD
    
    RENTHOME_GOOGLE_CLIENT_ID=
    RENTHOME_FACEBOOK_CLIENT_ID=
    RENTHOME_FACEBOOK_CLIENT_ID=
```
### Init
Spin up the project for development
```
$ make init
```
### Start server

Start server listening on port 8000

```
$ make serve
```

### Start web client
Start web client on port 3000

```
$ make web-watch
```
