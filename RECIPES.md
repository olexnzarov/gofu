# Recipes <!-- omit from toc -->

This page contains descriptions of **gofu** commands and practical examples of how to use them. It aims to help you start cooking with **gofu** with little experience. If you have a use case not specified here, create a [pull request](https://github.com/olexnzarov/gofu/pulls) and help us improve this documentation! 

- [Start a process - gofu run](#start-a-process---gofu-run)
  - [Start a process that will persist throughout system restarts](#start-a-process-that-will-persist-throughout-system-restarts)
  - [Start a process that will run forever](#start-a-process-that-will-run-forever)
  - [Start a process with arguments and/or flags](#start-a-process-with-arguments-andor-flags)
  - [Start a process with environment variables](#start-a-process-with-environment-variables)
  - [Start a Node.js process](#start-a-nodejs-process)
- [List processes - gofu ps (gofu list)](#list-processes---gofu-ps-gofu-list)
  - [List processes as a JSON array](#list-processes-as-a-json-array)
- [Inspect a process - gofu inspect](#inspect-a-process---gofu-inspect)
- [Stop a process - gofu stop](#stop-a-process---gofu-stop)
  - [Stop a process named '80s-anime-collection-backup'](#stop-a-process-named-80s-anime-collection-backup)
- [Restart a process - gofu restart](#restart-a-process---gofu-restart)
- [Remove a process - gofu rm (gofu remove)](#remove-a-process---gofu-rm-gofu-remove)
  - [Stop and remove a process without any prompts](#stop-and-remove-a-process-without-any-prompts)
- [Update a process - gofu update](#update-a-process---gofu-update)
  - [Rename a process](#rename-a-process)
  - [Make existing process persist throughout system restarts](#make-existing-process-persist-throughout-system-restarts)


## Start a process - gofu run

This command starts a new gofu-managed process, it will be created as a child process of gofu's daemon.

```
Usage: gofu run [FLAG ...] COMMAND [ARGUMENT ...]
```
```
Flags: 
      --directory string         set working directory for the process
  -e, --env stringArray          set an environment variable, usage: -e FOO=BAR -e HELLO=WORLD
      --env-file stringArray     read environment variables from a file, usage: --env-file default.env --env-file local.env
  -h, --help                     help for run
      --max-retries uint32       max number of restart tries
  -n, --name string              set the process name
  -r, --restart                  automatically restart a process when it exits
      --restart-delay duration   delay between automatic restarts
  -s, --save                     start the process on system startup
  -o, --output string            output format (text, json, or prettyjson) (default "text")
      --timeout duration         timeout for requests to the daemon (default 1m30s)
```

### Start a process that will persist throughout system restarts

```sh
gofu run --save $COMMAND
```
```sh
gofu run -s $COMMAND
```

### Start a process that will run forever

**Note:** You can set `--restart-delay` to `0`, making the process restart without any delay when it exits. But beware of cases where the process can exit on its startup, making gofu's process manager spam the system with new processes.

```sh
gofu run --save --restart --restart-delay 1s --max-retries 0 $COMMAND
```
```sh
gofu run -s -r --restart-delay 1s --max-retries 0 $COMMAND
```

### Start a process with arguments and/or flags

**Note:** Every flag after the `--` will be passed to the process. If you want to pass regular arguments, you can just write them after the `$COMMAND`.

```sh
gofu run $COMMAND -- myargumentvalue --foo=bar -abc
```
```sh
gofu run $COMMAND myargumentvalue foo bar
```

### Start a process with environment variables

```sh
gofu run -e FOO=BAR --env-file=/path/to/file.env $COMMAND
```

### Start a Node.js process

```sh
gofu run node /absolute/path/index.js
```

## List processes - gofu ps (gofu list)

This command lists all gofu-managed processes.

```
Usage: gofu ps   [FLAG ...] 
       gofu list [FLAG ...]
```
```
Flags:
  -h, --help               help for ps
  -o, --output string      output format (text, json, or prettyjson) (default "text")
      --timeout duration   timeout for requests to the daemon (default 1m30s)
```

### List processes as a JSON array

```sh
gofu ps -o prettyjson
```

## Inspect a process - gofu inspect

This commands provides detailed information about a gofu-managed process.

```
Usage: gofu inspect [FLAG ...] {NAME|PID}
```
```
Flags:
  -h, --help               help for inspect
  -o, --output string      output format (text, json, or prettyjson) (default "text")
      --timeout duration   timeout for requests to the daemon (default 1m30s)
```

## Stop a process - gofu stop

This command stops a gofu-managed process. 

```
Usage: gofu stop [FLAG ...] {NAME|PID}
```
```
Flags:
  -h, --help               help for stop
  -o, --output string      output format (text, json, or prettyjson) (default "text")
      --timeout duration   timeout for requests to the daemon (default 1m30s)
```

### Stop a process named '80s-anime-collection-backup'

```sh
gofu stop 80s-anime-collection-backup
```

## Restart a process - gofu restart

This command restarts a gofu-managed process. 

```
Usage: gofu restart [FLAG ...] {NAME|PID}
```
```
Flags:
  -h, --help               help for restart
  -o, --output string      output format (text, json, or prettyjson) (default "text")
      --timeout duration   timeout for requests to the daemon (default 1m30s)
```

## Remove a process - gofu rm (gofu remove)

```
Usage: gofu rm     [FLAG ...] {NAME|PID}
       gofu remove [FLAG ...] {NAME|PID}
```
```
Flags:
  -h, --help               help for rm
  -o, --output string      output format (text, json, or prettyjson) (default "text")
      --timeout duration   timeout for requests to the daemon (default 1m30s)
```

### Stop and remove a process without any prompts 

```sh
gofu rm --force --always-yes $NAME
```
```sh
gofu rm -f -y $NAME
```

## Update a process - gofu update

This command updates existing gofu-managed process with specified configuration properties.

```
Usage: gofu update [FLAG ...] {NAME|PID}
```
```
Flags:
      --directory string         update the working directory of a process
  -h, --help                     help for update
      --max-retries uint32       update the max number of restart tries (default 1)
  -n, --name string              update the name of a process
  -r, --restart                  update whether a process should be automatically restarted (default true)
      --restart-delay duration   update the delay between automatic restarts
  -s, --save                     update whether a process should be started on system startup
  -o, --output string            output format (text, json, or prettyjson) (default "text")
      --timeout duration         timeout for requests to the daemon (default 1m30s)
```

### Rename a process

```sh
gofu update --name $NEW_NAME $NAME
```
```sh
gofu update -n $NEW_NAME $NAME
```

### Make existing process persist throughout system restarts

```sh
gofu update --save $NAME
```
```sh
gofu update -s $NAME
```