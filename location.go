package ngxParse

//Location represents a location in nginx config
type Location struct {
	*Directive
	Modifier string
	Match    string
}
