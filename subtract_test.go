package main

import (
	"reflect"
	"testing"
)

func TestSubtract(t *testing.T) {
	patterns := []string{
		"cmd/kubeadm/app",
		"cmd/kubelet/app",
		"pkg/apis/*",
		"pkg/capabilities",
	}

	fails := []string{
		"cmd/kubeadm/app/apis/kubeadm/v1beta1",
		"cmd/kubeadm/app/util/system",
		"pkg/apis/abac/latest",
		"pkg/apis/admission",
		"pkg/apis/admissionregistration",
	}

	expected := []string{
		"cmd/kubeadm/app/apis/kubeadm/v1beta1",
		"cmd/kubeadm/app/util/system",
	}

	if !reflect.DeepEqual(subtractPatternsFromFails(patterns, fails), expected) {
		t.Fail()
	}
}
