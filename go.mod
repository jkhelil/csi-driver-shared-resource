module github.com/openshift/csi-driver-shared-resource

go 1.14

require (
	github.com/container-storage-interface/spec v1.3.0
	github.com/go-openapi/spec v0.19.7 // indirect
	github.com/kubernetes-csi/csi-lib-utils v0.7.0
	github.com/openshift/api v0.0.0-20211007134530-4cb30f221b89
	github.com/openshift/client-go v0.0.0-20211007143529-7ab6242249ff
	github.com/prometheus/client_golang v1.11.1
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.28.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	golang.org/x/net v0.0.0-20211209124913-491a49abca63
	google.golang.org/grpc v1.40.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/api v0.23.5
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/klog/v2 v2.30.0
	k8s.io/kubectl v0.21.2
	k8s.io/utils v0.0.0-20211116205334-6203023598ed
	sigs.k8s.io/controller-runtime v0.11.2
)

replace k8s.io/client-go => k8s.io/client-go v0.22.1
