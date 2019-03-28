# autoapp 🤖

An example of CI/CD with a simple go application.

## Run

### 🐳ized

```bash
# creates a new image
$ make build

# runs the image
$ make run
```

### 🐹ized

```bash
# installs deps etc
$ make install

# execute binary
$ ./bin/cmd
```

## Example

![example view](./assets/example.png)

## Environment Variables

```console
# defaults to "8080"
PORT=3000

# defaults to "Gopher"
NAME="Edgar Allan Poe"
```

## Commands

Run `make help` to list available commands:

```console
  $  make help

Choose a command run in gateway:

  install    Install missing dependencies. Builds binary in ./bin
  build      Creates a docker image of the app
  run        Runs the current docker image on port 8080
  clean      Clean build files. Runs `go clean` internally
  fmt        Runs gofmt on all source files
  test       Runs all the tests.
  coverage   Tests code coverage
  missing    Displays lines of code missing from coverage
```
