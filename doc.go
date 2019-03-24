/*
Package main will scrape the golint failures Google Sheet for Kubernetes repo issue #68026, and reconcile it against a
list of packages with golint failures from the same repo, to surface a list of packages that still need attention.

Sheet: https://docs.google.com/spreadsheets/d/1VhU6zCk-vaAnbu_XkgsIq5I7jVnezUquWBMhcIN6AJM/edit#gid=0
Failures: https://github.com/kubernetes/kubernetes/raw/master/hack/.golint_failures
*/
package main
