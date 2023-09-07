# gofu

**gofu** is a process manager for any application you can think of. It allows you to keep processes alive, manage their environment, balance the load, schedule different tasks to run on system startup, and much more. It also provides an option to start a local web interface to manage the processes more easily.

**Important:** Work in progress. Some features may not work correctly or even not exist at the moment. Use at your discretion.

## Binaries

- [gofu-daemon](./cmd/gofu-daemon/main.go) - a daemon that manages the processes started by it.
- gofu - a command-line utility that lets you control the daemon.
- gofu-web - a web interface that lets you control the daemon.

## Build

As this is a work in progress, there are no detailed instructions on building or running the application locally. But you can look at the [Makefile](./Makefile) to get a good feel for how it all works.

```bash
make build-proto  # to compile protobuf
make run-daemon   # to run gofu-daemon
```

## License

**gofu** is available under the [MIT](./LICENSE) license, allowing for free use, modification, and distribution.

