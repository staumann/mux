package mux

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("mux", func() {
	It("should return a new mux with empty routes", func() {
		mux := GetNewMux()

		Expect(mux).NotTo(BeNil())
		Expect(mux.routes).To(HaveLen(0))
	})
	It("should register a new handler function", func() {
		mux := GetNewMux()

		mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {

		})
		allRoutes := mux.GetAllRoutes()
		Expect(allRoutes).To(HaveLen(1))
		Expect(allRoutes).To(ContainElement("/foo"))
		Expect(mux.HasHandlerForPattern("/foo")).To(BeTrue())
	})
	It("should register a new handler", func() {
		mux := GetNewMux()

		mux.Handle("/handle-pattern", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		}))
		allRoutes := mux.GetAllRoutes()
		Expect(allRoutes).To(HaveLen(1))
		Expect(allRoutes).To(ContainElement("/handle-pattern"))
		Expect(mux.HasHandlerForPattern("/handle-pattern")).To(BeTrue())
	})
})
