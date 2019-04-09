# Setting Up Quay

### 1. Select the `+` icon at the top right

> Select `New Repository`

![Select Create New Repo](../assets/quay-new-repo.png)

### 2. Set it up this repo to be `public`

> Public repos are free

![Creating a new repo](../assets/quay-repo.png)

### 3. Select autoapp

> This is assuming you signed up with Github

![Select autoapp](../assets/quay-github-app.png)

### 4. Select yourself

![Select quay user](../assets/quay-org.png)

### 5. Setting up the build triggers

![Configure build trigger](../assets/quay-build-trigger.png)

> This indicates to quay to pull and build on any push to github's `origin` with a new tag

### 6. Configs regarding docker builds and Dockerfile

> Nothing special here (defaults)

![Configure docker](../assets/quay-docker.png)

## Sections

| Previous                                             | Next                                                      |
| ---------------------------------------------------- | --------------------------------------------------------- |
| [Deploying Spinnaker to k8s](03-deploy-spinnaker.md) | [Configuring Spinnaker for CD](05-configure-spinnaker.md) |
