<h2 align="center">tomaranai - fast async bomber</h2>

###

<img align="right" height="250" src="https://media.tenor.com/eRwdyxU7MEAAAAAC/shokugeki-no.gif"  />

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
  <img width="0" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="200" alt="docker logo"  />
</div>

###

### project structure:
```
.
├── api
│   ├── email.go
│   └── sms.go
├── cmd
│   └── main.go
├── core
│   ├── email_bomber.go
│   └── sms_bomber.go
├── Dockerfile
├── go.mod
└── README.md
```

## installation

use go get:

```
go get github.com/kenjitheman/tomaranai_bomber
```

use git clone:

```
git clone https://github.com/kenjitheman/tomaranai_bomber
```

## usage

- run it using: go run main.go

```
cd cmd
go run main.go
```

### run it using docker ->

```
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

## contributing

- pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

- please make sure to update tests as appropriate.

## license

- [MIT](https://choosealicense.com/licenses/mit/)

###
