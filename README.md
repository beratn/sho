# Sho: Serverless URL Shortener
**sho** is a basic url shortener written in Go. Sho is a small API developed for those who want to shorten links to be used in business. It aims to shorten urls and redirecting requests that coming to the urls. No any additional feature. That's all.

*Contributions and bug reports are welcome.*


## Stack
- PostgreSQL
- Golang
- Redis

## Endpoints

### Short a URL

Request:

**POST** /
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

## Development Requirements

* AWS CLI already configured with Administrator permission
* [Docker installed](https://www.docker.com/community-edition)
* [Golang](https://golang.org)
* SAM CLI - [Install the SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

## Setup process

### Installing dependencies & building the target 

In this example we use the built-in `sam build` to automatically download all the dependencies and package our build target.   
Read more about [SAM Build here](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-cli-command-reference-sam-build.html) 

The `sam build` command is wrapped inside of the `Makefile`. To execute this simply run
 
```shell
make
```

### Local development

**Invoking function locally through local API Gateway**

```bash
    sam local start-api --env-vars env.json
```

If the previous command ran successfully you should now be able to hit the following local endpoint to invoke your function `http://localhost:3000`

**SAM CLI** is used to emulate both Lambda and API Gateway locally and uses our `template.yaml` to understand how to bootstrap this environment (runtime, where the source code is, etc.) - The following excerpt is what the CLI will read in order to initialize an API and its routes:

## TODO
- URL Stats
- Url Expirity