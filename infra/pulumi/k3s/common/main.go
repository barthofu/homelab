package common

type Ip struct {
	Adress  string
	Gateway string
}

type Network struct {
	Domain string
	Server string
	Ipv4   Ip
}

type MinMax struct {
	Min int
	Max int
}
