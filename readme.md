#  PromQL & MetricsQL Prettify!

Make your PromQL or MetricsQL beautiful.

<!-- vim-markdown-toc GFM -->

- [Usage](#usage)
  - [Online Tools](#online-tools)
  - [Command-line](#command-line)
  - [Use xbin.io (Without installation!)](#use-xbinio-without-installation)
- [Thanks](#thanks)

<!-- vim-markdown-toc -->

## Usage

### Online Tools

https://laixintao.github.io/promql-metricsql-prettify/

![](./assets/promql-metricsql-prettify-online-demo.png)

### Command-line

Installation:

```shell
go install github.com/laixintao/promql-metricsql-prettify@latest
```

Usage: pass your PromQL or MetricsQL into `promql-metricsql-prettify` as stdin:

```shell
$ echo 'count(sum(label_replace(node_uname_info, "kernel", "$1", "release", "([0-9]+.[0-9]+.[0-9]+).*")) by (kernel)) > 1' | promql-metricsql-prettify
count(
  sum(
    label_replace(
      node_uname_info,
      "kernel",
      "$1",
      "release",
      "([0-9]+.[0-9]+.[0-9]+).*"
    )
  ) by(kernel)
)
  >
1
```

### Use xbin.io (Without installation!)

[xbin.io](https://xbin.io/) is a collection of cli tools, you can send your
input via HTTP Request (cURL) to xbin.io, and xbin.io will run the command, then return
output via HTTP Response. (You can think it as... serverless shell?)

Use this tool on: https://xbin.io/w/tool/promql-metricsql-prettify

Use it through cURL from terminal:

```shell
$ echo 'count(sum(label_replace(node_uname_info, "kernel", "$1", "release", "([0-9]+.[0-9]+.[0-9]+).*")) by (kernel)) > 1' | \
    curl -X POST --data-binary @- https://xbin.io/promql-metricsql-prettify
count(
  sum(
    label_replace(
      node_uname_info,
      "kernel",
      "$1",
      "release",
      "([0-9]+.[0-9]+.[0-9]+).*"
    )
  ) by(kernel)
)
  >
1
```

How does it work?

1. `|` pip will send your PromQL into curl
2. curl will send the conent to xbin.io/promql-metricsql-prettify
3. xbin.io will run `promql-metricsql-prettify`
4. You get the result from cURL

## Thanks

- This project was inspired by https://github.com/jiacai2050/promql-prettier
- This project uses [VictoriaMetrics](https://github.com/VictoriaMetrics)'s code, [`prettier.go`](https://github.com/VictoriaMetrics/metricsql/blob/master/prettifier.go) in [MetricsQL](https://github.com/VictoriaMetrics/metricsql/)
