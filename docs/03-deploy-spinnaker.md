# Deploying Spinnaker

Most of this setup and deployment is straight from [GCP's documentation](https://cloud.google.com/solutions/continuous-delivery-spinnaker-kubernetes-engine), with the exception of utilizing GCP's [Container Registry](https://cloud.google.com/container-registry/). But we will use [Google Cloud Storage](https://cloud.google.com/storage/) which is more expensive than [AWS S3](https://aws.amazon.com/s3/) but for the rest of this tutorial we'll just try and stick to GCP as much as possible for convenience.

## Set up the environment

### 1. Create an IAM service account

You create a Cloud Identity and Access Management (Cloud IAM) service account to delegate permissions to Spinnaker, allowing it to store data in Cloud Storage. Spinnaker stores its pipeline data in Cloud Storage to ensure reliability and resiliency. If your Spinnaker deployment unexpectedly fails, you can create an identical deployment in minutes with access to the same pipeline data as the original.

> Basically persists your Spinnaker configurations

```bash
$ gcloud iam service-accounts \
    create spinnaker-account \
    --display-name spinnaker-account

  Created service account [spinnaker-account].
```

### 2. Store the service account email address and your current project ID in environment variables for use in later commands:

```bash
$ export SA_EMAIL=$(gcloud iam service-accounts list \
    --filter="displayName:spinnaker-account" \
    --format='value(email)')

$ export PROJECT=$(gcloud info --format='value(config.project)')

# Check the newly set vars
$ env

  ...
  SA_EMAIL=spinnaker-account@autoapping.iam.gserviceaccount.com
  PROJECT=autoapping
  ...
```

### 3. Bind the storage.admin role to your service account

```bash
$ gcloud projects add-iam-policy-binding $PROJECT \
    --role roles/storage.admin \
    --member serviceAccount:$SA_EMAIL

  Updated IAM policy for project [autoapping].
  bindings:
  - members:
    - serviceAccount:service-xxxxxxxxxxxx@compute-system.iam.gserviceaccount.com
    role: roles/compute.serviceAgent
  - members:
    - serviceAccount:service-xxxxxxxxxxxx@container-engine-robot.iam.gserviceaccount.com
    role: roles/container.serviceAgent
  - members:
    - serviceAccount:xxxxxxxxxxxx-compute@developer.gserviceaccount.com
    - serviceAccount:xxxxxxxxxxxx@cloudservices.gserviceaccount.com
    - serviceAccount:service-xxxxxxxxxxxx@containerregistry.iam.gserviceaccount.com
    role: roles/editor
  - members:
    - user:your.email@gmail.com
    role: roles/owner
  - members:
    - serviceAccount:spinnaker-account@autoapping.iam.gserviceaccount.com
    role: roles/storage.admin
  etag: XXXXXXXXXXX=
  version: 1
```

### 4.Download the service account key. You need this key later when you install Spinnaker and upload the key to GKE.

```bash
# Be inside the projects deploy/ for the generated files to come
$ cd ~/<PATH TO PROJECT>/autoapp/deploy/

$ gcloud iam service-accounts keys create spinnaker-sa.json --iam-account $SA_EMAIL

  created key [xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx] of type [json] as [spinnaker-sa.json] for [spinnaker-account@autoapping.iam.gserviceaccount.com]
```

## Deploying Spinnaker using Helm

Helm should already be installed from section [Installing the Client Tools](01-client-tools.md).

### 1. Grant Tiller, the server side of Helm, the cluster-admin role in your cluster

```bash
$ kubectl create clusterrolebinding user-admin-binding \
  --clusterrole=cluster-admin \
  --user=$(gcloud config get-value account)

  clusterrolebinding.rbac.authorization.k8s.io/user-admin-binding created

$ kubectl create serviceaccount tiller --namespace kube-system

  serviceaccount/tiller created

$ kubectl create clusterrolebinding tiller-admin-binding \
  --clusterrole=cluster-admin \
  --serviceaccount=kube-system:tiller

  clusterrolebinding.rbac.authorization.k8s.io/tiller-admin-binding created

```

### 2. Grant Spinnaker the cluster-admin role so it can deploy resources across all namespaces:

```bash
$ kubectl create clusterrolebinding \
  --clusterrole=cluster-admin \
  --serviceaccount=default:default spinnaker-admin

  clusterrolebinding.rbac.authorization.k8s.io/spinnaker-admin created
```

### 3. Initialize Helm to install Tiller in your cluster:

```bash
$ helm init --service-account=tiller

  $HELM_HOME has been configured at /Users/orion/.helm.

  Tiller (the Helm server-side component) has been installed into your Kubernetes Cluster.

  Please note: by default, Tiller is deployed with an insecure 'allow unauthenticated users' policy.
  To prevent this, run `helm init` with the --tiller-tls-verify flag.
  For more information on securing your installation see: https://docs.helm.sh/using_helm/#securing-your-helm-installation
  Happy Helming!

# update all the things
$ helm update

  Command "update" is deprecated, use 'helm repo update'

  Hang tight while we grab the latest from your chart repositories...
  ...Skip local chart repository
  ...Successfully got an update from the "appscode" chart repository
  ...Successfully got an update from the "stable" chart repository
  Update Complete. ⎈ Happy Helming!⎈
```

### 4. Ensure that Helm is properly installed by running the following command. If Helm is correctly installed, v2.10.0 appears for both client and server.

```bash
$ helm version
  Client: &version.Version{SemVer:"v2.13.1", GitCommit:"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", GitTreeState:"clean"}
  Server: &version.Version{SemVer:"v2.13.1", GitCommit:"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", GitTreeState:"clean"}
```

## Configure Spinnaker

### 1. Create a bucket for Spinnaker to store its pipeline configuration:

```bash
$ export PROJECT=$(gcloud info \
    --format='value(config.project)')

$ export BUCKET=$PROJECT-spinnaker-config

# Double check env variables
$ env

  ...
  BUCKET=autoapping-spinnaker-config
  PROJECT=autoapping
  ...

$ gsutil mb -c regional -l us-central1 gs://$BUCKET

  Creating gs://autoapping-spinnaker-config/...
```

### 2. Set and run `deploy/spinnaker/spinnaker-config.yaml`

```bash
$ export SA_JSON=$(cat spinnaker-sa.json)

$ export PROJECT=$(gcloud info --format='value(config.project)')

$ export BUCKET=$PROJECT-spinnaker-config

$ export QUAY_USER=<USERNAME>

# Don't foget to escape special characters
$ export QUAY_PASS=<PASSWORD>

$ export QUAY_EMAIL=<ADDRESS>@<DOMAIN>.<TLD>

# Double check env variables
$ env

  ...
  SA_JSON={
    "type": "service_account",
    "project_id": "autoapping",
    "private_key_id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "private_key": "-----BEGIN PRIVATE KEY---XXXXXXXXXX==\n-----END PRIVATE KEY-----\n",
    "client_email": "spinnaker-account@autoapping.iam.gserviceaccount.com",
    "client_id": "xxxxxxxxxxxxxxxxxxxxx",
    "auth_uri": "https://accounts.google.com/o/oauth2/auth",
    "token_uri": "https://oauth2.googleapis.com/token",
    "auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
    "client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/spinnaker-account%40autoapping.iam.gserviceaccount.com"
  }
  PROJECT=autoapping
  BUCKET=autoapping-spinnaker-config
  QUAY_USER=groot
  QUAY_PASS=password
  QUAY_EMAIL=groot@gmail.com
  ...
```

## Deploy the Spinnaker Chart

### Use the Helm command-line interface to deploy the chart with your configuration set

```bash
# this could take >=5 minutes, so grab a ☕
$ helm install -n cd stable/spinnaker \
  -f config.yaml \
  --timeout 600 \
  --version 1.8.1 --wait

  NAME:   cd
  LAST DEPLOYED: Fri Mar 29 10:09:55 2019
  NAMESPACE: default
  STATUS: DEPLOYED

  RESOURCES:
  ==> v1/ClusterRoleBinding
  NAME                    AGE
  cd-spinnaker-spinnaker  4m2s

  ==> v1/ConfigMap
  NAME                             DATA  AGE
  cd-spinnaker-additional-scripts  2     4m2s
  cd-spinnaker-halyard-config      3     4m2s

  ==> v1/Pod(related)
  NAME                    READY  STATUS   RESTARTS  AGE
  cd-redis-master-0       1/1    Running  0         4m2s
  cd-spinnaker-halyard-0  1/1    Running  0         4m2s

  ==> v1/RoleBinding
  NAME                  AGE
  cd-spinnaker-halyard  4m2s

  ==> v1/Secret
  NAME                   TYPE    DATA  AGE
  cd-redis               Opaque  1     4m2s
  cd-spinnaker-gcs       Opaque  1     4m2s
  cd-spinnaker-registry  Opaque  1     4m2s

  ==> v1/Service
  NAME                  TYPE       CLUSTER-IP    EXTERNAL-IP  PORT(S)   AGE
  cd-redis-master       ClusterIP  10.15.243.85  <none>       6379/TCP  4m2s
  cd-spinnaker-halyard  ClusterIP  None          <none>       8064/TCP  4m2s

  ==> v1/ServiceAccount
  NAME                  SECRETS  AGE
  cd-spinnaker-halyard  1        4m2s

  ==> v1/StatefulSet
  NAME                  READY  AGE
  cd-spinnaker-halyard  1/1    4m2s

  ==> v1beta2/StatefulSet
  NAME             READY  AGE
  cd-redis-master  1/1    4m2s


  NOTES:
  1. You will need to create 2 port forwarding tunnels in order to access the Spinnaker UI:
    export DECK_POD=$(kubectl get pods --namespace default -l "cluster=spin-deck" -o jsonpath="{.items[0].metadata.name}")
    kubectl port-forward --namespace default $DECK_POD 9000

  2. Visit the Spinnaker UI by opening your browser to: http://127.0.0.1:9000

  To customize your Spinnaker installation. Create a shell in your Halyard pod:

    kubectl exec --namespace default -it cd-spinnaker-halyard-0 bash

  For more info on using Halyard to customize your installation, visit:
    https://www.spinnaker.io/reference/halyard/

  For more info on the Kubernetes integration for Spinnaker, visit:
    https://www.spinnaker.io/reference/providers/kubernetes-v2/
```

### Port forward spinnaker to expose locally

```bash
$ export DECK_POD=$(kubectl get pods --namespace default -l "cluster=spin-deck" \
    -o jsonpath="{.items[0].metadata.name}")

$ kubectl port-forward --namespace default $DECK_POD 8080:9000 >> /dev/null &
```

## Sections

| Previous                                     | Next                                |
| -------------------------------------------- | ----------------------------------- |
| [Create A New Cluster](02-create-cluster.md) | [Setting up Quay](04-setup-quay.md) |
