package istio

import (
	"context"

	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/supergloo/pkg/translator/istio/plugins"

	"github.com/solo-io/supergloo/pkg/api/external/istio/networking/v1alpha3"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/supergloo/test/inputs"

	v1 "github.com/solo-io/supergloo/pkg/api/v1"
)

var _ = Describe("appliesToDestination", func() {
	Context("upstream selector match", func() {
		It("returns true", func() {
			applies, err := appliesToDestination("details.default.svc.cluster.local", &v1.PodSelector{
				SelectorType: &v1.PodSelector_UpstreamSelector_{
					UpstreamSelector: &v1.PodSelector_UpstreamSelector{
						Upstreams: []core.ResourceRef{
							{Name: "default-details-v1-9080", Namespace: "gloo-system"},
						},
					},
				},
			}, inputs.BookInfoUpstreams())
			Expect(err).NotTo(HaveOccurred())
			Expect(applies).To(BeTrue())
		})
	})
	Context("namespace selector match", func() {
		It("returns true", func() {
			applies, err := appliesToDestination("details.default.svc.cluster.local", &v1.PodSelector{
				SelectorType: &v1.PodSelector_NamespaceSelector_{
					NamespaceSelector: &v1.PodSelector_NamespaceSelector{
						Namespaces: []string{"default"},
					},
				},
			}, inputs.BookInfoUpstreams())
			Expect(err).NotTo(HaveOccurred())
			Expect(applies).To(BeTrue())
		})
	})
	Context("label selector match", func() {
		It("returns true", func() {
			applies, err := appliesToDestination("details.default.svc.cluster.local", &v1.PodSelector{
				SelectorType: &v1.PodSelector_LabelSelector_{
					LabelSelector: &v1.PodSelector_LabelSelector{
						LabelsToMatch: map[string]string{"version": "v1", "app": "details"},
					},
				},
			}, inputs.BookInfoUpstreams())
			Expect(err).NotTo(HaveOccurred())
			Expect(applies).To(BeTrue())
		})
	})
})

var _ = Describe("labelSetsFromSelector", func() {
	Context("PodSelector_UpstreamSelector", func() {
		It("returns labels for each upstream found", func() {
			labelSets, err := labelSetsFromSelector(&v1.PodSelector{
				SelectorType: &v1.PodSelector_UpstreamSelector_{
					UpstreamSelector: &v1.PodSelector_UpstreamSelector{
						Upstreams: []core.ResourceRef{
							{Name: "default-details-v1-9080", Namespace: "gloo-system"},
							{Name: "default-reviews-v2-9080", Namespace: "gloo-system"},
							{Name: "default-reviews-9080", Namespace: "gloo-system"},
						},
					},
				},
			}, inputs.BookInfoUpstreams())
			Expect(err).NotTo(HaveOccurred())
			Expect(labelSets).To(Equal([]map[string]string{
				{"version": "v1", "app": "details"},
				{"app": "reviews", "version": "v2"},
				{"app": "reviews"},
			}))
		})
	})
	Context("PodSelector_NamespaceSelector", func() {
		It("returns labels for each upstream in the namespace", func() {
			labelSets, err := labelSetsFromSelector(&v1.PodSelector{
				SelectorType: &v1.PodSelector_NamespaceSelector_{
					NamespaceSelector: &v1.PodSelector_NamespaceSelector{
						Namespaces: []string{"default"},
					},
				},
			}, inputs.BookInfoUpstreams())
			Expect(err).NotTo(HaveOccurred())
			Expect(labelSets).To(Equal([]map[string]string{
				{"app": "details"},
				{"version": "v1", "app": "details"},
				{"app": "productpage"},
				{"app": "productpage", "version": "v1"},
				{"app": "ratings"},
				{"app": "ratings", "version": "v1"},
				{"app": "reviews"},
				{"version": "v1", "app": "reviews"},
				{"app": "reviews", "version": "v2"},
				{"version": "v3", "app": "reviews"},
			}))
		})
	})
})

var _ = Describe("convertMatcher", func() {
	It("converts a gloo match to an istio match", func() {
		istioMatch := convertMatcher(map[string]string{"app": "details", "version": "v1"}, 1234, &gloov1.Matcher{
			PathSpecifier: &gloov1.Matcher_Exact{
				Exact: "hi",
			},
			Methods: []string{"GET", "ME", "OUTTA", "HERE"},
			Headers: []*gloov1.HeaderMatcher{
				{Name: "k", Value: "v", Regex: true},
				{Name: "a", Value: "z", Regex: false},
			},
		})
		Expect(istioMatch).To(Equal(&v1alpha3.HTTPMatchRequest{
			Uri: &v1alpha3.StringMatch{
				MatchType: &v1alpha3.StringMatch_Exact{Exact: "hi"},
			},
			Method: &v1alpha3.StringMatch{
				MatchType: &v1alpha3.StringMatch_Regex{Regex: "GET|ME|OUTTA|HERE"},
			},
			Headers: map[string]*v1alpha3.StringMatch{
				"a": {MatchType: &v1alpha3.StringMatch_Exact{Exact: "z"}},
				"k": {MatchType: &v1alpha3.StringMatch_Regex{Regex: "v"}},
			},
			Port:         1234,
			SourceLabels: map[string]string{"app": "details", "version": "v1"},
		}))
	})
})

var _ = Describe("createIstioMatcher", func() {
	Context("yes matcher, yes source labels", func() {
		It("creates a copy of each matcher for each set of source labels", func() {
			istioMatchers := createIstioMatcher(
				[]map[string]string{
					{"app": "details"},
					{"app": "reviews"},
				}, 1234, []*gloov1.Matcher{
					{
						PathSpecifier: &gloov1.Matcher_Exact{
							Exact: "hi",
						},
					},
					{
						PathSpecifier: &gloov1.Matcher_Exact{
							Exact: "bye",
						},
					},
				})
			Expect(istioMatchers).To(Equal([]*v1alpha3.HTTPMatchRequest{
				{
					Uri: &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Exact{Exact: "hi"},
					},
					Port:         1234,
					SourceLabels: map[string]string{"app": "details"},
				},
				{
					Uri: &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Exact{Exact: "hi"},
					},
					Port:         1234,
					SourceLabels: map[string]string{"app": "reviews"},
				},
				{
					Uri: &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Exact{Exact: "bye"},
					},
					Port:         1234,
					SourceLabels: map[string]string{"app": "details"},
				},
				{
					Uri: &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Exact{Exact: "bye"},
					},
					Port:         1234,
					SourceLabels: map[string]string{"app": "reviews"},
				},
			}))
		})
	})
})

type testRoutingPlugin struct {
	collectedRoutes []*v1alpha3.HTTPRoute
}

func (t *testRoutingPlugin) Init(params plugins.InitParams) error {
	return nil
}

func (t *testRoutingPlugin) ProcessRoute(params plugins.Params, in v1.RoutingRuleSpec, out *v1alpha3.HTTPRoute) error {
	t.collectedRoutes = append(t.collectedRoutes, out)
	return nil
}

var _ = Describe("createRoute", func() {
	Context("with a route plugin", func() {
		It("creates an http route with the corresponding destination, and calls the plugin for each route", func() {
			resourceErrs := make(reporter.ResourceErrors)
			plug := testRoutingPlugin{}
			t := NewTranslator("hi", []plugins.Plugin{&plug}).(*translator)
			route := t.createRoute(
				plugins.Params{Ctx: context.TODO()},
				"details.default.svc.cluster.local",
				inputs.BookInfoRoutingRules("namespace-where-rules-crds-live"),
				createIstioMatcher(
					[]map[string]string{
						{"app": "details"},
						{"app": "reviews"},
					}, 1234, []*gloov1.Matcher{
						{
							PathSpecifier: &gloov1.Matcher_Exact{
								Exact: "hi",
							},
						},
						{
							PathSpecifier: &gloov1.Matcher_Exact{
								Exact: "bye",
							},
						},
					}),
				inputs.BookInfoUpstreams(),
				resourceErrs,
			)
			Expect(route.Route).To(HaveLen(1))
			Expect(route.Route[0].Destination.Host).To(Equal("details.default.svc.cluster.local"))
			Expect(plug.collectedRoutes).To(HaveLen(1))
			Expect(plug.collectedRoutes[0]).To(Equal(route))
		})
	})
})

type destinationRule struct {
	host    string
	subsets []subset
}

type subset struct {
	name   string
	labels map[string]string
}
type virtualService struct {
	host   string
	routes []route
}

type route struct {
	matches []match
}

type match struct {
	sourceLabels map[string]string
	port         uint32
}

var _ = Describe("Translator", func() {
	It("translates a snapshot into a corresponding meshconfig, returns ResourceErrors", func() {
		plug := testRoutingPlugin{}
		inputRoutingRules := inputs.BookInfoRoutingRules("namespace-where-rules-crds-live")

		t := NewTranslator("hi", []plugins.Plugin{&plug}).(*translator)
		meshConfig, resourceErrs, err := t.Translate(context.TODO(), &v1.ConfigSnapshot{
			Upstreams:    map[string]gloov1.UpstreamList{"": inputs.BookInfoUpstreams()},
			Routingrules: map[string]v1.RoutingRuleList{"": inputRoutingRules},
		})
		Expect(err).NotTo(HaveOccurred())
		Expect(meshConfig).NotTo(BeNil())
		Expect(meshConfig.DesinationRules).To(HaveLen(4))
		// we should have 1 subset for the service selector followed by 1 subset for each set of tags
		for i, expected := range []destinationRule{
			{"details.default.svc.cluster.local",
				[]subset{
					{"app-details",
						map[string]string{"app": "details"}},
					{"app-details-version-v1",
						map[string]string{"app": "details", "version": "v1"}},
				}},
			{"productpage.default.svc.cluster.local",
				[]subset{
					{"app-productpage",
						map[string]string{"app": "productpage"}},
					{"app-productpage-version-v1",
						map[string]string{"app": "productpage", "version": "v1"}},
				}},
			{"ratings.default.svc.cluster.local",
				[]subset{
					{"app-ratings",
						map[string]string{"app": "ratings"}},
					{"app-ratings-version-v1",
						map[string]string{"app": "ratings", "version": "v1"}},
				}},
			{"reviews.default.svc.cluster.local",
				[]subset{
					{"app-reviews",
						map[string]string{"app": "reviews"}},
					{"app-reviews-version-v1",
						map[string]string{"app": "reviews", "version": "v1"}},
					{"app-reviews-version-v2",
						map[string]string{"app": "reviews", "version": "v2"}},
					{"app-reviews-version-v3",
						map[string]string{"app": "reviews", "version": "v3"}},
				}},
		} {
			Expect(meshConfig.DesinationRules[i].Host).To(Equal(expected.host))
			for j, sub := range expected.subsets {
				Expect(meshConfig.DesinationRules[i].Subsets[j].Name).To(Equal(sub.name))
				Expect(meshConfig.DesinationRules[i].Subsets[j].Labels).To(Equal(sub.labels))
			}

			Expect(meshConfig.VirtualServices).To(HaveLen(4))
			for i, expected := range []virtualService{
				{"details.default.svc.cluster.local",
					[]route{{[]match{{
						map[string]string{"app": "reviews"}, 9080,
					}}}}},
				{"productpage.default.svc.cluster.local",
					[]route{{[]match{{
						map[string]string{"app": "reviews"}, 9080,
					}}}}},
				{"ratings.default.svc.cluster.local",
					[]route{{[]match{{
						map[string]string{"app": "reviews"}, 9080,
					}}}}},
				{"reviews.default.svc.cluster.local",
					[]route{{[]match{{
						map[string]string{"app": "reviews"}, 9080,
					}}}}},
			} {
				Expect(meshConfig.VirtualServices[i].Metadata.Name).To(Equal(expected.host))
				Expect(meshConfig.VirtualServices[i].Hosts).To(Equal([]string{expected.host}))
				Expect(meshConfig.VirtualServices[i].Gateways).To(Equal([]string{"mesh"}))
				Expect(meshConfig.VirtualServices[i].Http).To(HaveLen(len(inputRoutingRules)))
				for j, rr := range inputRoutingRules {
					if reqMatchers := len(rr.RequestMatchers); reqMatchers > 0 {
						Expect(meshConfig.VirtualServices[i].Http[j].Match).To(HaveLen(reqMatchers))
					} else {
						Expect(meshConfig.VirtualServices[i].Http[j].Match[0].Port).To(Equal(expected.routes[j].matches[0].port))
						Expect(meshConfig.VirtualServices[i].Http[j].Match[0].SourceLabels).To(Equal(expected.routes[j].matches[0].sourceLabels))
					}
				}
			}
		}

		Expect(resourceErrs).NotTo(BeNil())
	})
})