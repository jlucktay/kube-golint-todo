# kube-golint-todo

## Purpose

An aid for tackling [kubernetes#68026](https://github.com/kubernetes/kubernetes/issues/68026).

This program will scrape the golint failures Google Sheet for Kubernetes repo issue #68026, and reconcile it against a list of packages with golint failures from the same repo, to surface a list of packages that still need attention.

## Installation

``` shell
go get github.com/jlucktay/kube-golint-todo
```

## Usage

This tool routes requests to the Google Sheets API through a GCP project, so you will need to generate a `credentials.json` file and store it in `${HOME}/.config/kube-golint-todo/`.

Once you start making requests, a `token.json` file will also be generated, which should automatically be stored in the same directory alongside `credentials.json`.

Further details and guidance on this authentication/setup process can be found [here][1].

## See also

Based on [this Go Quickstart][1] for the Google Sheets API.

[1]: https://developers.google.com/sheets/api/quickstart/go
