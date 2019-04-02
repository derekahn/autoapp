# Installing the Client Tools

## Google Cloud Platform

This tutorial leverages the [Google Cloud Platform](https://cloud.google.com/) to streamline provisioning of the compute infrastructure required to bootstrap a Kubernetes cluster from the ground up. [Sign up](https://cloud.google.com/free/) for \$300 in free credits.

[Estimated cost](https://cloud.google.com/products/calculator/#id=a0e9772e-d50b-40e4-9f7a-13f6916b5321) to run this tutorial: $0.022 per hour ($1.03 per day).

> The compute resources required for this tutorial exceed the Google Cloud Platform free tier.

## Install Gcloud

### [Install the Google Cloud SDK](https://cloud.google.com/sdk/)

Follow the Google Cloud SDK [documentation](https://cloud.google.com/sdk/) to install and configure the `gcloud` command line utility.

Verify the Google Cloud SDK version is 240.0.0 or higher:

```
gcloud version
```

### Set a Default Compute Region and Zone

This tutorial assumes a default compute region and zone have been configured.

If you are using the `gcloud` command-line tool for the first time `init` is the easiest way to do this:

```
gcloud init
```

Otherwise set a default compute region:

> I pick region `us-west2` just because it's the closest to me, you can pick any region you desire

```
gcloud config set compute/region us-west2
```

Set a default compute zone:

```
gcloud config set compute/zone us-west2-a
```

## Install Kubectl

### 🍎

Easist way if you're on an 🍎 💻 is with [homebrew](https://brew.sh/):

```bash
$ brew install kubernetes-cli
```

Else just copy 🍝

```bash
$ curl -o kubectl https://storage.googleapis.com/kubernetes-release/release/v1.12.0/bin/darwin/amd64/kubectl
$ chmod +x kubectl
$ sudo mv kubectl /usr/local/bin/
```

### 🐧

Linux users just copy 🍝

```bash
$ wget https://storage.googleapis.com/kubernetes-release/release/v1.12.0/bin/linux/amd64/kubectl

$ chmod +x kubectl

$ sudo mv kubectl /usr/local/bin/
```

### Otherwise refer to [this](https://kubernetes.io/docs/tasks/tools/install-kubectl/)!

## Install Helm

Binary downloads of the Helm client can be found on the [Releases page](https://github.com/helm/helm/releases/tag/v2.13.1).

If you want to use a package manager:

- [Homebrew](https://brew.sh/) users can use

  ```bash
  $ brew install kubernetes-helm.
  ```

- [Chocolatey](https://chocolatey.org/) users can use:
  ```bash
  $ choco install kubernetes-helm
  ```
- [Scoop](https://scoop.sh/) users can use:
  ```bash
  $ scoop install helm
  ```
- [GoFish](https://gofi.sh/index.html) users can use:
  ```bash
  $ gofish install helm
  ```

## Install Spin

### 🍎

```bash
curl -LO https://storage.googleapis.com/spinnaker-artifacts/spin/$(curl -s https://storage.googleapis.com/spinnaker-artifacts/spin/latest)/darwin/amd64/spin

chmod +x spin

sudo mv spin /usr/local/bin/spin
```

### 🐧

```bash
curl -LO https://storage.googleapis.com/spinnaker-artifacts/spin/$(curl -s https://storage.googleapis.com/spinnaker-artifacts/spin/latest)/linux/amd64/spin

chmod +x spin

sudo mv spin /usr/local/bin/spin
```

## Sections

| Previous                             | Next                                         |
| ------------------------------------ | -------------------------------------------- |
| [Prerequisites](00-prerequisites.md) | [Create A New Cluster](02-create-cluster.md) |
