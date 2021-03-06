package crd_uninstall_test

import (
	"context"

	kubernetes_apiext "github.com/solo-io/external-apis/pkg/api/k8s/apiextensions.k8s.io/v1beta1"
	kubernetes_apiext_providers "github.com/solo-io/external-apis/pkg/api/k8s/apiextensions.k8s.io/v1beta1/providers"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	mock_k8s_extension_clients "github.com/solo-io/external-apis/pkg/api/k8s/apiextensions.k8s.io/v1beta1/mocks"
	"github.com/solo-io/go-utils/testutils"
	crd_uninstall "github.com/solo-io/service-mesh-hub/pkg/common/kube/crd"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/rest"
)

var _ = Describe("Crd Uninstaller", func() {
	var (
		ctx        context.Context
		ctrl       *gomock.Controller
		restConfig = &rest.Config{
			// these fields aren't relevant to anything
			Host:        "example.com",
			BearerToken: "service-account-token",
		}
		crdClientFactoryBuilder = func(crdClient kubernetes_apiext.CustomResourceDefinitionClient) kubernetes_apiext_providers.CustomResourceDefinitionClientFromConfigFactory {
			return func(cfg *rest.Config) (client kubernetes_apiext.CustomResourceDefinitionClient, err error) {
				Expect(cfg).To(Equal(restConfig))
				return crdClient, nil
			}
		}
	)

	BeforeEach(func() {
		ctrl, ctx = gomock.WithContext(context.TODO(), GinkgoT())
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	It("only removes smh CRDs", func() {
		crdClient := mock_k8s_extension_clients.NewMockCustomResourceDefinitionClient(ctrl)

		crd1 := v1beta1.CustomResourceDefinition{
			ObjectMeta: v1.ObjectMeta{
				Name: "test.abc.smh.solo.io",
			},
		}
		crd2 := v1beta1.CustomResourceDefinition{
			ObjectMeta: v1.ObjectMeta{
				Name: "test.def.smh.solo.io",
			},
		}
		crdClient.EXPECT().
			ListCustomResourceDefinition(ctx).
			Return(&v1beta1.CustomResourceDefinitionList{
				Items: []v1beta1.CustomResourceDefinition{
					crd1,
					crd2,
					{
						ObjectMeta: v1.ObjectMeta{
							Name: "unrelated.crd",
						},
					},
				},
			}, nil)

		crdClient.EXPECT().
			GetCustomResourceDefinition(ctx, crd1.GetName()).
			Return(&crd1, nil)

		crdClient.EXPECT().
			DeleteCustomResourceDefinition(ctx, crd1.GetName()).
			Return(nil)

		crdClient.EXPECT().
			GetCustomResourceDefinition(ctx, crd2.GetName()).
			Return(&crd2, nil)

		crdClient.EXPECT().
			DeleteCustomResourceDefinition(ctx, crd2.GetName()).
			Return(nil)

		deletedCrds, err := crd_uninstall.NewCrdRemover(crdClientFactoryBuilder(crdClient)).RemovesmhCrds(ctx, "cluster-1", restConfig)
		Expect(deletedCrds).To(BeTrue())
		Expect(err).NotTo(HaveOccurred())
	})

	It("responds with the appropriate error if the list call fails", func() {
		testErr := eris.New("test-err")
		crdClient := mock_k8s_extension_clients.NewMockCustomResourceDefinitionClient(ctrl)
		crdClient.EXPECT().
			ListCustomResourceDefinition(ctx).
			Return(nil, testErr)

		removedCrds, err := crd_uninstall.NewCrdRemover(crdClientFactoryBuilder(crdClient)).RemovesmhCrds(ctx, "cluster-1", restConfig)
		Expect(removedCrds).To(BeFalse())
		Expect(err).To(testutils.HaveInErrorChain(crd_uninstall.FailedToListCrds(testErr, "cluster-1")))
	})

	It("responds with the appropriate error if the delete call fails", func() {
		testErr := eris.New("test-err")
		crdClient := mock_k8s_extension_clients.NewMockCustomResourceDefinitionClient(ctrl)

		crd := v1beta1.CustomResourceDefinition{
			ObjectMeta: v1.ObjectMeta{
				Name: "test.abc.smh.solo.io",
			},
		}
		crdClient.EXPECT().
			ListCustomResourceDefinition(ctx).
			Return(&v1beta1.CustomResourceDefinitionList{
				Items: []v1beta1.CustomResourceDefinition{
					crd,
					{
						ObjectMeta: v1.ObjectMeta{
							Name: "test.def.smh.solo.io",
						},
					},
					{
						ObjectMeta: v1.ObjectMeta{
							Name: "unrelated.crd",
						},
					},
				},
			}, nil)

		crdClient.EXPECT().
			GetCustomResourceDefinition(ctx, crd.GetName()).
			Return(&crd, nil)

		crdClient.EXPECT().
			DeleteCustomResourceDefinition(ctx, crd.GetName()).
			Return(testErr)

		removedCrds, err := crd_uninstall.NewCrdRemover(crdClientFactoryBuilder(crdClient)).RemovesmhCrds(ctx, "cluster-1", restConfig)
		Expect(removedCrds).To(BeFalse())
		Expect(err).To(testutils.HaveInErrorChain(crd_uninstall.FailedToDeleteCrd(testErr, "cluster-1", "test.abc.smh.solo.io")))
	})

	It("does not return an error if the CRDs have been deleted concurrently in the background", func() {
		crdClient := mock_k8s_extension_clients.NewMockCustomResourceDefinitionClient(ctrl)
		crd1 := v1beta1.CustomResourceDefinition{
			ObjectMeta: v1.ObjectMeta{
				Name: "test.abc.smh.solo.io",
			},
		}
		crd2 := v1beta1.CustomResourceDefinition{
			ObjectMeta: v1.ObjectMeta{
				Name: "test.def.smh.solo.io",
			},
		}
		crdClient.EXPECT().
			ListCustomResourceDefinition(ctx).
			Return(&v1beta1.CustomResourceDefinitionList{
				Items: []v1beta1.CustomResourceDefinition{
					crd1,
					crd2,
					{
						ObjectMeta: v1.ObjectMeta{
							Name: "unrelated.crd",
						},
					},
				},
			}, nil)

		crdClient.EXPECT().
			GetCustomResourceDefinition(ctx, crd1.GetName()).
			Return(nil, errors.NewNotFound(schema.GroupResource{}, "test-name"))
		crdClient.EXPECT().
			GetCustomResourceDefinition(ctx, crd2.GetName()).
			Return(nil, errors.NewNotFound(schema.GroupResource{}, "test-name"))

		removedCrds, err := crd_uninstall.NewCrdRemover(crdClientFactoryBuilder(crdClient)).RemovesmhCrds(ctx, "cluster-1", restConfig)
		Expect(removedCrds).To(BeTrue())
		Expect(err).NotTo(HaveOccurred())
	})
})
