package main

type Output interface {
	Write(data []CSPRequest)
}
