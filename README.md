# onosjob-operator

> The operator for creating a K8s job and using ONOS REST API to create rules or host.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Usage](#usage)
- [Development setup](#development-setup)
- [Contributing](#contributing)

## Features

Please refer http://ONOS-IP:ONOS-UI-PORT/onos/v1/docs  

- hosts : Manage inventory of end-station hosts
    - [x] POST /hosts
- flows : Query and program flow rules
    - [x] POST /flows/{deviceId}
    - [ ] POST /flows

## Getting Started

### Prerequisites

- Access to a Kubernetes v1.11.3+ cluster.
- [go](https://golang.org/dl/) version v1.12+.
- [docker](https://docs.docker.com/install/) version 17.03+.
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) version v1.11.3+.
- [ONOS](https://onosproject.org/) v1.13.5 (Helm chart onos-1.1.0).
- [Helm](https://helm.sh/) v2.9.1+.

**Note**: This guide uses kubeadm v1.13.5 as the local Kubernetes cluster, please refer to [vagrant](vagrant).

### Installation

Register the CRD with the Kubernetes apiserver:  

```sh
$ kubectl create -f deploy/crds/onosjob_v1alpha1_onosjob_crd.yaml
```

Setup RBAC and deploy the onosjob-operator:  

```sh
$ kubectl create -f deploy/service_account.yaml
$ kubectl create -f deploy/role.yaml
$ kubectl create -f deploy/role_binding.yaml
$ kubectl create -f deploy/operator.yaml
```

### Usage

Create a onosjob CR

Edit `deploy/crds/onosjob_v1alpha1_onosjob_cr.yaml`

```sh
kubectl create -f deploy/crds/onosjob_v1alpha1_onosjob_cr.yaml
```

## Development setup

### Install the Operator SDK CLI

On Ubuntu
```
ELEASE_VERSION=v0.8.1
curl -OJL https://github.com/operator-framework/operator-sdk/releases/download/${RELEASE_VERSION}/operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu
chmod +x operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu && sudo cp operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu /usr/local/bin/operator-sdk && rm operator-sdk-${RELEASE_VERSION}-x86_64-linux-gnu
```

Ref: https://github.com/operator-framework/operator-sdk/blob/master/doc/user/install-operator-sdk.md

or `cd vagrant && vagrant up`

### Run command when you change Custom Resource structure

After modify the followings files, 

```
deploy/crds/onosjob_v1alpha1_onosjob_cr.yaml
pkg/apis/onosjob/v1alpha1/onosjob_types.go
```

run the below command.

```sh
$ operator-sdk generate k8s
```

### Run locally outside the cluster

After modify the followings file, 

```
pkg/controller/onosjob/onosjob_controller.go
```

and you want to check the program is fine, please run the below command.

```
$ export OPERATOR_NAME=onosjob-operator
$ operator-sdk up local --namespace=default
```

Ref: https://github.com/operator-framework/operator-sdk/blob/master/doc/user-guide.md#2-run-locally-outside-the-cluster

### Test Custom Resource

Edit `deploy/crds/onosjob_v1alpha1_onosjob_cr.yaml`

Run mininet  

```sh
sudo mn --controller remote,ip=ONOS-IP,port=ONOS-openflow-port
# e.g. sudo mn --controller remote,ip=10.0.2.15,port=31653
```

Create CR  

```sh
kubectl create -f deploy/crds/onosjob_v1alpha1_onosjob_cr.yaml
```

---

**Note**: This project is for Master's degree. The method is not very good, cause the better way is not to create a K8s job to call ONOS REST API but call the ONOS REST API by the operator.

<!--
```
operator-sdk new onosjob-operator --cluster-scoped
operator-sdk add api  --api-version=onosjob.com/v1alpha1 --kind=ONOSJob
```
-->
