package aws

import (
	"time"

	"github.com/spf13/pflag"
	"sigs.k8s.io/aws-load-balancer-controller/pkg/aws/throttle"
)

const (
	flagAWSRegion        = "aws-region"
	flagAWSAPIEndpoints  = "aws-api-endpoints"
	flagAWSAPIThrottle   = "aws-api-throttle"
	flagAWSVpcID         = "aws-vpc-id"
	flagAWSVpcCacheTTL   = "aws-vpc-cache-ttl"
	flagAWSMaxRetries    = "aws-max-retries"
	flagLogLevel         = "log-level"
	defaultVpcID         = ""
	defaultRegion        = ""
	defaultAPIMaxRetries = 10
	defaultLogLevel      = "info"
)

type CloudConfig struct {
	// AWS Region for the kubernetes cluster
	Region string

	// Throttle settings for AWS APIs
	ThrottleConfig *throttle.ServiceOperationsThrottleConfig

	// VpcID for the LoadBalancer resources.
	VpcID string

	// VPC cache TTL in minutes
	VpcCacheTTL time.Duration

	// Max retries configuration for AWS APIs
	MaxRetries int

	// AWS endpoints configuration
	AWSEndpoints map[string]string

	// Log level for the AWS API client
	LogLevel string
}

func (cfg *CloudConfig) BindFlags(fs *pflag.FlagSet) {
	fs.StringVar(&cfg.Region, flagAWSRegion, defaultRegion, "AWS Region for the kubernetes cluster")
	fs.Var(cfg.ThrottleConfig, flagAWSAPIThrottle, "throttle settings for AWS APIs, format: serviceID1:operationRegex1=rate:burst,serviceID2:operationRegex2=rate:burst")
	fs.StringVar(&cfg.VpcID, flagAWSVpcID, defaultVpcID, "AWS VpcID for the LoadBalancer resources")
	fs.IntVar(&cfg.MaxRetries, flagAWSMaxRetries, defaultAPIMaxRetries, "Maximum retries for AWS APIs")
	fs.StringToStringVar(&cfg.AWSEndpoints, flagAWSAPIEndpoints, nil, "Custom AWS endpoint configuration, format: serviceID1=URL1,serviceID2=URL2")
	fs.StringVar(&cfg.LogLevel, flagLogLevel, defaultLogLevel, "Set the controller log level - info(default), debug")
}
