package main

import "os"

var configBasePath string
var credentialsJson string
var tokenJson string

func init() {
	configBasePath = os.ExpandEnv("${HOME}/.config/kube-golint-todo/")
	credentialsJson = configBasePath + "credentials.json"
	tokenJson = configBasePath + "token.json"
}
