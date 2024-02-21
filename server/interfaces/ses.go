package interfaces

import (
	"github.com/aws/aws-sdk-go/service/ses"
)

type SESClient interface {
	SendEmail(input *ses.SendEmailInput) (*ses.SendEmailOutput, error)
}
