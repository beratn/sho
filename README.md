# Sho: Yet Another URL Shortener
**sho** is a basic url shortener written in Go. Sho is a small API developed for those who want to shorten links to be used in business. It aims to shorten urls and redirecting requests that coming to the urls. No any additional feature. That's all.

*Contributions and bug reports are welcome.*


## Stack
- PostgreSQL
- Golang
- Mux
- Gorm
- Redis (Soon)

## APIs

### Short a URL

Request:

**POST** /links
```
{
    "target: "https://beratnakip.com"
}

```

Response:

```
{
    "ID": 12,
    "CreatedAt": "2021-08-07T20:02:25.494307616+03:00",
    "UpdatedAt": "2021-08-07T20:02:25.494307616+03:00",
    "DeletedAt": null,
    "target": "http://beratnakip.com",
    "address": "BpLnfg"
}
```
### Redirect

**GET** /{id}

```
http://localhost:8010/9F3qNf
```

## TODO
- Authentication
- Cache Support
- Url Expirity
- Connection Pooling