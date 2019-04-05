# autoapp 🤖

[![license](https://img.shields.io/github/license/derekahn/autoapp.svg)](https://github.com/derekahn/autoapp/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/derekahn/autoapp)](https://goreportcard.com/report/github.com/derekahn/autoapp)
![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/derekahn/autoapp.svg)

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

## CI/CD How To 🧙

This tutorial assumes you have access to the [Google Cloud Platform](https://cloud.google.com). While GCP is used for basic infrastructure requirements the lessons learned in this tutorial can be applied to other platforms.

> It's not a requirement. But reading through this tutorial will be easier and available offline utilizing [vmd](https://www.npmjs.com/package/vmd). It requires [node.js](https://nodejs.org/en/) installed which you can do with `$ brew install node` and then `$ npm install -g vmd`. Then in the root of the project `$ cd ~/<WORK_DIR>/autoapp/ && vmd`.

| Section                                                | Description                                                     |
| ------------------------------------------------------ | --------------------------------------------------------------- |
| [Prerequisites](docs/00-prerequisites.md)              | Preface, disclaimers and setting expectations                   |
| [Installing the Client Tools](docs/01-client-tools.md) | Setup and install of CLIs required                              |
| [Creating A New Cluster](docs/02-create-cluster.md)    | Creating and connecting to a GCP GKE cluster                    |
| [Deploying Spinnaker](docs/03-deploy-spinnaker.md)     | Deploying [Spinnaker](https://www.spinnaker.io/) to our cluster |
