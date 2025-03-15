package main

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLinkedLists(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Linked List Suite")
}
