package service

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/samber/do"

	mock "hello-do/mocks/store"

	"hello-do/store"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Service Suite")
}

var _ = Describe("GetItems", func() {
	var (
		mockStore *mock.Store
		injector  *do.Injector
	)

	BeforeEach(func() {
		injector = do.New()

		mockStore = mock.NewStore(GinkgoT())
		mockStore.On("HealthCheck").Return(nil)
		mockStore.On("Shutdown").Return(nil)

		do.Provide(injector, func(_ *do.Injector) (store.Store, error) {
			return mockStore, nil
		})

		do.Provide(injector, NewService)
	})

	AfterEach(func() {
		Expect(injector.Shutdown()).To(Succeed())
	})

	When("method is called", func() {
		It("should return items", func(ctx SpecContext) {
			mockItems := []string{"mockA", "mockB"}
			mockStore.On("GetItems").Return(mockItems, nil)

			s, err := do.Invoke[Service](injector)
			Expect(err).ToNot(HaveOccurred())
			Expect(s).ToNot(BeNil())
			Expect(s.HealthCheck()).To(Succeed())
			Expect(do.HealthCheck[Service](injector)).To(Succeed())

			items, err := s.GetItems()
			Expect(err).ToNot(HaveOccurred())
			Expect(items).To(Equal(mockItems))
		})
	})
})
