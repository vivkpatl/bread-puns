package snsController

import (
	"os"

	"github.com/aws/aws-sdk-go/service/sns"
)

type SnsController struct {
	stopCh chan os.Signal
	topic  string

	client *sns.SNS
}

func main() {
}
