# autoapp 🤖

[![license](https://img.shields.io/github/license/derekahn/autoapp.svg)](https://github.com/derekahn/autoapp/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/derekahn/autoapp)](https://goreportcard.com/report/github.com/derekahn/autoapp)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/a3b36a7aab4d43088d45f213ecb1a423)](https://www.codacy.com/app/git.derek/autoapp?utm_source=github.com&utm_medium=referral&utm_content=derekahn/autoapp&utm_campaign=Badge_Grade)
[![Codacy Badge](https://api.codacy.com/project/badge/Coverage/a3b36a7aab4d43088d45f213ecb1a423)](https://www.codacy.com/app/git.derek/autoapp?utm_source=github.com&utm_medium=referral&utm_content=derekahn/autoapp&utm_campaign=Badge_Coverage)
[![Build Status](https://cloud.drone.io/api/badges/derekahn/autoapp/status.svg)](https://cloud.drone.io/derekahn/autoapp)
[![Docker Repository on Quay](https://quay.io/repository/derekahn/autoapp/status 'Docker Repository on Quay')](https://quay.io/repository/derekahn/autoapp)
![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/derekahn/autoapp.svg)

An example of CI/CD with a simple go application utilizing a multistage docker build.

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

# forces tests to fail
FAILED=true
```

## Commands

Run `make help` to list available commands:

```console
  $  make help

Choose a command run in autoapp:

  install    Install missing dependencies. Builds binary in ./bin
  build      Creates a docker image of the app
  fail       Forces a failed docker build of the app due to failing tests
  run        Runs the current docker image on port 8080
  clean      Clean build files. Runs `go clean` internally
  fmt        Runs gofmt on all source files
  test       Runs all the tests.
  coverage   Tests code coverage
  missing    Displays lines of code missing from coverage
```

## CI/CD How To 🧙

This tutorial assumes you have access to the [Google Cloud Platform](https://cloud.google.com). While GCP is used for basic infrastructure requirements, the lessons learned in this tutorial can be applied to other platforms.

> It's not a requirement. But reading through this tutorial will be easier and available offline utilizing [vmd](https://www.npmjs.com/package/vmd). It requires [node.js](https://nodejs.org/en/) installed which you can do with `$ brew install node` and then `$ npm install -g vmd`. Then in the root of the project `$ cd ~/<WORK_DIR>/autoapp/ && vmd`.

| Section                                                | Description                                                                                                              |
| ------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------ |
| [Prerequisites](docs/00-prerequisites.md)              | Preface, disclaimers and setting expectations                                                                            |
| [Installing the Client Tools](docs/01-client-tools.md) | Setup and install of CLIs required                                                                                       |
| [Creating A New Cluster](docs/02-create-cluster.md)    | Creating and connecting to a [GCP](https://cloud.google.com/) [GKE](https://cloud.google.com/kubernetes-engine/) cluster |
| [Deploying Spinnaker](docs/03-deploy-spinnaker.md)     | Deploying [Spinnaker](https://www.spinnaker.io/) to our cluster                                                          |
| [Setting up Quay](docs/04-setup-quay.md)               | Step by step configuration for [Quay](https://quay.io)'s hooks                                                           |
| [Setting up Spinnaker](docs/05-setup-spinnaker.md)     | Step by step configuration for spinnaker to automate deployments                                                         |
| [Setting up Drone](docs/06-setup-drone.md)             | Wire up a simple `CI` pre-merge check                                                                                    |
| [Deploying 🚀](docs/07-deploy.md)                      | Submitting a pull-request which on successful merge will trigger our `CI/CD`                                             |
