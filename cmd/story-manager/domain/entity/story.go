package entity

type Option struct {
	Text string
	Arc  string
}

type Chapter struct {
	Arc     string
	Title   string
	Story   string
	Options []Option
}
