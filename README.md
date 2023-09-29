<picture><img src=".github/content/banner.png" /></picture>

**gofu** is a process manager for any application you can think of. It allows you to keep processes alive, manage their environment, balance the load, schedule different tasks to run on system startup, and much more. It also provides an option to start a local web interface to manage the processes more easily.

**Important:** Work in progress. Some features may not work correctly or even not exist at the moment. Use at your discretion.

## Usage

We highly encourage you to check out our [recipes](./RECIPES.md) for more detailed descriptions and advanced examples of using gofu.

```sh
gofu run -n $NAME $COMMAND  # to create a named process 
gofu run -s $COMMAND        # to create a persistent process
```

```sh
gofu ps               # to list all processes
gofu ps -o prettyjson # to list all processes in formatted json
gofu inspect $NAME    # to get information about the process
```

```sh
gofu stop $NAME     # to stop the running process
gofu restart $NAME  # to restart the process
gofu rm $NAME       # to remove the process from the daemon
```

## Packages

These packages help you interact with gofu, providing client interfaces and utility functions:

- [gofu](./pkg/gofu)
- [envfmt](./pkg/envfmt)

These libraries were developed for use in gofu, but you can use them in your projects as standalone packages:

- [protomask](https://github.com/olexnzarov/protomask)
- [processinfo](https://github.com/olexnzarov/processinfo) 

## License

gofu is available under the [MIT](./LICENSE) license, allowing for free use, modification, and distribution.

