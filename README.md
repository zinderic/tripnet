# TripNet

A basic tripwire system that will hash all files in given path and send the intel to a server.

In lockdown mode it will alert the server if changes to the filesystem are observed.

# Send data to the gRPC server via Connect using curl

```
curl \
    --header "Content-Type: application/json" \
    --data '{"file_path": "myfile", "file_hash": "myhash"}' \
    http://localhost:8080/tripserv.v1.TripnetService/FileHash
```

# State of the project

It's still WIP and most functionality is not developed yet.
