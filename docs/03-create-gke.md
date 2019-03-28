# Create A New Cluster

## Google Cloud Platform

This tutorial leverages the [Google Cloud Platform](https://cloud.google.com/) to streamline provisioning of the compute infrastructure required to bootstrap a Kubernetes cluster from the ground up. [Sign up](https://cloud.google.com/free/) for \$300 in free credits.

[Estimated cost](https://cloud.google.com/products/calculator/#id=a0e9772e-d50b-40e4-9f7a-13f6916b5321) to run this tutorial: $0.022 per hour ($1.03 per day).

> The compute resources required for this tutorial exceed the Google Cloud Platform free tier.

### Create a new cluster

##### Create a project with a desired project-id

> Make sure billing is setup for the project

![create a project](../assets/gcp-create-project.png)

##### Create a 2 node cluster

> I chose us-west2 just because I always choose California regions (closest to me)

![create a new cluster](../assets/gcp-create-cluster.png)

## Connect to your cluster

Assuming the `gcloud` cli is successfully installed and setup on your machine let's connect to it, by clicking `connect`

![connect to cluster from gcloud cli](../assets/gcp-click-connect.png)

Next copy 🍝 into your terminal

```bash
$ gcloud container clusters get-credentials autoapp --zone us-west2-a --project autoapping

# You can use kubectl to list the nodes
$ kubectl get nodes

NAME                                    STATUS    ROLES     AGE       VERSION
gke-autoapp-default-pool-abcx123-abcd   Ready     <none>    41m       v1.11.7-gke.12
gke-autoapp-default-pool-abcx123-efgh   Ready     <none>    41m       v1.11.7-gke.12
```

| Previous                                               | Next |
| ------------------------------------------------------ | ---- |
| [Installing the Client Tools](02-client-tools.md) |      |
