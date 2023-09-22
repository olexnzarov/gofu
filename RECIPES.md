# 🥙 Recipes

**Table of contents** 

- [🥙 Recipes](#-recipes)
  - [Create a process that will persist throughout system restarts](#create-a-process-that-will-persist-throughout-system-restarts)
  - [Create a process that will run forever](#create-a-process-that-will-run-forever)
  - [Pass arguments/flags to the process](#pass-argumentsflags-to-the-process)
  - [Create a process with environment variables](#create-a-process-with-environment-variables)
  - [Create a Node.js process](#create-a-nodejs-process)
  - [Rename a process](#rename-a-process)
  - [Stop and remove a process without any prompts](#stop-and-remove-a-process-without-any-prompts)
  - [Make existing process persist throughout system restarts](#make-existing-process-persist-throughout-system-restarts)
- [📎 Arguments](#-arguments)
  - [Global](#global)
  - [gofu run](#gofu-run)


## Create a process that will persist throughout system restarts 

```sh
gofu run --save $COMMAND
```

```sh
gofu run -s $COMMAND
```

## Create a process that will run forever

**Note:** You can set `--restart-delay` to `0`, making the process restart without any delay when it exits. But beware of cases where the process can exit on its startup, making gofu's process manager spam the system with new processes.

```sh
gofu run --save --restart --restart-delay 1s --max-retries 0 $COMMAND
```

```sh
gofu run -s -r --restart-delay 1s --max-retries 0 $COMMAND
```

## Pass arguments/flags to the process

**Note:** Every flag after the `--` will be passed to the process. If you want to pass regular arguments, you can just write them after the `$COMMAND`.

```sh
gofu run $COMMAND -- myargumentvalue --foo=bar -abc
```

```sh
gofu run $COMMAND myargumentvalue foo bar
```

## Create a process with environment variables

```sh
gofu run -e FOO=BAR --env-file=/path/to/file.env $COMMAND
```

## Create a Node.js process

```sh
gofu run node /absolute/path/index.js
```

## Rename a process

```sh
gofu update --name $NEW_NAME $NAME
```

```sh
gofu update -n $NEW_NAME $NAME
```

## Stop and remove a process without any prompts 

```sh
gofu rm --force --always-yes $NAME
```

```sh
gofu rm -f -y $NAME
```

## Make existing process persist throughout system restarts

```sh
gofu update --save $NAME
```

```sh
gofu update -s $NAME
```

# 📎 Arguments

## Global

| Flag            | Default     | Description 
| :---            | :---:       | :---
| --output (-o)   | text        | output format (text, json, prettyjson)
| --timeout       | 60s         | request timeout to daemon

## gofu run

| Flag            | Default     | Description 
| :---            | :---:       | :---
| --name (-n)     | random name | non-empty string
| --save (-s)     | false       | make process persistent
| --env (-e)      | none        | set environment variable
| --env-file      | none        | add file with environment variables
| --restart (-r)  | false       | enable autorestarts
| --max-retries   | 0           | max number of autorestart
| --restart-delay | 1s          | delay between autorestarts
| --directory     | $HOME       | sets current working directory