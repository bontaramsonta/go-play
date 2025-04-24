package main

// this switch on type is facinating
func getExpenseReportSwitch(e expense) (string, float64) {
	switch obj := e.(type) {
	case email:
		return obj.toAddress, obj.cost()
	case sms:
		return obj.toPhoneNumber, obj.cost()
	default:
		return "", 0.0
	}
}

func getExpenseReport(e expense) (string, float64) {
	emailObj, ok := e.(email)
	if ok {
		return emailObj.toAddress, emailObj.cost()
	}
	smsObj, ok := e.(sms)
	if ok {
		return smsObj.toPhoneNumber, smsObj.cost()
	}
	return "", 0.0
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
