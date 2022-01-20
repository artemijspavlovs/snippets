package test

import (
	"testing"
	"time"

	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
)

func TestLocalIngress(t *testing.T) {
	t.Parallel()

	kubeResourcePath := "./test-app.yml"

	options := k8s.NewKubectlOptions("", "", "foo-bar-app")

	defer k8s.KubectlDelete(t, options, kubeResourcePath)

	k8s.KubectlApply(t, options, kubeResourcePath)

	k8s.WaitUntilIngressAvailable(t, options, "foo.artpav.here", 10, 10*time.Second)
	k8s.WaitUntilIngressAvailable(t, options, "bar.artpav.here", 10, 10*time.Second)

	k8s.GetIngress(t, options, "foo.artpav.here")

	http_helper.HttpGetWithRetry(t, "http://foo.artpav.here", nil, 200, "foo", 30, 10*time.Second)
	http_helper.HttpGetWithRetry(t, "http://bar.artpav.here", nil, 200, "bar", 30, 10*time.Second)
}
