module github.com/americanas-go/ignite

go 1.21

// replace github.com/americanas-go/log => ../log/
// replace github.com/americanas-go/multiserver => ../multiserver/
// replace github.com/americanas-go/config => ../config/
// replace k8s.io/client-go => k8s.io/client-go v0.19.16
// replace github.com/nats-io/nats-server/v2 => github.com/nats-io/nats-server/v2 v2.7.4
// replace github.com/opencontainers/runc => github.com/opencontainers/runc v1.1.1
// replace github.com/opencontainers/image-spec => github.com/opencontainers/image-spec v1.0.2
// replace github.com/spf13/viper => github.com/spf13/viper v1.10.1
// replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.12.1
// replace github.com/smartystreets/assertions => github.com/smarty/assertions v1.15.1
// replace github.com/Shopify/sarama => github.com/IBM/sarama v1.42.1

replace gocloud.dev/pubsub/kafkapubsub => github.com/google/go-cloud/pubsub/kafkapubsub v0.0.0-20231101181203-6d5d2890816a

replace gocloud.dev/pubsub/mempubsub => github.com/google/go-cloud/pubsub/mempubsub v0.0.0-20231101181203-6d5d2890816a

replace gocloud.dev/pubsub/awssnssqs => github.com/google/go-cloud/pubsub/awssnssqs v0.0.0-20231101181203-6d5d2890816a

replace gocloud.dev/pubsub => github.com/google/go-cloud/pubsub v0.0.0-20231101181203-6d5d2890816a

require (
	cloud.google.com/go/bigquery v1.57.1
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/allegro/bigcache/v3 v3.1.0
	github.com/americanas-go/cache v1.0.0-beta.12
	github.com/americanas-go/config v1.8.5
	github.com/americanas-go/errors v1.1.1
	github.com/americanas-go/grapper v1.0.0-beta.15
	github.com/americanas-go/health v1.0.0
	github.com/americanas-go/log v1.8.9
	github.com/americanas-go/multiserver v1.1.1
	github.com/americanas-go/rest-response v1.0.7
	github.com/ansrivas/fiberprometheus/v2 v2.6.1
	github.com/aws/aws-sdk-go v1.48.12
	github.com/aws/aws-sdk-go-v2 v1.23.5
	github.com/aws/aws-sdk-go-v2/config v1.25.11
	github.com/aws/aws-sdk-go-v2/credentials v1.16.9
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.24.2
	github.com/aws/aws-sdk-go-v2/service/s3 v1.47.2
	github.com/aws/aws-sdk-go-v2/service/sns v1.26.2
	github.com/aws/aws-sdk-go-v2/service/sqs v1.29.2
	github.com/bytedance/sonic v1.10.2
	github.com/cloudevents/sdk-go/v2 v2.14.0
	github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be
	github.com/coocood/freecache v1.2.4
	github.com/dubonzi/otelresty v1.2.0
	github.com/elastic/go-elasticsearch/v7 v7.17.10
	github.com/elastic/go-elasticsearch/v8 v8.11.1
	github.com/globocom/echo-prometheus v0.1.2
	github.com/go-chi/chi/v5 v5.0.10
	github.com/go-chi/cors v1.2.1
	github.com/go-playground/validator/v10 v10.16.0
	github.com/go-redis/redis/v7 v7.4.1
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-resty/resty/v2 v2.10.0
	github.com/goccy/go-json v0.10.2
	github.com/gocql/gocql v1.6.0
	github.com/godror/godror v0.40.4
	github.com/gofiber/fiber/v2 v2.51.0
	github.com/google/uuid v1.4.0
	github.com/graphql-go/graphql v0.8.1
	github.com/graphql-go/handler v0.2.3
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/hashicorp/go-memdb v1.3.4
	github.com/hashicorp/vault/api v1.10.0
	github.com/hiko1129/echo-pprof v1.0.1
	github.com/jedib0t/go-pretty/v6 v6.4.9
	github.com/jlaffaye/ftp v0.2.0
	github.com/labstack/echo/v4 v4.11.3
	github.com/labstack/gommon v0.4.1
	github.com/mittwald/vaultgo v0.1.4
	github.com/nats-io/nats.go v1.31.0
	github.com/newrelic/go-agent/v3 v3.28.1
	github.com/newrelic/go-agent/v3/integrations/nrecho-v4 v1.0.4
	github.com/newrelic/go-agent/v3/integrations/nrgrpc v1.4.1
	github.com/newrelic/go-agent/v3/integrations/nrmongo v1.1.2
	github.com/newrelic/go-agent/v3/integrations/nrredis-v7 v1.0.1
	github.com/newrelic/go-agent/v3/integrations/nrredis-v8 v1.0.1
	github.com/opentracing-contrib/echo v0.0.0-20190807091611-5fe2e1308f06
	github.com/opentracing-contrib/go-grpc v0.0.0-20210225150812-73cb765af46e
	github.com/opentracing/opentracing-go v1.2.0
	github.com/panjf2000/ants/v2 v2.9.0
	github.com/prometheus/client_golang v1.17.0
	github.com/ravernkoh/cwlogsfmt v0.0.0-20180121032441-917bad983b4c
	github.com/segmentio/kafka-go v0.4.46
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.8.0
	github.com/stretchr/testify v1.8.4
	github.com/swaggo/echo-swagger v1.4.1
	github.com/tidwall/buntdb v1.3.0
	github.com/valyala/fasthttp v1.51.0
	github.com/wesovilabs/beyond v1.0.1
	go.mongodb.org/mongo-driver v1.13.0
	go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho v0.46.1
	go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo v0.46.1
	go.opentelemetry.io/otel v1.21.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.21.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.21.0
	go.opentelemetry.io/otel/sdk v1.21.0
	go.opentelemetry.io/otel/trace v1.21.0
	go.uber.org/fx v1.20.1
	gocloud.dev v0.34.0
	gocloud.dev/pubsub/kafkapubsub v0.34.0
	golang.org/x/net v0.19.0
	golang.org/x/sync v0.5.0
	google.golang.org/api v0.152.0
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.58.0
	gopkg.in/matryer/try.v1 v1.0.0-20150601225556-312d2599e12e
	k8s.io/client-go v0.28.4
	storj.io/drpc v0.0.33
)

require (
	github.com/DataDog/appsec-internal-go v1.2.0 // indirect
	github.com/DataDog/datadog-agent/pkg/util/log v0.49.0 // indirect
	github.com/DataDog/datadog-agent/pkg/util/scrubber v0.49.0 // indirect
	github.com/DataDog/go-libddwaf/v2 v2.2.1 // indirect
	github.com/IBM/sarama v1.42.1 // indirect
	github.com/apache/arrow/go/v12 v12.0.1 // indirect
	github.com/apache/thrift v0.19.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.26.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/eventbridge v1.26.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.8.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/sfn v1.24.2 // indirect
	github.com/chenzhuoyu/iasm v0.9.1 // indirect
	github.com/cihub/seelog v0.0.0-20170130134532-f561c5e57575 // indirect
	github.com/ebitengine/purego v0.5.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/go-jose/go-jose/v3 v3.0.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/flatbuffers v23.5.26+incompatible // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/klauspost/asmfmt v1.3.2 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/minio/asm2plan9s v0.0.0-20200509001527-cdd76441f9d8 // indirect
	github.com/minio/c2goasm v0.0.0-20190812172519-36a3d3bbc4f3 // indirect
	github.com/newrelic/csec-go-agent v0.5.0 // indirect
	github.com/outcaste-io/ristretto v0.2.3 // indirect
	github.com/smartystreets/goconvey v1.8.0 // indirect
	github.com/swaggo/files/v2 v2.0.0 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	go.opentelemetry.io/otel/metric v1.21.0 // indirect
	golang.org/x/exp v0.0.0-20231127185646-65229373498e // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/oauth2 v0.15.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20231127180814-3a041ad873d4 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231127180814-3a041ad873d4 // indirect
)

require (
	cloud.google.com/go v0.111.0 // indirect
	cloud.google.com/go/compute v1.23.3 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v1.1.5 // indirect
	cloud.google.com/go/pubsub v1.33.0 // indirect
	github.com/DataDog/datadog-agent/pkg/obfuscate v0.49.0 // indirect
	github.com/DataDog/datadog-agent/pkg/remoteconfig/state v0.51.0-devel // indirect
	github.com/DataDog/datadog-go/v5 v5.3.0 // indirect
	github.com/DataDog/go-tuf v1.0.2-0.5.2 // indirect
	github.com/DataDog/gostackparse v0.7.0 // indirect
	github.com/DataDog/sketches-go v1.4.3 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/andybalholm/brotli v1.0.6 // indirect
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.5.3 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.9 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.1 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.2.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.2.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.16.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.18.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.21.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.26.2 // indirect
	github.com/aws/smithy-go v1.18.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/cheekybits/is v0.0.0-20150225183255-68e9c0620927 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eapache/go-resiliency v1.4.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/elastic/elastic-transport-go/v8 v8.3.0 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.3.0 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/jsonpointer v0.20.0 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/spec v0.20.11 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/gobeam/stringy v0.0.6 // indirect
	github.com/godror/knownpb v0.1.1 // indirect
	github.com/gofiber/adaptor/v2 v2.2.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/pprof v0.0.0-20231205033806-a5a03c77bf08 // indirect
	github.com/google/wire v0.5.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.1 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.5 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.8 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.6 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/hcl v1.0.1-vault-5 // indirect
	github.com/imdario/mergo v0.3.16 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.17.4 // indirect
	github.com/klauspost/cpuid/v2 v2.2.6 // indirect
	github.com/knadh/koanf v1.5.0 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/matryer/try v0.0.0-20161228173917-9ac251b645a2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nats-io/nkeys v0.4.6 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/richardartoul/molecule v1.0.1-0.20221107223329-32cfee06a052 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/rs/zerolog v1.31.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/secure-systems-lab/go-securesystemslib v0.7.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.5.1 // indirect
	github.com/swaggo/swag v1.16.2 // indirect
	github.com/tidwall/btree v1.7.0 // indirect
	github.com/tidwall/gjson v1.17.0 // indirect
	github.com/tidwall/grect v0.1.4 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/rtred v0.1.2 // indirect
	github.com/tidwall/tinyqueue v0.1.1 // indirect
	github.com/tinylib/msgp v1.1.9 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	github.com/zeebo/errs v1.3.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.21.0
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/dig v1.17.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	go4.org/intern v0.0.0-20230525184215-6c62f75575cb // indirect
	go4.org/unsafe/assume-no-moving-gc v0.0.0-20231121144256-b99613f794b6 // indirect
	golang.org/x/arch v0.6.0 // indirect
	golang.org/x/crypto v0.16.0 // indirect; indirect221014153046-6fdb5e3db783 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/term v0.15.0 // indirect
	golang.org/x/text v0.14.0
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/tools v0.16.0 // indirect
	golang.org/x/xerrors v0.0.0-20231012003039-104605ab7028 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20231127180814-3a041ad873d4 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	inet.af/netaddr v0.0.0-20230525184311-b8eac61e914a // indirect
	k8s.io/api v0.28.4 // indirect
	k8s.io/apimachinery v0.28.4 // indirect
	k8s.io/klog/v2 v2.110.1 // indirect
	k8s.io/kube-openapi v0.0.0-20231129212854-f0671cc7e66a // indirect
	k8s.io/utils v0.0.0-20231127182322-b307cd553661 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
