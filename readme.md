###DSYS5: Replication

This is a distributed system with clients, frontend(s) and replica servers.

Start by starting the replicas, in a terminal from the /replica repository with

```sh
    go run replica.go 5000
```
```sh
    go run replica.go 5001
```
```sh
    go run replica.go 5002
```

Start the frontend/frontends in /frontend on port 5003 with
```sh
    go run frontend.go 5003
```
Just choose a port which is not reserved for the replicas (5000, 5001, 5002).

For the clients, start them in seperate terminals from the /client repository with 

```sh
    go run client.go 1 5003
```
Clients take 2 arguments, the first is an id (integer) of the client, and the next is the frontend it connects to. 
It is possible to have multiple clients connected to different frontends if they exist, e.g. client 1 speaks with frontend 5003, client 2 speaks with frontend 5004 and so on.

Clients issue request from the command line. The system supports 2 request types on the form 
```sh
    bid 200
```
To bid 200 on the auction. And
```sh
    result
```
To query the state of the auction. Calls to "bid" will receive messages with SUCCESS if they made a succesful bid, FAIL if the bid is too low/the auction ended, EXCEPTION if something technical goes wrong. Calls to "result" will return the highest bid and if the action is over, the winner bid.

The first call to bid starts the auction, which will run for 1 min. Unfortunately, the system has no UI for starting and ending auctions. If you want to start a new auction, you have to restart all the replicas so their timeout resets. 

####Logs

Each process has its respective log in /logs. The client's log contain the responses from the requests to the frontend. The frontend logs replies from its request to the replicas, and the replicas log request from the frontend. 

Clients: 1 is the ID, bid returns ack and result returns outcome.
```sh
   Client (1): -> Ack: FAIL
```
Frontends log on the form: 5003 is the frontend's ID, "Ack: FAIL req by 1" is the reply from the frontend's request to a replica.  
```sh
 Frontend (5003): Ack: FAIL req by 1
```
Replicas log: Contains the replica's ID and the request received from the frontend. 
```sh
 Replica (5000): Bid: 21 from 2 (sendSeq 1)
```

All these logs messaged are also combined in a shared log file, name "combined.txt". 


####Crash

To crash a replica, type ctrl + C (Windows) in the terminal. This will be present in the frontend's log as this error:
```sh
 Frontend (5004): Failed to forward request to replica: rpc error: code = Unavailable desc = connection error: desc = "transport: Error while dialing: dial tcp [::1]:5001: connectex: No connection could be made because the target machine actively refused it."
```

