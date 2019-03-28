package maven_test

import (
	"github.com/locrep/locrep-go/config"
	"github.com/locrep/locrep-go/server"
	. "github.com/locrep/locrep-go/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
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

		cmd := exec.Command("mvn", "package", "-Dmaven.repo.remote=http://localhost:8888")
		cmd.Dir = "dummy-repo/"

		result, err := cmd.Output()
		println(string(result))
		Expect(err).Should(BeNil())
	})

	It("should return 200 status ok", func() {
		Expect(recorder.Result().StatusCode).Should(Equal(http.StatusOK))
	})

	It("should return hello world", func() {
		greeting, err := ioutil.ReadAll(recorder.Result().Body)
		Expect(string(greeting)).Should(Equal("hello-world"))
		Expect(err).Should(BeNil())
	})

	AfterAll(func() {
		testServer.Close()
		os.RemoveAll("dummy-repo/target/")
	})
})
