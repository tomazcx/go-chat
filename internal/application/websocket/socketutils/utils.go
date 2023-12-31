package socketutils

import (
	"bytes"
	"context"

	"github.com/tomazcx/go-chat-app/internal/application/httpapi/templates/components"
)

func BufferizeMessage(message, author string) []byte {
	buffer := &bytes.Buffer{}
	components.Message(author, message).Render(context.Background(), buffer)
	return buffer.Bytes()
}
