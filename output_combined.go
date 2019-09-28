package csprcollector

type CombinedOutput struct {
	Outputs []Output
}

func (o *CombinedOutput) Write(data []CSPRequest) {
	for _, output := range o.Outputs {
		output.Write(data)
	}
}
