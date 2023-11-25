###DSYS5: Replication

This is a distributed system with clients, a frontend and replica servers.

In a terminal, start the 3 replicas from the /replica repository with

```sh
    go run replica.go 5000
```
```sh
    go run replica.go 5001
```
```sh
    go run replica.go 5002
```

Start the frontend in /frontend with
```sh
    go run frontend.go 5003
```

For the clients, start them in seperate terminals from the /client repository with 

```sh
    go run client.go 5003
```





