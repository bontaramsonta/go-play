package main

import "fmt"

func (e email) cost() int {
	cost := 0
	if e.isSubscribed {
		cost = 2
	} else {
		cost = 5
	}
	return len(e.body) * cost
}

func (e email) format() string {
	subscriptionText := ""
	if e.isSubscribed {
		subscriptionText = "Subscribed"
	} else {
		subscriptionText = "Not Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, subscriptionText)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
