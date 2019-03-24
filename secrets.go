package main

import "os"

func secretsBasePath() string {
	return os.ExpandEnv("${HOME}/.config/kube-golint-todo/")
}

func credentialsJSON() string {
	return secretsBasePath() + "credentials.json"
}

func tokenJSON() string {
	return secretsBasePath() + "token.json"
}
