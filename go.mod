module github.com/owncloud/ocis-markdown-editor

go 1.13

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.1
	contrib.go.opencensus.io/exporter/ocagent v0.6.0
	contrib.go.opencensus.io/exporter/zipkin v0.1.1
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/micro/cli/v2 v2.1.2
	github.com/oklog/run v1.0.0
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/owncloud/ocis-pkg/v2 v2.3.0
	github.com/spf13/viper v1.6.1
	go.opencensus.io v0.22.4
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2
)

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
