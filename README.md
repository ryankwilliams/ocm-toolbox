# ocm-toolbox

Collection of command line tools to help with day to day tasks while working
with OCM

## Install

You can use `ocm-toolbox` by any of the following ways:

### Go Install

```shell
go install github.com/ryankwilliams/ocm-toolbox@main
```

### Container Image

```shell
podman run ghcr.io/ryankwilliams/ocm-toolbox:main --help
```

### Build Binary

```shell
make build
```

On completion, a directory named `out` will contain the `ocm-toolbox`
binary.

## Usage

`ocm-toolbox` has a command line interface to it with various sub-commands.
Each sub-command will have different options and a help menu.

```shell
ocm-toolbox --help
```

## Examples

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
