# networkplan
Document network infrastructure

## Features

``` text
{{ exec (printf "%s --help" $.binary_path) | trim }}
```

## Plan

``` shell
{{ exec (printf "%s plan --help" $.binary_path) | trim }}
```

## Hostsfile

``` shell
{{ exec (printf "%s hostsfile --help" $.binary_path) | trim }}
```


## Exports

- IP list (by network)
- hosts file block (see https://github.com/jojomi/io#auto-generate-etchosts for how to merge that into the `/etc/hosts` file)
