package demoserver

type Options struct {
	Port                     int
	JsonLog                  bool
	SubscriptionID           string
	EnableAzureSDKCalls      bool
	HTTPPort                 int
	IdentityResourceID       string
	OperationContainerAddr   string
	ServiceBusHostName       string
	ServiceBusQueueName      string
	DatabaseServerUrl        string
	DatabasePort             int
	DatabaseName             string
	DatabaseConnectionString string
	EntityTableName          string
}
