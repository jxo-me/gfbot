package telebot

import (
	"fmt"
	"strconv"
)

func StateKey(ctx IContext, strategy KeyStrategy) string {
	switch strategy {
	case KeyStrategySender:
		return strconv.FormatInt(ctx.Sender().ID, 10)
	case KeyStrategyChat:
		return strconv.FormatInt(ctx.Chat().ID, 10)
	case KeyStrategySenderAndChat:
		fallthrough
	default:
		// Default to KeyStrategySenderAndChat if unknown strategy
		return fmt.Sprintf("%d/%d", ctx.Sender().ID, ctx.Chat().ID)
	}
}
