module github.com/containers/podman/v3

go 1.13

require (
	github.com/BurntSushi/toml v0.4.1
	github.com/blang/semver v3.5.1+incompatible
	github.com/buger/goterm v0.0.0-20181115115552-c206103e1f37
	github.com/checkpoint-restore/checkpointctl v0.0.0-20210922093614-c31748bec9f2
	github.com/checkpoint-restore/go-criu/v5 v5.1.0
	github.com/container-orchestrated-devices/container-device-interface v0.0.0-20210325223243-f99e8b6c10b9
	github.com/containernetworking/cni v1.0.1
	github.com/containernetworking/plugins v1.0.1
	github.com/containers/buildah v1.23.1
	github.com/containers/common v0.46.1-0.20211026130826-7abfd453c86f
	github.com/containers/conmon v2.0.20+incompatible
	github.com/containers/image/v5 v5.16.1
	github.com/containers/ocicrypt v1.1.2
	github.com/containers/psgo v1.7.1
	github.com/containers/storage v1.37.1-0.20211014130921-5c5bf639ed01
	github.com/coreos/go-systemd/v22 v22.3.2
	github.com/coreos/stream-metadata-go v0.0.0-20210225230131-70edb9eb47b3
	github.com/cyphar/filepath-securejoin v0.2.3
	github.com/davecgh/go-spew v1.1.1
	github.com/digitalocean/go-qemu v0.0.0-20210209191958-152a1535e49f
	github.com/docker/distribution v2.7.1+incompatible
	github.com/docker/docker v20.10.10+incompatible
	github.com/docker/go-connections v0.4.0
	github.com/docker/go-plugins-helpers v0.0.0-20200102110956-c9a8a2d92ccc
	github.com/docker/go-units v0.4.0
	github.com/dtylman/scp v0.0.0-20181017070807-f3000a34aef4
	github.com/fsnotify/fsnotify v1.5.1
	github.com/ghodss/yaml v1.0.0
	github.com/godbus/dbus/v5 v5.0.6
	github.com/google/shlex v0.0.0-20181106134648-c34317bd91bf
	github.com/google/uuid v1.3.0
	github.com/gorilla/handlers v0.0.0-20150720190736-60c7bfde3e33
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/schema v1.2.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hpcloud/tail v1.0.0
	github.com/json-iterator/go v1.1.12
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.14
	github.com/moby/term v0.0.0-20210619224110-3f7ff695adc6
	github.com/mrunalp/fileutils v0.5.0
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.16.0
	github.com/opencontainers/go-digest v1.0.0
	github.com/opencontainers/image-spec v1.0.2-0.20210819154149-5ad6f50d6283
	github.com/opencontainers/runc v1.0.2
	github.com/opencontainers/runtime-spec v1.0.3-0.20210326190908-1c3f411f0417
	github.com/opencontainers/runtime-tools v0.9.1-0.20211020193359-09d837bf40a7
	github.com/opencontainers/selinux v1.9.1
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/rootless-containers/rootlesskit v0.14.5
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/gocapability v0.0.0-20200815063812-42c35b437635
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/vbauerster/mpb/v6 v6.0.4
	github.com/vishvananda/netlink v1.1.1-0.20210330154013-f5de75959ad5
	go.etcd.io/bbolt v1.3.6
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20210910150752-751e447fb3d0
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
	k8s.io/api v0.22.2
	k8s.io/apimachinery v0.22.2
)
