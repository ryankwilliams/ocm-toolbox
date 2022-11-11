# ocm-toolbox

Collection of command line tools to help with day to day tasks while working with OCM

## Build Binary

To use ocm-toolbox locally, you will need to build the binary. A make
target exists for this:

```shell
make build
```

On completion, a directory named `out` will contain the `ocm-toolbox`
binary.

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
