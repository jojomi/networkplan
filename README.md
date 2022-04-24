# networkplan
Document network infrastructure

## Features

``` text
networkplan

Usage:
  networkplan [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  hostsfile   Export the network hosts in hostsfile format
  plan        Generate a network plan in an HTML file

Flags:
  -h, --help   help for networkplan

Use "networkplan [command] --help" for more information about a command.
```

## Plan

``` text
Generate a network plan in an HTML file

Usage:
  networkplan plan [flags]

Flags:
  -h, --help              help for plan
  -o, --open              output generated document
      --print-all-ipv4s   Also print unused IPv4 addresses
```

## Hostsfile

``` text
Export the network hosts in hostsfile format

Usage:
  networkplan hostsfile [flags]

Flags:
  -h, --help   help for hostsfile
```


## Exports

- IP list (by network)
- hosts file block (see https://github.com/jojomi/io#auto-generate-etchosts for how to merge that into the `/etc/hosts` file)
