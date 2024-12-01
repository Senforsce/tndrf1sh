package cliquablelist

type Item struct {
	Name string
	URL  string
}

type Triple = struct {
	Subject   string `turtle:"subject"`
	Predicate string `turtle:"predicate"`
	Object    string `turtle:"object"`
}
