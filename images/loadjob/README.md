# Loadjob
The service makes CPU and RAM load.

# Run
Use environment variables to setup the service:
- CPUSleep is nanoseconds which control CPU load. If it is not present, it will load 100% CPU.
- RAM is amount of RAM in megabytes which you want to allocate.
- Duration is the running interval in seconds. After the duration service is exit. (default 180 seconds)
- ExitCode is code in which service will finish (0 - success, otherwise error)  

```shell
export CPUSleep=100000
export RAM=100
export Duration=60
export ExitCode=0
go run main.go
```
