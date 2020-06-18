# url-short

## Configure

The app is configured with environment variables described below. With docker the `urlshort.env` file is used. Minimal working example for building and running the app is empty file.  
  
| Name                    | Default    | Comment                       |
|-------------------------|------------|-------------------------------|
| URLSHORT_LOGLEVEL       | DEBUG      |                               |
| URLSHORT_LOGFILE        | ""         | Empty for stderr              |
| URLSHORT_LOGFULLREQ     | false      |                               |
| URLSHORT_PORT           | 12321      | Listen port                   |
| URLSHORT_GRACEFULLWAIT  | 15s        |                               |
| URLSHORT_USEDB          | true       | Use redis or inmemory storage |
| URLSHORT_REDISADDR      | redis:6379 |                               |
| URLSHORT_REDISPASSWORD  | ""         |                               |
| URLSHORT_REDISDB        | 0          |                               |
| URLSHORT_REDISRWTIMEOUT | 5s         |                               |
| URLSHORT_BASEURL        | ""         | Base url before hash          |
| URLSHORT_MAXREHASHTRIES | 5          |                               |



## Build

To build image and run linters use
```
docker-compose build
```

To build image without lint uses
```
SKIP_LINTERS=1 docker-compose build
```

## Run 

```
docker-compose pull
docker-compose up -d
```

