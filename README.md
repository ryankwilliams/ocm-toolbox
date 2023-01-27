# OCM Toolbox (ocm-toolbox)

Command line utility to help with day to day tasks when working with
OCM (OpenShift Cluster Manager).

## Install

You can install/run `ocm-toolbox` by any of the following methods:

1. Go install
  ```
  go install github.com/ryankwilliams/ocm-toolbox@main
  ```

2. Container image
  ```
  podman run ghcr.io/ryankwilliams/ocm-toolbox:main --help
  ```

3. Build binary
  ```
  make build
  ./out/ocm-toolbox
  ```

## Usage

Different sub-commands exist within `ocm-toolbox`. Each sub-command will
have different options and a help menu.

```shell
ocm-toolbox --help
```

### Examples

1. List cluster details

```shell
ocm-toolbox cluster-details
```

2. Get cluster credentials

```shell
ocm-toolbox cluster-credentials --cluster-id <CLUSTER_ID>
```

3. Update cluster expiration timestamp

```shell
ocm-toolbox set-cluster-expiration --cluster-id <CLUSTER_ID> \
--duration 60
```
