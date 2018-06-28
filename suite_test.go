// +build spec

package fpgo

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFpGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "FP Go Suite")
}
