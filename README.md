# Gobra-CLI

Gobra-CLI is a powerful command-line interface (CLI) tool developed using the Go programming language and the Cobra framework. It provides a versatile and user-friendly way to interact with your applications, perform a wide range of tasks, and streamline your workflow.

## Table of Contents

- [Packages](#packages)
  - [net](#net-package)
  - [info](#info-package)
- [Usage](#usage)
  - [net ping](#net-ping)
  - [info diskUsage](#info-diskusage)
- [Building the CLI](#building-the-cli)

## Packages

This CLI pallet contains two packages:

### net (Network)

The `net` package is designed for network-based commands.

### info (Information)

The `info` package is for information-based commands.

## Usage

Here's how you can use Gobra-CLI to perform various tasks using its packages.

### net ping

The `net` package includes a subcommand called `ping`, which is used to ping any hostname domain to verify its existence and validity. It utilizes the `--url` flag to specify the domain you want to ping.

Example:

```sh
cli-pallet net ping --url google.com
```

### info diskUsage

The `info` package includes a subcommand called `diskUsage`. As the name suggests, it is used to provide information about the disk usage of the current directory.

Example:

```sh
cli-pallet info diskUsage
```

## Building the CLI

To build the Gobra-CLI, use the following command:

```sh
go build
```

This command will compile the Go source code and generate an executable binary for the CLI. After building, you can use the CLI as described in the usage section above.

For more information on how to use Gobra-CLI and its features, please refer to the documentation or explore the source code in the respective packages. Feel free to contribute or report issues on the project's GitHub repository.

It's a basic as of now, will evolve into more interactive.

Enjoy using Gobra-CLI for your command-line tasks and interactions!