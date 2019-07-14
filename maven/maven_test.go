package maven_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"

	"github.com/locrep/go/config"
	"github.com/locrep/go/server"
	. "github.com/locrep/go/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("when doing maven install", func() {
	var (
		testServer *httptest.Server
		recorder   *httptest.ResponseRecorder
	)

	BeforeAll(func() {
		_, err := exec.Command("rm", "-rf", "~/.m2/repository").Output()
		Expect(err).Should(BeNil())

		conf := config.Config()
		testServer = httptest.NewServer(server.NewServer(conf))
		recorder = httptest.NewRecorder()

		cmd := exec.Command("mvn", "package", "-Dmaven.repo.remote=http://localhost:"+conf.Environment.Port())
		cmd.Dir = "dummy-repo/"

		result, err := cmd.Output()
		println(string(result))
		Expect(err).Should(BeNil())
	})

	It("should return 200 status ok", func() {
		Expect(recorder.Result().StatusCode).Should(Equal(http.StatusOK))
	})

	AfterAll(func() {
		testServer.Close()
		os.RemoveAll("dummy-repo/target/")
	})
})
