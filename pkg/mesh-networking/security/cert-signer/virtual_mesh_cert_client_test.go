package cert_signer_test

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rotisserie/eris"
	mock_kubernetes_core "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/mocks"
	. "github.com/solo-io/go-utils/testutils"
	smh_core_types "github.com/solo-io/service-mesh-hub/pkg/api/core.smh.solo.io/v1alpha1/types"
	smh_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"
	mock_smh_networking "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/mocks"
	smh_networking_types "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1/types"
	container_runtime "github.com/solo-io/service-mesh-hub/pkg/common/container-runtime"
	mock_certgen "github.com/solo-io/service-mesh-hub/pkg/common/csr/certgen/mocks"
	cert_secrets "github.com/solo-io/service-mesh-hub/pkg/common/csr/certgen/secrets"
	cert_signer "github.com/solo-io/service-mesh-hub/pkg/mesh-networking/security/cert-signer"
	k8s_core_types "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	k8s_meta_types "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ = Describe("virtual mesh cert client", func() {
	var (
		ctrl                  *gomock.Controller
		ctx                   context.Context
		secretClient          *mock_kubernetes_core.MockSecretClient
		virtualMeshClient     *mock_smh_networking.MockVirtualMeshClient
		virtualMeshCertClient cert_signer.VirtualMeshCertClient
		rootCertGenerator     *mock_certgen.MockRootCertGenerator
		testErr               = eris.New("hello")
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		ctx = context.TODO()
		secretClient = mock_kubernetes_core.NewMockSecretClient(ctrl)
		virtualMeshClient = mock_smh_networking.NewMockVirtualMeshClient(ctrl)
		rootCertGenerator = mock_certgen.NewMockRootCertGenerator(ctrl)
		virtualMeshCertClient = cert_signer.NewVirtualMeshCertClient(secretClient, virtualMeshClient, rootCertGenerator)
	})

	It("will fail if virtualMesh cannot be found", func() {
		meshRef := &smh_core_types.ResourceRef{
			Name:      "name",
			Namespace: "namespace",
		}
		virtualMeshClient.EXPECT().GetVirtualMesh(ctx, client.ObjectKey{Name: meshRef.Name, Namespace: meshRef.Namespace}).Return(nil, testErr)
		_, err := virtualMeshCertClient.GetRootCaBundle(ctx, meshRef)
		Expect(err).To(HaveOccurred())
		Expect(err).To(HaveInErrorChain(testErr))
	})

	It("will use user provided trust bundle in vm if set", func() {
		meshRef := &smh_core_types.ResourceRef{
			Name:      "name",
			Namespace: "namespace",
		}
		vm := &smh_networking.VirtualMesh{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      meshRef.Name,
				Namespace: meshRef.Namespace,
			},
			Spec: smh_networking_types.VirtualMeshSpec{
				CertificateAuthority: &smh_networking_types.VirtualMeshSpec_CertificateAuthority{
					Type: &smh_networking_types.VirtualMeshSpec_CertificateAuthority_Provided_{
						Provided: &smh_networking_types.VirtualMeshSpec_CertificateAuthority_Provided{
							Certificate: &smh_core_types.ResourceRef{
								Name:      "tb_name",
								Namespace: "tb_namespace",
							},
						},
					},
				},
			},
		}
		virtualMeshClient.EXPECT().GetVirtualMesh(ctx, client.ObjectKey{Name: meshRef.Name, Namespace: meshRef.Namespace}).Return(vm, nil)
		secretClient.
			EXPECT().
			GetSecret(ctx,
				client.ObjectKey{
					Name:      vm.Spec.GetCertificateAuthority().GetProvided().GetCertificate().GetName(),
					Namespace: vm.Spec.GetCertificateAuthority().GetProvided().GetCertificate().GetNamespace()}).
			Return(nil, testErr)
		_, err := virtualMeshCertClient.GetRootCaBundle(ctx, meshRef)
		Expect(err).To(HaveOccurred())
		Expect(err).To(HaveInErrorChain(testErr))
	})

	It("will return proper CA data", func() {
		meshRef := &smh_core_types.ResourceRef{
			Name:      "name",
			Namespace: "namespace",
		}
		vm := &smh_networking.VirtualMesh{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      meshRef.Name,
				Namespace: meshRef.Namespace,
			},
			Spec: smh_networking_types.VirtualMeshSpec{
				CertificateAuthority: &smh_networking_types.VirtualMeshSpec_CertificateAuthority{
					Type: &smh_networking_types.VirtualMeshSpec_CertificateAuthority_Provided_{
						Provided: &smh_networking_types.VirtualMeshSpec_CertificateAuthority_Provided{
							Certificate: &smh_core_types.ResourceRef{
								Name:      "tb_name",
								Namespace: "tb_namespace",
							},
						},
					},
				},
			},
		}
		matchData := &cert_secrets.RootCAData{
			PrivateKey: []byte("private_key"),
			RootCert:   []byte("root_cert"),
		}
		matchSecret := &k8s_core_types.Secret{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      vm.Spec.GetCertificateAuthority().GetProvided().GetCertificate().GetName(),
				Namespace: vm.Spec.GetCertificateAuthority().GetProvided().GetCertificate().GetNamespace(),
			},
			Data: map[string][]byte{
				cert_secrets.RootCertID:       matchData.RootCert,
				cert_secrets.RootPrivateKeyID: matchData.PrivateKey,
			},
			Type: cert_secrets.RootCertSecretType,
		}
		virtualMeshClient.EXPECT().GetVirtualMesh(ctx, client.ObjectKey{Name: meshRef.Name, Namespace: meshRef.Namespace}).Return(vm, nil)
		secretClient.
			EXPECT().
			GetSecret(ctx,
				client.ObjectKey{
					Name:      vm.Spec.GetCertificateAuthority().GetProvided().GetCertificate().GetName(),
					Namespace: vm.Spec.GetCertificateAuthority().GetProvided().GetCertificate().GetNamespace()},
			).
			Return(matchSecret, nil)
		data, err := virtualMeshCertClient.GetRootCaBundle(ctx, meshRef)
		Expect(err).NotTo(HaveOccurred())
		Expect(data).To(Equal(matchData))
	})

	It("will create auto-generated root cert if CertificateAuthority is not user provided", func() {
		meshRef := &smh_core_types.ResourceRef{
			Name:      "name",
			Namespace: "namespace",
		}
		vm := &smh_networking.VirtualMesh{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      meshRef.Name,
				Namespace: meshRef.Namespace,
			},
			Spec: smh_networking_types.VirtualMeshSpec{
				CertificateAuthority: &smh_networking_types.VirtualMeshSpec_CertificateAuthority{
					Type: &smh_networking_types.VirtualMeshSpec_CertificateAuthority_Builtin_{
						Builtin: &smh_networking_types.VirtualMeshSpec_CertificateAuthority_Builtin{},
					},
				},
			},
		}
		virtualMeshClient.EXPECT().GetVirtualMesh(ctx, client.ObjectKey{Name: meshRef.Name, Namespace: meshRef.Namespace}).Return(vm, nil)
		secretClient.
			EXPECT().
			GetSecret(ctx, client.ObjectKey{Name: cert_signer.DefaultRootCaName(vm), Namespace: container_runtime.GetWriteNamespace()}).
			Return(nil, errors.NewNotFound(k8s_core_types.Resource("secret"), "non-extant-secret"))
		expectedRootCaData := &cert_secrets.RootCAData{}
		rootCertGenerator.
			EXPECT().
			GenRootCertAndKey(vm.Spec.GetCertificateAuthority().GetBuiltin()).
			Return(expectedRootCaData, nil)
		secretClient.
			EXPECT().
			CreateSecret(ctx, expectedRootCaData.BuildSecret(cert_signer.DefaultRootCaName(vm), container_runtime.GetWriteNamespace())).
			Return(nil)
		rootCaData, err := virtualMeshCertClient.GetRootCaBundle(ctx, meshRef)
		Expect(err).ToNot(HaveOccurred())
		Expect(rootCaData).To(Equal(expectedRootCaData))
	})

	It("will get auto-generated root cert if CertificateAuthority is not user provided and already exists", func() {
		meshRef := &smh_core_types.ResourceRef{
			Name:      "name",
			Namespace: "namespace",
		}
		vm := &smh_networking.VirtualMesh{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      meshRef.Name,
				Namespace: meshRef.Namespace,
			},
			Spec: smh_networking_types.VirtualMeshSpec{
				CertificateAuthority: &smh_networking_types.VirtualMeshSpec_CertificateAuthority{
					Type: &smh_networking_types.VirtualMeshSpec_CertificateAuthority_Builtin_{
						Builtin: &smh_networking_types.VirtualMeshSpec_CertificateAuthority_Builtin{},
					},
				},
			},
		}
		virtualMeshClient.EXPECT().GetVirtualMesh(ctx, client.ObjectKey{Name: meshRef.Name, Namespace: meshRef.Namespace}).Return(vm, nil)
		expectedRootCaData := &cert_secrets.RootCAData{}
		secretClient.
			EXPECT().
			GetSecret(ctx, client.ObjectKey{Name: cert_signer.DefaultRootCaName(vm), Namespace: container_runtime.GetWriteNamespace()}).
			Return(expectedRootCaData.BuildSecret(cert_signer.DefaultRootCaName(vm), container_runtime.GetWriteNamespace()), nil)
		rootCaData, err := virtualMeshCertClient.GetRootCaBundle(ctx, meshRef)
		Expect(err).ToNot(HaveOccurred())
		Expect(rootCaData).To(Equal(expectedRootCaData))
	})

	It("will default to using builtin cert", func() {
		meshRef := &smh_core_types.ResourceRef{
			Name:      "name",
			Namespace: "namespace",
		}
		vm := &smh_networking.VirtualMesh{
			ObjectMeta: k8s_meta_types.ObjectMeta{
				Name:      meshRef.Name,
				Namespace: meshRef.Namespace,
			},
			Spec: smh_networking_types.VirtualMeshSpec{},
		}
		virtualMeshClient.EXPECT().GetVirtualMesh(ctx, client.ObjectKey{Name: meshRef.Name, Namespace: meshRef.Namespace}).Return(vm, nil)
		expectedRootCaData := &cert_secrets.RootCAData{}
		secretClient.
			EXPECT().
			GetSecret(ctx, client.ObjectKey{Name: cert_signer.DefaultRootCaName(vm), Namespace: container_runtime.GetWriteNamespace()}).
			Return(expectedRootCaData.BuildSecret(cert_signer.DefaultRootCaName(vm), container_runtime.GetWriteNamespace()), nil)
		rootCaData, err := virtualMeshCertClient.GetRootCaBundle(ctx, meshRef)
		Expect(err).ToNot(HaveOccurred())
		Expect(rootCaData).To(Equal(expectedRootCaData))
	})
})
