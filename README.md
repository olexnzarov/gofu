<picture><img src=".github/content/banner.png" /></picture>

**gofu** is a process manager for any application you can think of. It allows you to keep processes alive, manage their environment, balance the load, schedule different tasks to run on system startup, and much more. It also provides an option to start a local web interface to manage the processes more easily.

**Important:** Work in progress. Some features may not work correctly or even not exist at the moment. Use at your discretion.

## Binaries

- [gofu-daemon](./cmd/gofu-daemon/main.go) - a daemon that manages the processes started by it.
- [gofu](./cmd/gofu/main.go) - a command-line utility that lets you control the daemon.
- gofu-web - a web interface that lets you control the daemon.

## Build

As this is a work in progress, there are no detailed instructions on building or running the application locally. But you can look at the [Makefile](./Makefile) to get a good feel for how it all works.

```bash
make build-proto      # to compile protobuf
make run-daemon       # to run gofu-daemon

make install-gofu     # to build and install gofu cli
gofu ps               # then you can use gofu like this
```

## Design

_gofu-daemon_ is an application that manages the processes. It uses gRPC for communication, which should decrease the difficulty of creating alternative clients. Every process started with _gofu_ will be a child process of this daemon. 

_gofu_ is a command-line utility with commands similar to `docker` or `pm2`. It lets you control the running daemon from the terminal.

_gofu-web_ is an optional web interface that lets you enjoy the ease of use of graphical interfaces. It should allow you to do everything that's possible with _gofu_, but with a user-friendly interface. 

### Philosophy

- Modern and widely used technologies. We don't create our own protocols or use some archaic libraries. The interfaces that gofu exposes should be easily accessible and extendable.
- Cross-platform consistency. The behavior should be the same across different operating systems.  
- Ease of use. There should be a recipe for every possible use case that the user can copy and paste. It also entails that we shouldn't expect the user to understand what they just run, so we must build safeguards and protect them from themselves.
- An alternative to dated utilities. We can't stop the user from using `systemd`, but we can offer a cross-platform option that might be a tad easier to use for a newbie or even an experienced user.

## License

gofu is available under the [MIT](./LICENSE) license, allowing for free use, modification, and distribution.

