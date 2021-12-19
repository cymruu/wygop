package responses

import (
	"encoding/json"
	"fmt"
)

type APIResponse struct {
	Data  json.RawMessage `json:"data"`
	Error *WykopError     `json:"error"`
}

type WykopError struct {
	Code      uint16 `json:"code"`
	Field     string `json:"field"`
	MessageEn string `json:"message_en"`
	MessagePl string `json:"message_pl"`
}

func (e *WykopError) Error() string {
	return fmt.Sprintf("%s [%d]", e.MessageEn, e.Code)
}
