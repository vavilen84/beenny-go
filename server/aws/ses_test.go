package aws_test

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vavilen84/beenny-go/aws"
	"github.com/vavilen84/beenny-go/mocks"
	"github.com/vavilen84/beenny-go/test"
	"testing"
)

func Test_SendResetPasswordEmail(t *testing.T) {
	mockSESclient := mocks.NewSESClient(t)
	aws.SetSESClient(mockSESclient)
	mockSESclient.On("SendEmail", mock.Anything).Return(&ses.SendEmailOutput{}, nil)
	err := aws.SendResetPasswordEmail("user@example.com", "123456")
	assert.Nil(t, err)
	mockSESclient.AssertCalled(t, "SendEmail", mock.Anything)
	mockSESclient.AssertExpectations(t)
}

func Test_SendEmailVerificationMail(t *testing.T) {
	mockSESclient := mocks.NewSESClient(t)
	aws.SetSESClient(mockSESclient)
	mockSESclient.On("SendEmail", mock.Anything).Return(&ses.SendEmailOutput{}, nil)
	err := aws.SendEmailVerificationMail("user@example.com", "123456")
	assert.Nil(t, err)
	mockSESclient.AssertCalled(t, "SendEmail", mock.Anything)
	mockSESclient.AssertExpectations(t)
}

func Test_SendLoginTwoFaCode(t *testing.T) {
	mockSESclient := mocks.NewSESClient(t)
	aws.SetSESClient(mockSESclient)
	mockSESclient.On("SendEmail", mock.Anything).Return(&ses.SendEmailOutput{}, nil)
	err := aws.SendEmailVerificationMail("user@example.com", "123456")
	assert.Nil(t, err)
	mockSESclient.AssertCalled(t, "SendEmail", mock.Anything)
	mockSESclient.AssertExpectations(t)
}

func Test_notOk(t *testing.T) {
	mockSESclient := mocks.NewSESClient(t)
	aws.SetSESClient(mockSESclient)

	var sesErr awserr.Error
	sesErr = test.SesError{
		CodeData: ses.ErrCodeMessageRejected,
	}

	mockSESclient.On("SendEmail", mock.Anything).Return(&ses.SendEmailOutput{}, sesErr)
	err := aws.SendEmailVerificationMail("user@example.com", "123456")
	assert.NotNil(t, err)
	mockSESclient.AssertCalled(t, "SendEmail", mock.Anything)
	mockSESclient.AssertExpectations(t)
}
