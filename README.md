## Betwixt - A LWM2M Client and Server in Go
[![GoDoc](https://godoc.org/github.com/zubairhamed/betwixt?status.svg)](https://godoc.org/github.com/zubairhamed/betwixt)
[![Build Status](https://drone.io/github.com/zubairhamed/betwixt/status.png)](https://drone.io/github.com/zubairhamed/betwixt/latest)
[![Coverage Status](https://coveralls.io/repos/zubairhamed/betwixt/badge.svg?branch=master)](https://coveralls.io/r/zubairhamed/betwixt?branch=master)

## Requirements & Building
Go v1.4 (but should work on 1.3 as well)

## Running
There are a couple of examples in the /examples package.

### basic_client.go
This runs a client with objects and values based on the LWM2M spec examples.
Code for the objects are found in /examples/obj/*

You can either run the Leshan Standalone server and have the client register to it or run basic_server.go and have it register to it.

### basic_server.go
This runs a simple server which accepts registrations and updates.

By default, it runs a CoAP server on port 5683 and a HTTP server on port 8081

### Using

#### Defining a client

#### Defining an object

## Limitations
- No dTLS support.

## LWM2M - Short of it
- Device Management Standard out of OMA
- Lightweight and compact binary protocol based on CoAP
- Targets as light as 8-bit MCUs

## Links
[A primer on LWM2M](http://www.slideshare.net/zdshelby/oma-lightweightm2-mtutorial)
[Specifications and Technical Information](http://technical.openmobilealliance.org/Technical/technical-information/release-program/current-releases/oma-lightweightm2m-v1-0)
[Leshan - A fairly complete Java-based LWM2M implementation](https://github.com/eclipse/leshan)


