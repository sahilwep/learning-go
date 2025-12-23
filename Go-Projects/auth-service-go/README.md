# Build Walkthrough:

## Docker-Compose for (Postgres & redis)

- Running compose:
```sh
$ docker compose up -d
$ docker ps
```

- Additionally, we can check the running status using `nc -vv localhost PORT` to get the name of the running services...

- To stop all the running services, we can use this command.
```sh
$ docker compose stop
```