# networkplan
Document network infrastructure

{{- $virtual_binary_path := "networkplan" }}

## Features

``` text
{{ exec (printf "%s --help" $.binary_path) | trim }}
```

## Plan

``` text
{{ exec (printf "%s plan --help" $.binary_path) | trim }}
```

## Hostsfile

``` text
{{ exec (printf "%s hostsfile --help" $.binary_path) | trim }}
```


## Exports

- IP list (by network)
- hosts file block (see https://github.com/jojomi/io#auto-generate-etchosts for how to merge that into the `/etc/hosts` file)


## Example

### Input

{{ $exampleInputFile := "testdata/example-network.yml" -}}
``` yml
{{ include $exampleInputFile }}
```
[ {{- $exampleInputFile -}} ]( {{- $exampleInputFile -}} )

### Hostsfile

{{ $cmdHostsfile := printf "%s hostsfile --config %s" $.binary_path $exampleInputFile }}

Execute `{{ $cmdHostsfile | replace $.binary_path $virtual_binary_path }}` to get this output:

``` yml
{{ exec $cmdHostsfile | trim }}
```

### Plan

{{- $outputFilename := "docu/output-plan.html" }}
{{- $cmdPlan := printf "%s plan --config %s" $.binary_path $exampleInputFile }}

Execute `{{ $cmdPlan | replace $.binary_path $virtual_binary_path }}` to get this output:

{{ exec $cmdPlan | writeFile $outputFilename -}}

[Click here](https://htmlpreview.github.io/?https://github.com/jojomi/networkplan/blob/master/ {{- $outputFilename -}} )
