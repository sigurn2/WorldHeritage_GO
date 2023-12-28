package data

type Attribute struct {
	Name   string
	Values []string
}

var AttributeList = []Attribute{
	architectureStyle,
}
var architectureStyle = Attribute{
	Name: "architecture_style",
	Values: []string{
		"Ancient and Prehistoric Architecture",
		"Classical and European Architectural Styles (includes Renaissance, Baroque, Gothic, Romanesque, Byzantine, Roman)",
		"Islamic and Middle Eastern Architecture",
		"Colonial and Industrial Revolution Architecture",
		"Asian Architectural Styles",
		"South American and Mesoamerican Architecture",
		"Polynesian and Pacific Island Architecture",
		"African and Native American Architecture",
		"Norse and Icelandic Architecture",
		"Art and Cave Paintings",
	},
}
