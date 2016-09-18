package alarmer

import (
	"github.com/zorkian/go-datadog-api"
)


type AlarmerClient struct {
    Datadog *datadog.Client
}

func NewAlarmerClient(DatadogApiKey string, DatadogAppKey string) (*AlarmerClient, error) {
    client := AlarmerClient{
        Datadog: datadog.NewClient(DatadogApiKey, DatadogAppKey),
    }

    return client, nil
}
