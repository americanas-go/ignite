module github.com/americanas-go/ignite

go 1.19

// replace github.com/americanas-go/log => ../log/
// replace github.com/americanas-go/multiserver => ../multiserver/
// replace github.com/americanas-go/config => ../config/
// replace k8s.io/client-go => k8s.io/client-go v0.19.16
// replace github.com/nats-io/nats-server/v2 => github.com/nats-io/nats-server/v2 v2.7.4
// replace github.com/opencontainers/runc => github.com/opencontainers/runc v1.1.1
// replace github.com/opencontainers/image-spec => github.com/opencontainers/image-spec v1.0.2
// replace github.com/spf13/viper => github.com/spf13/viper v1.10.1
// replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.12.1

require (
	cloud.google.com/go/bigquery v1.41.0
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/allegro/bigcache/v3 v3.1.0
	github.com/americanas-go/cache v1.0.0-beta.9
	github.com/americanas-go/config v1.8.0
	github.com/americanas-go/errors v1.1.0
	github.com/americanas-go/grapper v1.0.0-beta.5
	github.com/americanas-go/health v1.0.0
	github.com/americanas-go/log v1.8.6
	github.com/americanas-go/multiserver v1.1.0
	github.com/americanas-go/rest-response v1.0.7
	github.com/ansrivas/fiberprometheus/v2 v2.4.1
	github.com/aws/aws-sdk-go v1.44.101
	github.com/aws/aws-sdk-go-v2 v1.16.15
	github.com/aws/aws-sdk-go-v2/config v1.17.6
	github.com/aws/aws-sdk-go-v2/credentials v1.12.19
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.15.18
	github.com/aws/aws-sdk-go-v2/service/s3 v1.27.10
	github.com/aws/aws-sdk-go-v2/service/sns v1.18.0
	github.com/aws/aws-sdk-go-v2/service/sqs v1.19.9
	github.com/bytedance/sonic v1.5.0-rc
	github.com/cloudevents/sdk-go/v2 v2.11.0
	github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be
	github.com/coocood/freecache v1.2.3
	github.com/elastic/go-elasticsearch/v8 v8.4.0
	github.com/globocom/echo-prometheus v0.1.2
	github.com/go-chi/chi/v5 v5.0.7
	github.com/go-chi/cors v1.2.1
	github.com/go-playground/validator/v10 v10.11.1
	github.com/go-redis/redis/v7 v7.4.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-resty/resty/v2 v2.7.0
	github.com/goccy/go-json v0.10.0
	github.com/gocql/gocql v1.2.1
	github.com/godror/godror v0.34.0
	github.com/gofiber/fiber/v2 v2.37.1
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/graphql-go/graphql v0.8.0
	github.com/graphql-go/handler v0.2.3
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/hashicorp/go-memdb v1.3.3
	github.com/hashicorp/vault/api v1.7.2
	github.com/hiko1129/echo-pprof v1.0.1
	github.com/jedib0t/go-pretty/v6 v6.3.8
	github.com/jlaffaye/ftp v0.1.0
	github.com/labstack/echo/v4 v4.9.0
	github.com/labstack/gommon v0.3.1
	github.com/mittwald/vaultgo v0.1.1
	github.com/nats-io/nats.go v1.17.0
	github.com/newrelic/go-agent/v3 v3.19.1
	github.com/newrelic/go-agent/v3/integrations/nrecho-v4 v1.0.2
	github.com/newrelic/go-agent/v3/integrations/nrgrpc v1.3.1
	github.com/newrelic/go-agent/v3/integrations/nrmongo v1.0.2
	github.com/newrelic/go-agent/v3/integrations/nrnats v1.1.1
	github.com/newrelic/go-agent/v3/integrations/nrredis-v7 v1.0.1
	github.com/newrelic/go-agent/v3/integrations/nrredis-v8 v1.0.0
	github.com/opentracing-contrib/echo v0.0.0-20190807091611-5fe2e1308f06
	github.com/opentracing-contrib/go-grpc v0.0.0-20210225150812-73cb765af46e
	github.com/opentracing/opentracing-go v1.2.0
	github.com/panjf2000/ants/v2 v2.5.0
	github.com/prometheus/client_golang v1.13.0
	github.com/ravernkoh/cwlogsfmt v0.0.0-20180121032441-917bad983b4c
	github.com/segmentio/kafka-go v0.4.35
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.5.0
	github.com/stretchr/testify v1.8.1
	github.com/swaggo/echo-swagger v1.3.4
	github.com/tidwall/buntdb v1.2.10
	github.com/valyala/fasthttp v1.40.0
	github.com/wesovilabs/beyond v1.0.1
	go.mongodb.org/mongo-driver v1.11.0
	go.uber.org/fx v1.18.1
	gocloud.dev v0.26.0
	gocloud.dev/pubsub/kafkapubsub v0.26.0
	golang.org/x/net v0.2.0
	golang.org/x/sync v0.0.0-20220907140024-f12130a52804
	google.golang.org/api v0.96.0
	google.golang.org/grpc v1.49.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/DataDog/dd-trace-go.v1 v1.45.1
	gopkg.in/matryer/try.v1 v1.0.0-20150601225556-312d2599e12e
	k8s.io/client-go v0.25.1
	storj.io/drpc v0.0.32
)

require (
	cloud.google.com/go v0.102.1 // indirect
	cloud.google.com/go/compute v1.7.0 // indirect
	cloud.google.com/go/iam v0.3.0 // indirect
	cloud.google.com/go/pubsub v1.19.0 // indirect
	github.com/DataDog/datadog-agent/pkg/obfuscate v0.0.0-20211129110424-6491aa3bf583 // indirect
	github.com/DataDog/datadog-agent/pkg/remoteconfig/state v0.42.0-rc.1 // indirect
	github.com/DataDog/datadog-go v4.8.2+incompatible // indirect
	github.com/DataDog/datadog-go/v5 v5.0.2 // indirect
	github.com/DataDog/go-tuf v0.3.0--fix-localmeta-fork // indirect
	github.com/DataDog/gostackparse v0.5.0 // indirect
	github.com/DataDog/sketches-go v1.2.1 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/Shopify/sarama v1.32.0 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/armon/go-metrics v0.3.9 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.4.8 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.16 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.22 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.16 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.23 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.0.13 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.1.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.13.16 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.22 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.13.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.16.18 // indirect
	github.com/aws/smithy-go v1.13.3 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v3 v3.0.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cheekybits/is v0.0.0-20150225183255-68e9c0620927 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20211019084208-fb5309c8db06 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/elastic/elastic-transport-go/v8 v8.1.0 // indirect
	github.com/emicklei/go-restful/v3 v3.8.0 // indirect
	github.com/fatih/color v1.9.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/go-logfmt/logfmt v0.5.1 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.19.6 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/gobeam/stringy v0.0.5 // indirect
	github.com/godror/knownpb v0.1.0 // indirect
	github.com/gofiber/adaptor/v2 v2.1.25 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/gnostic v0.5.7-v3refs // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/pprof v0.0.0-20210720184732-4bb14d4b1be1 // indirect
	github.com/google/wire v0.5.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.1.0 // indirect
	github.com/googleapis/gax-go/v2 v2.5.1 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v0.16.2 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-plugin v1.4.3 // indirect
	github.com/hashicorp/go-retryablehttp v0.6.6 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-secure-stdlib/mlock v0.1.1 // indirect
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.6 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/go-version v1.2.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/vault/sdk v0.5.1 // indirect
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.0.0 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.2 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/knadh/koanf v1.4.2 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matryer/try v0.0.0-20161228173917-9ac251b645a2 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.0.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nats-io/nats-server/v2 v2.9.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/oklog/run v1.0.0 // indirect
	github.com/onsi/ginkgo/v2 v2.5.0 // indirect
	github.com/onsi/gomega v1.24.1 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/richardartoul/molecule v1.0.1-0.20221107223329-32cfee06a052 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rs/zerolog v1.28.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/secure-systems-lab/go-securesystemslib v0.4.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/swaggo/files v0.0.0-20220728132757-551d4a08d97a // indirect
	github.com/swaggo/swag v1.8.1 // indirect
	github.com/tidwall/btree v1.4.2 // indirect
	github.com/tidwall/gjson v1.14.3 // indirect
	github.com/tidwall/grect v0.1.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/rtred v0.1.2 // indirect
	github.com/tidwall/tinyqueue v0.1.1 // indirect
	github.com/tinylib/msgp v1.1.6 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	github.com/zeebo/errs v1.2.2 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/dig v1.15.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.23.0 // indirect
	go4.org/intern v0.0.0-20211027215823-ae77deb06f29 // indirect
	go4.org/unsafe/assume-no-moving-gc v0.0.0-20220617031537-928513b29760 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/crypto v0.0.0-20220829220503-c86fa9a7ed90 // indirect
	golang.org/x/oauth2 v0.0.0-20220909003341-f21342109be1 // indirect
	golang.org/x/sys v0.2.0 // indirect
	golang.org/x/term v0.2.0 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/time v0.0.0-20220722155302-e5dcc9cfc0b9 // indirect
	golang.org/x/tools v0.2.0 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20220913154956-18f8339a66a5 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/square/go-jose.v2 v2.5.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	inet.af/netaddr v0.0.0-20220617031823-097006376321 // indirect
	k8s.io/api v0.25.1 // indirect
	k8s.io/apimachinery v0.25.1 // indirect
	k8s.io/klog/v2 v2.70.1 // indirect
	k8s.io/kube-openapi v0.0.0-20220803162953-67bda5d908f1 // indirect
	k8s.io/utils v0.0.0-20220728103510-ee6ede2d64ed // indirect
	sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.3 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
