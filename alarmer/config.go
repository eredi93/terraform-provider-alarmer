package alarmer

import "log"

type Config struct {
    DatadogApiKey string
    DatadogAppKey string
    AwsAccessKey string
    AwsSecretKey string
    AwsRegion string
}

func (c *Config) Client() (*AlarmerClient, error) {
    client := NewAlarmerClient(c.DatadogApiKey, c.DatadogApiKey,
        c.AwsAccessKey, c.AwsSecretKey, c.AwsRegion)
    log.Printf("[INFO] Alarmer Client configured")

    return client, nil
}
