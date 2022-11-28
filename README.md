# ocm-toolbox

Collection of command line tools to help with day to day tasks while working with OCM

## Usage

### Build Binary

To use ocm-toolbox locally, you will need to build the binary. A make
target exists for this:

```shell
make build
```

On completion, a directory named `out` will contain the `ocm-toolbox`
binary.

### Go Install

You can also get ocm-toolbox by using `go install`.

```shell
go install github.com/ryankwilliams/ocm-toolbox@latest
```

### Container Image

You can also use ocm-toolbox using the container image below:

```shell
podman pull ghcr.io/ryankwilliams/ocm-toolbox:main
podman run ghcr.io/ryankwilliams/ocm-toolbox:main --help
```

## Usage

`ocm-toolbox` has a command line interface to it with various sub-commands.
Each sub-command will have different options and a help menu.

```shell
./out/ocm-toolbox --help
```

## Examples

1. List cluster details

```shell
# OCM production
./out/ocm-toolbox cluster-details

# OCM stage
./out/ocm-toolbox cluster-details --url https://api.stage.openshift.com
```

2. Get cluster credentials

```shell
./out/ocm-toolbox cluster-credentials --cluster-id <CLUSTER_ID>
```

3. Update cluster expiration timestamp

```shell
./out/ocm-toolbox set-cluster-expiration --cluster-id <CLUSTER_ID> \
--duration 60
```
