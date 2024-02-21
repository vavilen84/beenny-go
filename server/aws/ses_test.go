package aws

import (
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vavilen84/beenny-go/mocks"
	"testing"
)

func Test_SendResetPasswordEmail(t *testing.T) {
	mockSESclient := mocks.NewSESClient(t)
	SetSESClient(mockSESclient)
	mockSESclient.On("SendEmail", mock.Anything).Return(&ses.SendEmailOutput{}, nil)
	err := SendResetPasswordEmail("user@example.com", "123456")
	assert.Nil(t, err)
	mockSESclient.AssertCalled(t, "SendEmail", mock.Anything)
	mockSESclient.AssertExpectations(t)
}

func Test_SendEmailVerificationMail(t *testing.T) {
	mockSESclient := mocks.NewSESClient(t)
	SetSESClient(mockSESclient)
	mockSESclient.On("SendEmail", mock.Anything).Return(&ses.SendEmailOutput{}, nil)
	err := SendEmailVerificationMail("user@example.com", "123456")
	assert.Nil(t, err)
	mockSESclient.AssertCalled(t, "SendEmail", mock.Anything)
	mockSESclient.AssertExpectations(t)
}

func Test_SendLoginTwoFaCode(t *testing.T) {
	mockSESclient := mocks.NewSESClient(t)
	SetSESClient(mockSESclient)
	mockSESclient.On("SendEmail", mock.Anything).Return(&ses.SendEmailOutput{}, nil)
	err := SendEmailVerificationMail("user@example.com", "123456")
	assert.Nil(t, err)
	mockSESclient.AssertCalled(t, "SendEmail", mock.Anything)
	mockSESclient.AssertExpectations(t)
}