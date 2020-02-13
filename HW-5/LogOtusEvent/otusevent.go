package otusevent

import (
	"fmt"
	"io"
	"time"
)

var data = time.Now().Format("2006-01-02")

//OtusEvent - interface LogMessage()
type OtusEvent interface {
	LogMessage() string
}

//HwAccepted - структура "домашняя работа принята"
type HwAccepted struct {
	ID    int
	Grade int
}

//LogMessage - функция печатает ID и оценку
func (h *HwAccepted) LogMessage() string {
	log := fmt.Sprintf("accepted %d %d", h.ID, h.Grade)
	return log
}

//HwSubmitted - структура "студент отправил дз"
type HwSubmitted struct {
	ID      int
	Code    string
	Comment string
}

//LogMessage - функция печатает ID и комментарий к дз
func (h *HwSubmitted) LogMessage() string {
	log := fmt.Sprintf("submitted %d \"%s\"", h.ID, h.Comment)
	return log
}

//LogOtusEvent - функция реализует интерфейс OtusEvent
func LogOtusEvent(o OtusEvent, w io.Writer) {
	result := fmt.Sprintf("%s %s\n", data, o.LogMessage())
	w.Write([]byte(result))
}
