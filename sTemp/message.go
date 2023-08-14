package sTemp

import (
	"fmt"

	"github.com/yasseldg/simplego/sLog"
)

type FlashMessage struct {
	Message string
	Type    string
}
type FlashMessages []*FlashMessage

func (fm *FlashMessages) Error(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	if len(msg) > 0 {
		sLog.Error(msg)
		*fm = append(*fm, &FlashMessage{Message: msg, Type: "error"})
	}
}

func (fm *FlashMessages) Info(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	if len(msg) > 0 {
		*fm = append(*fm, &FlashMessage{Message: msg, Type: "info"})
	}
}
