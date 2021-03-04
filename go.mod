module github.com/falcosecurity/falcosidekick

go 1.16

require (
	cloud.google.com/go/pubsub v1.9.1
	cloud.google.com/go/storage v1.14.0 // indirect
	github.com/Azure/azure-event-hubs-go/v3 v3.3.4
	github.com/DataDog/datadog-go v4.2.0+incompatible
	github.com/PagerDuty/go-pagerduty v1.3.0
	github.com/aws/aws-sdk-go v1.36.23
	github.com/cloudevents/sdk-go/v2 v2.3.1
	github.com/emersion/go-sasl v0.0.0-20200509203442-7bfe0ed36a21
	github.com/emersion/go-smtp v0.14.0
	github.com/google/uuid v1.1.2
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/nats-io/nats-streaming-server v0.19.0 // indirect
	github.com/nats-io/nats.go v1.10.0
	github.com/nats-io/stan.go v0.8.1
	github.com/prometheus/client_golang v1.9.0
	github.com/segmentio/kafka-go v0.4.8
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	golang.org/x/oauth2 v0.0.0-20210218202405-ba52d332ba99
	google.golang.org/api v0.40.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/client-go v0.20.1
)
