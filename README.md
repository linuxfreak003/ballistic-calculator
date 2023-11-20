# ballistic-calculator

### HTTP/gRPC ballistic-calculator service

`ballistic-calculator` is a simple ballistic calculator service.

It can be used to calculate drop charts.

`ballistic-calculator` uses a `gRPC` server to interact with the resources.

An HTTP proxy is also included that transcodes to/from HTTP and gRPC

The service is described in `pb/service.proto`

### Installation

To install you will need `Go 1.21+`

```
git clone https://github.com/linuxfreak003/ballistic-calculator.git
```

### Build

To build the gRPC server from the project root you can run

```
go build -o bin/ballistic ./ballistic
```

To build the proxy from the project root you can run

```
go build -o bin/proxy ./proxy
```

### Run

To run the gRPC server:

```
./bin/ballistic
```

To run the proxy:

```
./bin/proxy
```

## Usage

`curl` can be used to make simple requests:

```
curl -X POST http://localhost:8080/ballistic/createrifle -d '{"name": "Rifle","sightHeight": 1.5, "barrelTwist": 7, "zeroRange": 100}'
```

After a Scenario has been created (includes rifle, load, and environment) a solution can
be generated:

```
curl -X POST http://localhost:8080/ballistic/solve -d '{"scenarioId": 1, "range": 500}'
```

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.
