module github.com/keycloak/keycloak-operator

require (
	cloud.google.com/go v0.46.3 // indirect
	contrib.go.opencensus.io/exporter/ocagent v0.4.12 // indirect
	github.com/Azure/go-autorest v11.7.0+incompatible // indirect
	github.com/coreos/prometheus-operator v0.34.0
	github.com/fatih/color v1.9.0 // indirect
	github.com/go-openapi/spec v0.19.6
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golangci/gocyclo v0.0.0-20180528144436-0a533e8fa43d // indirect
	github.com/golangci/golangci-lint v1.23.3 // indirect
	github.com/golangci/revgrep v0.0.0-20180812185044-276a5c0a1039 // indirect
	github.com/google/uuid v1.1.1
	github.com/gostaticanalysis/analysisutil v0.0.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.11.2 // indirect
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/integr8ly/grafana-operator/v3 v3.0.2-0.20200103111057-03d7fa884db4
	github.com/jirfag/go-printf-func-name v0.0.0-20200119135958-7558a9eaa5af // indirect
	github.com/json-iterator/go v1.1.9
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/openshift/api v3.9.1-0.20190924102528-32369d4db2ad+incompatible
	github.com/operator-framework/operator-sdk v0.15.1
	github.com/pelletier/go-toml v1.6.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/securego/gosec v0.0.0-20200203094520-d13bb6d2420c // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.6.2 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/tommy-muehle/go-mnd v1.2.0 // indirect
	go.opencensus.io v0.22.1 // indirect
	golang.org/x/sys v0.0.0-20200202164722-d101bd2416d5 // indirect
	golang.org/x/tools v0.0.0-20200207131002-533eb2654509 // indirect
	google.golang.org/api v0.10.0 // indirect
	google.golang.org/grpc v1.24.0 // indirect
	gopkg.in/ini.v1 v1.52.0 // indirect
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20190918143330-0270cf2f1c1d
	mvdan.cc/unparam v0.0.0-20191111180625-960b1ec0f2c2 // indirect
	sigs.k8s.io/controller-runtime v0.4.0
	sourcegraph.com/sqs/pbtypes v1.0.0 // indirect
)

// Pinned to kubernetes-1.15.4
replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.2.0+incompatible
	k8s.io/api => k8s.io/api v0.0.0-20191016110408-35e52d86657a
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191016113550-5357c4baaf65
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191004115801-a2eda9f80ab8
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20191016112112-5190913f932d
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20191016114015-74ad18325ed5
	k8s.io/client-go => k8s.io/client-go v0.0.0-20191016111102-bec269661e48
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20191016115326-20453efc2458
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20191016115129-c07a134afb42
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20191004115455-8e001e5d1894
	k8s.io/component-base => k8s.io/component-base v0.0.0-20191016111319-039242c015a9
	k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190828162817-608eb1dad4ac
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20191016115521-756ffa5af0bd
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20191016112429-9587704a8ad4
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20191016114939-2b2b218dc1df
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20191016114407-2e83b6f20229
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20191016114748-65049c67a58b
	k8s.io/kubectl => k8s.io/kubectl v0.0.0-20191016120415-2ed914427d51
	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20191016114556-7841ed97f1b2
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20191016115753-cf0698c3a16b
	k8s.io/metrics => k8s.io/metrics v0.0.0-20191016113814-3b1a734dba6e
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20191016112829-06bb3c9d77c9
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.3.0
)

replace github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309 // Required by Helm

replace github.com/openshift/api => github.com/openshift/api v0.0.0-20190924102528-32369d4db2ad // Required until https://github.com/operator-framework/operator-lifecycle-manager/pull/1241 is resolved

go 1.13
