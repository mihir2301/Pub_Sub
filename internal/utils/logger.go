package utils

import "log"

func Log(msg string) {
	log.Printf("[PubSub]%s", msg)
}
