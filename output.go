package csprcollector

type Output interface {
	Write(data []CSPRequest)
}
