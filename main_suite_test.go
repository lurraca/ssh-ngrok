package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

func TestSSHNgrok(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ssh-ngrok Integration Suite")
}

var sshNgrokBinary string

var _ = SynchronizedBeforeSuite(func() []byte {
	binaryPath, err := gexec.Build("github.com/lurraca/ssh-ngrok")
	Expect(err).NotTo(HaveOccurred())

	return []byte(binaryPath)
}, func(binaryPath []byte) {
	sshNgrokBinary = string(binaryPath)
})
