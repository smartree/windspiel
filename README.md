# windspiel

windspiel is a simple canary for detection of suspicious network activity (port scans).
The idea is to deploy it on some machines across your infrastructure and to listen on interesting ports.
Each time a connection to the ports is opened, an event is generated and delivered to you, so you are notified when somebody is trying to map your network.

windspiel doesn't require additional infrastructure, so no subscriptions, Internet access or something else (except you want to send notifications to a foreign system, e.g. mail).
To keep deployment easy, windspiel is platform independent, runs without runtime dependencies and is contained in a single binary.
Administrative access is only required if you plan to listen on ports < 1024.

## Installation

Either grab a binary from the releases page or build your own using `go build` after you cloned the repo.
If you are having problems building your own, the Makefile that is used to build releases is a good starting point for investigations.

## Configuration

There is an example configuration file that serves as its own documentation.

## Extension

Especially the logging capabilities are easy to extend, e.g. in order to log to other formats or directly to your SIEM / whatever.
If you have 2 hours and already saw a few lines of golang in your life, have a look at `fileEventHandler.go` and `mailEventHandler.go` in order to get started with writing your own awesome event handler.

## Questions, Bugs, Features

Please open an issue.

## Legal foo, License

Copyright 2019 by Michael Eder, licensed under Apache 2
