# Terraform Provider Manager

[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://lbesson.mit-license.org/)

Terraform Provider Manager (`tpm`) is a command-line interface (CLI) tool designed to simplify the management of [Terraform](https://www.terraform.io/) providers in the [Terraform plugin cache directory](https://developer.hashicorp.com/terraform/cli/config/config-file#provider-plugin-cache). With `tpm` you can easily install, uninstall, and list providers, helping you to streamline your Terraform workflow.

## Installation

### From source

Install Go if it is not already installed. You can download it from the official [website](https://golang.org/dl).

Clone the Terraform Provider Manager repository:

```shell
git clone https://github.com/Madh93/tpm
cd tpm
```

Build the binary:

```shell
go build -o tpm
```

Move the binary to a directory in your `PATH`:

```shell
mv tpm /usr/local/bin/
```

This will allow you to run `tpm` from any directory in your terminal.

## Usage

```shell
Terraform Provider Manager is a simple CLI to manage Terraform providers in the Terraform plugin cache directory

Usage:
  tpm [flags]
  tpm [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  install     Install a provider
  uninstall   Uninstall a provider

Flags:
  -c, --config string                       config file for tpm
  -d, --debug                               enable debug mode
  -h, --help                                help for tpm
  -p, --terraform-plugin-cache-dir string   the location of the Terraform plugin cache directory (default "/home/user/.terraform.d/plugin-cache")
  -r, --terraform-registry string           the Terraform registry provider hostname (default "registry.terraform.io")
  -v, --version                             version for tpm

Use "tpm [command] --help" for more information about a command.
```

## License

This project is licensed under the [MIT license](LICENSE).
