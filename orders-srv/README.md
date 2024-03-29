# Order Service

This is the Order service

Generated with

```
micro new github.com/songxuexian/gogomicro/orders-srv --namespace=sxx.micro.book --alias=order --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: sxx.micro.book.srv.order
- Type: srv
- Alias: order

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./order-srv
```

Build a docker image
```
make docker
```