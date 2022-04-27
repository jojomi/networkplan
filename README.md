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
  -c, --config string   Network config (default "~/.networkplan/network.yml")
  -h, --help            help for networkplan

Use "networkplan [command] --help" for more information about a command.
```

## Plan

``` text
Generate a network plan in an HTML file

Usage:
  networkplan plan [flags]

Flags:
  -h, --help             help for plan
      --print-all-ipv4   Also print unused IPv4 addresses

Global Flags:
  -c, --config string   Network config (default "~/.networkplan/network.yml")
```

## Hostsfile

``` text
Export the network hosts in hostsfile format

Usage:
  networkplan hostsfile [flags]

Flags:
  -h, --help   help for hostsfile

Global Flags:
  -c, --config string   Network config (default "~/.networkplan/network.yml")
```


## Exports

- IP list (by network)
- hosts file block (see https://github.com/jojomi/io#auto-generate-etchosts for how to merge that into the `/etc/hosts` file)


## Example

### Input

``` yml
date: 2022-04-24

networks:
  - name: Home Network
    subnet: 192.168.12.0/24
    domains:
      - ""
      - local
      - .myhome.local
    sub:
      - name: DHCP devices
        subnet: 192.168.12.64/26
        devices:
          - name: albert
            description: My fancy computer
            network: DHCP devices
            hostnames:
              - fancy-computer
              - albert
            ipv4: nw+1
          - name: marie
            description: The Laptop
            network: DHCP devices
            ipv4: nw+2
```
[testdata/example-network.yml](testdata/example-network.yml)

### Hostsfile



Execute `networkplan hostsfile --config testdata/example-network.yml` to get this output:

``` yml
## BEGIN Home Network
## 192.168.12.1 – 192.168.12.254

  ## BEGIN DHCP devices
  ## 192.168.12.65 – 192.168.12.126

    # albert – My fancy computer
    192.168.12.65 fancy-computer albert fancy-computer.local albert.local fancy-computer.myhome.local albert.myhome.local

    # marie – The Laptop
    192.168.12.66 marie marie.local marie.myhome.local

  ## END DHCP devices

## END Home Network
```

### Plan

Execute `networkplan plan --config testdata/example-network.yml` to get this output:

<nil>[Click here](https://htmlpreview.github.io/?https://github.com/jojomi/networkplan/blob/master/docu/output-plan.html)
