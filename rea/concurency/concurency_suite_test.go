package concurency_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConcurency(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Concurency Suite")
}
