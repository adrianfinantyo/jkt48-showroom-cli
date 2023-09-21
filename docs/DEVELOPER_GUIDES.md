# Developer Setup Guide

If you are the one who want to use this application and not as a developer, this is the ideal setup for you.

## Pre-requisites

- [Go](https://golang.org/doc/install) is required to run the application.
- [VLC Media Player](https://www.videolan.org/vlc/index.html) is required to enable the live streaming feature.

## Installation

1. Install the application using Go:
   ```bash
   $ go install github.com/adrianfinantyo/jkt48-showroom-cli
   ```
2. Open your terminal and type the following command to verify the installation:
   ```bash
    $ jkt48-showroom-cli --version
   ```
   If the installation is successful, you should see the version of the application.
3. Add an alias to your `~/.bashrc` or `~/.zshrc` file (Optional) :
   ```bash
   alias jkt48sr='jkt48-showroom-cli'
   ```
4. Reload your `~/.bashrc` or `~/.zshrc` file:
   ```bash
    $ source ~/.bashrc
   ```
5. Open your terminal and type the following command to verify the installation:
   ```bash
    $ jkt48sr --version
   ```
   If the installation is successful, you should see the version of the application.

## Uninstallation

You can uninstall the application by deleting the binary file.

```bash
$ which jkt48-showroom-cli
$ rm -rf <path/to/jkt48-showroom-cli>
```

## Usage

For the complete usage guide, please refer to the [usage guide](./USAGE.md).
