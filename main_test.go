package main_test

import (
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"

	_ "github.com/lurraca/ssh-ngrok"
)

var _ = Describe("ssh-ngrok", func() {
	var (
		session *gexec.Session

		cliCommand struct {
			binary    string
			arguments string
		}
	)

	BeforeEach(func() {
		cliCommand.binary = sshNgrokBinary
	})

	JustBeforeEach(func() {
		var err error

		cmd := exec.Command(cliCommand.binary, cliCommand.arguments)
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		cliCommand.arguments = ""
		session.Kill()
	})

	Context("when all parameters are provided", func() {
		BeforeEach(func() {
			cliCommand.arguments = "-u test-user"
		})

		It("Displays the ssh command to access the host machine", func() {
			Eventually(session).Should(Say("ssh test-user@ngrok.io -p 9000"))
		})
	})

	Context("when the username parameter is missing", func() {
		It("Alert the user that it needs a param params", func() {
			Eventually(session).Should(Say("Missing Username. Correct usage: ssh-ngrok -u <username-on-host>"))
		})
	})

	Context("when the -u flag is empty", func() {
		BeforeEach(func() {
			cliCommand.arguments = "-u "
		})

		It("Alert the user that it needs a param params", func() {
			Eventually(session).Should(Say("Missing Username. Correct usage: ssh-ngrok -u <username-on-host>"))
		})
	})
	Context("when the -u flag is not defined", func() {
		BeforeEach(func() {
			cliCommand.arguments = "-u"
		})

		It("Alert the user that it needs a param params", func() {
			Eventually(session.Err).Should(Say("flag needs an argument: 'u' in -u"))
		})
	})
})
