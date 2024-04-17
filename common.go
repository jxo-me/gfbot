package telebot

import (
	"fmt"
	"strconv"
)

func StateKey(ctx Context, strategy KeyStrategy) string {
	switch strategy {
	case KeyStrategySender:
		return strconv.FormatInt(ctx.Sender().ID, 10)
	case KeyStrategyChat:
		return strconv.FormatInt(ctx.Chat().ID, 10)
	case KeyStrategySenderAndChat:
		fallthrough
	default:
		// Default to KeyStrategySenderAndChat if unknown strategy todo
		uid := ""
		cid := ""
		if ctx.Sender() != nil {
			uid = fmt.Sprintf("%d", ctx.Sender().ID)
		}
		if ctx.Chat() != nil {
			cid = fmt.Sprintf("%d", ctx.Chat().ID)
		}
		if uid != "" && cid != "" {
			return fmt.Sprintf("%d/%d", ctx.Sender().ID, ctx.Chat().ID)
		}
		if uid != "" {
			return uid
		}
		if cid != "" {
			return cid
		}
		return ""
	}
}
