<h3 align="center">tomaranai - fast, async, multithreaded bomber</h3>

###

<div align="center">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" height="200" alt="go logo"  />
  <img width="0" />
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" height="200" alt="docker logo"  />
</div>

###

### project structure:

```go
├── cmd
│   └── main.go
├── core
│   ├── email_bomber.go
│   ├── generator.go
│   └── sms_bomber.go
├── Dockerfile
├── go.mod
├── LICENSE
└── README.md
```

## installation

```shell
git clone https://github.com/kenjitheman/tomaranai_bomber
```

## usage

- run it:

```shell
cd cmd
go run main.go
```

- or run it using docker:

```shell
docker build -t your_image_name .
docker run -d -p 8080:80 your_image_name
```

## contributing

- pull requests are welcome, for major changes, please open an issue first to
  discuss what you would like to change

## license

- [MIT](https://choosealicense.com/licenses/mit/)
