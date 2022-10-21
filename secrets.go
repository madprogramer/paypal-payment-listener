package paypalwebhook
import "github.com/plutov/paypal/v4"
import "os"

var(
	//KEYS FOR REST API ACCESSES!
	//By default, these are assigned to environment variables,
	//`secrets.go` is provided so that you may modify this code
	//if you are storing your keys elsewhere
	clientID = os.Getenv("PAYPAL_CLIENT_ID")
	secretID = os.Getenv("PAYPAL_SECRET")
	webhookID = os.Getenv("PAYPAL_WEBHOOK_ID")
	//apiMode = paypal.APIBaseLive
	apiMode = paypal.APIBaseSandBox
)
