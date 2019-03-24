# kube-golint-todo

## Purpose

An aid for tackling [kubernetes#68026](https://github.com/kubernetes/kubernetes/issues/68026).

This program will scrape the golint failures Google Sheet for Kubernetes repo issue #68026, and reconcile it against a list of packages with golint failures from the same repo, to surface a list of packages that still need attention.

## See also

Based on [this Go Quickstart](https://developers.google.com/sheets/api/quickstart/go) for the Google Sheets API.
