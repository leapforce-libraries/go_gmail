package gmail

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type Message struct {
	Id           *string      `json:"id,omitempty"`
	ThreadId     *string      `json:"threadId,omitempty"`
	LabelIds     *[]string    `json:"labelIds,omitempty"`
	Snippet      *string      `json:"snippet,omitempty"`
	HistoryId    *string      `json:"historyId,omitempty"`
	InternalDate *string      `json:"internalDate,omitempty"`
	Payload      *MessagePart `json:"payload,omitempty"`
	SizeEstimate *int64       `json:"sizeEstimate,omitempty"`
	Raw          *string      `json:"raw,omitempty"`
}

type MessagePart struct {
	PartId   *string          `json:"partId,omitempty"`
	MimeType *string          `json:"mimeType,omitempty"`
	Filename *string          `json:"filename,omitempty"`
	Headers  *[]Header        `json:"headers,omitempty"`
	Body     *MessagePartBody `json:"body,omitempty"`
	Parts    *[]MessagePart   `json:"parts,omitempty"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MessagePartBody struct {
	AttachmentId *string `json:"attachmentId,omitempty"`
	Size         *int64  `json:"size,omitempty"`
	Data         *string `json:"data,omitempty"`
}

type SendMessageConfig struct {
	UserId  string
	Message MimeMessage
}

func (service *Service) SendMessage(config *SendMessageConfig) (*Message, *errortools.Error) {
	if config == nil {
		return nil, errortools.ErrorMessage("SendMessageConfig cannot be nil.")
	}

	message := Message{}

	raw := config.Message.Raw(true)

	requestConfig := go_http.RequestConfig{
		Method: http.MethodPost,
		Url:    service.url(fmt.Sprintf("users/%s/messages/send", config.UserId)),
		BodyModel: Message{
			Raw: &raw,
		},
		ResponseModel: &message,
	}
	_, _, e := service.googleService.HttpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &message, nil
}
