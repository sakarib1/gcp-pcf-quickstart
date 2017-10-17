package ert_routing_test

import (
	. "omg-cli/certification/environment"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TcpRouter", func() {
	var (
		ert TileQuery
	)
	BeforeSuite(func() {
		ert = Target().OpsManager().MustGetTile("cf")
	})
	XIt("has highly available TCP routing", func() {
		tcpRouter := ert.Resource("tcp_router")
		Expect(tcpRouter.Instances).To(BeNumerically(">=", 3))
	})
})
