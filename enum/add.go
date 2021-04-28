package enum

import (
	log "github.com/sirupsen/logrus"
)

func Add(a, b interface{}) interface{} {

	switch typeAB := a.(type) {
	default:
		log.WithFields(log.Fields{
			"type": typeAB,
		}).Fatal("Unexpected type!")
		return nil
	case int:
		return a.(int) + b.(int)
	case float64:
		return a.(float64) + b.(float64)
	}
}
