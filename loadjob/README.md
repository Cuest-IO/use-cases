# Loadjob

## This service generates CPU and RAM load.

## Technologies Used

This service has been developed using the Go 18 programming language.

## Usage

To use this service, follow these steps:

1. Install [Golang](https://go.dev/doc/install) 
2. Navigate to the root directory of the project.
3. Set the appropriate environment variables and initiate the service.
```shell
export CPUSleep=100000 # measured in nanoseconds, controls the CPU load of the service. If it is not specified, the service will utilize 100% of the CPU.
export RAM=100 # refers to the amount of memory, measured in megabytes, that you want to allocate for the service.
export Duration=60 # refers to the running interval of the service, measured in seconds. After the specified duration has elapsed, the service will exit. (default 180 seconds)
export ExitCode=0 # the code that the service will return upon completion (0 for success, and any other value indicating an error).

go run main.go
```

Alternatively, you can use this service within a container. [link](https://github.com/Cuest-IO/use-cases/blob/main/images/loadjob/deployment.yaml) 

## Contributing

If you wish to contribute to this service, please submit a pull request with your proposed changes.

