module github.com/americanas-go/ignite

go 1.16

// replace github.com/americanas-go/log => ../log/
// replace github.com/americanas-go/multiserver => ../multiserver/

require (
	github.com/Microsoft/go-winio v0.5.0 // indirect
	github.com/americanas-go/config v1.3.0
	github.com/americanas-go/errors v1.1.0
	github.com/americanas-go/health v1.0.0
	github.com/americanas-go/log v1.6.0
	github.com/americanas-go/multiserver v1.1.0
	github.com/americanas-go/rest-response v1.0.2
	github.com/ansrivas/fiberprometheus/v2 v2.1.2
	github.com/aws/aws-sdk-go v1.40.34
	github.com/aws/aws-sdk-go-v2 v1.9.0
	github.com/aws/aws-sdk-go-v2/config v1.8.0
	github.com/aws/aws-sdk-go-v2/credentials v1.4.0
	github.com/aws/aws-sdk-go-v2/service/kinesis v1.6.0
	github.com/aws/aws-sdk-go-v2/service/s3 v1.15.0
	github.com/aws/aws-sdk-go-v2/service/sns v1.8.0
	github.com/aws/aws-sdk-go-v2/service/sqs v1.9.0
	github.com/cheekybits/is v0.0.0-20150225183255-68e9c0620927 // indirect
	github.com/cloudevents/sdk-go/v2 v2.5.0
	github.com/coocood/freecache v1.1.1
	github.com/elastic/go-elasticsearch/v8 v8.0.0-20210823151005-3b1f3aef208c
	github.com/globocom/echo-prometheus v0.1.2
	github.com/go-chi/chi/v5 v5.0.4
	github.com/go-chi/cors v1.2.0
	github.com/go-playground/validator/v10 v10.9.0
	github.com/go-redis/redis/v7 v7.4.1
	github.com/go-redis/redis/v8 v8.11.3
	github.com/go-resty/resty/v2 v2.6.0
	github.com/gocql/gocql v0.0.0-20210817081954-bc256bbb90de
	github.com/gofiber/fiber/v2 v2.18.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.3.0
	github.com/graphql-go/graphql v0.8.0
	github.com/graphql-go/handler v0.2.3
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/hashicorp/go-memdb v1.3.2
	github.com/hiko1129/echo-pprof v1.0.1
	github.com/jlaffaye/ftp v0.0.0-20210307004419-5d4190119067
	github.com/labstack/echo/v4 v4.5.0
	github.com/labstack/gommon v0.3.0
	github.com/matryer/try v0.0.0-20161228173917-9ac251b645a2 // indirect
	github.com/mattn/go-sqlite3 v1.14.8 // indirect
	github.com/nats-io/nats.go v1.12.1
	github.com/newrelic/go-agent/v3 v3.15.0
	github.com/newrelic/go-agent/v3/integrations/nrecho-v4 v1.0.1
	github.com/newrelic/go-agent/v3/integrations/nrgrpc v1.3.1
	github.com/newrelic/go-agent/v3/integrations/nrmongo v1.0.2
	github.com/newrelic/go-agent/v3/integrations/nrnats v1.1.1
	github.com/newrelic/go-agent/v3/integrations/nrredis-v7 v1.0.1
	github.com/newrelic/go-agent/v3/integrations/nrredis-v8 v1.0.0
	github.com/opentracing-contrib/echo v0.0.0-20190807091611-5fe2e1308f06
	github.com/opentracing-contrib/go-grpc v0.0.0-20210225150812-73cb765af46e
	github.com/opentracing/opentracing-go v1.2.0
	github.com/panjf2000/ants/v2 v2.4.6
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/prometheus/client_golang v1.11.0
	github.com/ravernkoh/cwlogsfmt v0.0.0-20180121032441-917bad983b4c
	github.com/segmentio/kafka-go v0.4.18
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.2.1
	github.com/stretchr/testify v1.7.0
	github.com/swaggo/echo-swagger v1.1.3
	github.com/tidwall/buntdb v1.2.6
	github.com/valyala/fasthttp v1.30.0
	github.com/wesovilabs/beyond v1.0.1
	go.mongodb.org/mongo-driver v1.7.2
	go.uber.org/fx v1.14.2
	gocloud.dev v0.24.0
	gocloud.dev/pubsub/kafkapubsub v0.24.0
	golang.org/x/net v0.0.0-20210825183410-e898025ed96a
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.40.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.33.0
	gopkg.in/matryer/try.v1 v1.0.0-20150601225556-312d2599e12e
	k8s.io/client-go v0.22.1
)
