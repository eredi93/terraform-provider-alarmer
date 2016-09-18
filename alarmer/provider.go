package alarmer

import (
  "log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"datadog_api_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
                DefaultFunc: schema.EnvDefaultFunc("DATADOG_API_KEY", nil),
			},

			"datadog_app_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
                DefaultFunc: schema.EnvDefaultFunc("DATADOG_APP_KEY", nil),
			},

			"aws_access_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
                DefaultFunc: schema.EnvDefaultFunc("AWS_ACCESS_KEY", nil),
			},

			"aws_secret_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
                DefaultFunc: schema.EnvDefaultFunc("AWS_SECRET_KEY", nil),
			},

			"aws_region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
                DefaultFunc: schema.EnvDefaultFunc("AWS_REGION", nil),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"alarmer_datadog_monitor":   resourceAlarmerDatadogMonitor(),
			"alarmer_datadog_timeboard": resourceAlarmerDatadogTimeboard(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
    config := Config{
		DatadogApiKey: d.Get("datadog_api_key").(string),
		DatadogAppKey: d.Get("datadog_app_key").(string),
		AwsAccessKey: d.Get("aws_access_key").(string),
		AwsSecretKey: d.Get("aws_secret_key").(string),
		AwsRegion: d.Get("aws_region").(string),
	}

    log.Println("[INFO] Initializing Alarmer client")
    return config.Client()
}
