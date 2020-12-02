package nonces_test

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func TestNonces(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "nonces")
}

type configuration struct {
	size int
}

func (c *configuration) Size() int {
	return c.size
}

func OnlyIncludeValidURLCharacters() types.GomegaMatcher {
	return &onlyIncludeValidURLCharacters{}
}

type onlyIncludeValidURLCharacters struct{}

func (*onlyIncludeValidURLCharacters) Match(actual interface{}) (bool, error) {
	value, ok := actual.(string)
	if !ok {
		return false, errors.New("OnlyIncludeValidURLCharacters expects string")
	}
	matched, err := regexp.Match("^[A-Za-z0-9-_]*$", []byte(value))
	if err != nil {
		return false, err
	}
	return matched, nil
}

func (*onlyIncludeValidURLCharacters) FailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected\n\t%#v\nto contain only valid URL characters", actual)
}

func (*onlyIncludeValidURLCharacters) NegatedFailureMessage(actual interface{}) string {
	return fmt.Sprintf("Expected\n\t%#v\nto not contain only valid URL characters", actual)
}
