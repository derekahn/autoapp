# Create A New Cluster

## Google Cloud Platform

This tutorial leverages the [Google Cloud Platform](https://cloud.google.com/) to streamline provisioning of the compute infrastructure required to bootstrap a Kubernetes cluster from the ground up. [Sign up](https://cloud.google.com/free/) for \$300 in free credits.

[Estimated cost](https://cloud.google.com/products/calculator/#id=348b4f56-d024-4832-8799-61410fe460f8) to run this tutorial: $0.132 per hour ($3.19 per day).

### Create a new cluster

##### Create a project with a desired project-id

> Note that the `project-id` must be unique in the 🌎.
> Also Make sure billing is setup for the project

![create a project](../assets/gcp-create-project.png)

### Now let's create a 2 node cluster

Assuming the `gcloud` cli is successfully installed and setup on your machine:

> `hydra` is the name of the cluster (named after the greek mythical creature). You can name the cluster whatever you wish,

```bash
# This could take a while (~5 minutes), grab a ☕
$ gcloud container clusters create hydra --machine-type=n1-standard-2 --num-nodes=2

# If the kubernetes context isn't set automatically
$ kubectl config use-context gke_autoapping_us-west2-a_hydra

# Double check the context
$ kubectl config get-contexts

  CURRENT   NAME                              CLUSTER                           AUTHINFO                          NAMESPACE
            docker-for-desktop                docker-for-desktop-cluster        docker-for-desktop
  *         gke_autoapping_us-west2-a_hydra   gke_autoapping_us-west2-a_hydra   gke_autoapping_us-west2-a_hydra

# You can use kubectl to list the nodes
$ kubectl get nodes

  NAME                                   STATUS    ROLES     AGE       VERSION
  gke-hydra-default-pool-abc12345-abcd   Ready     <none>    9m        v1.11.7-gke.12
  gke-hydra-default-pool-abc12345-efgh   Ready     <none>    9m        v1.11.7-gke.12
```

## Sections

| Previous                                          | Next                                                 |
| ------------------------------------------------- | ---------------------------------------------------- |
| [Installing the Client Tools](01-client-tools.md) | [Deploying Spinnaker to k8s](03-deploy-spinnaker.md) |
