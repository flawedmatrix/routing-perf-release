package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCpumonitor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cpumonitor Suite")
}
