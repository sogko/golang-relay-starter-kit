package data

// Model structs
type User struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Widgets []*Widget `json:"widgets"`
}

type Widget struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// Mock data
var viewer = &User{
	Id:   "1",
	Name: "Anonymous",
}
var widgets = []*Widget{
	&Widget{"0", "What's-it"},
	&Widget{"1", "Who's-it"},
	&Widget{"2", "How's-it"},
}

// Data accessors
func GetUser(id string) *User {
	if id == viewer.Id {
		return viewer
	}
	return nil
}
func GetViewer() *User {
	return viewer
}
func GetWidget(id string) *Widget {
	for _, widget := range widgets {
		if widget.Id == id {
			return widget
		}
	}
	return nil
}
func GetWidgets() []*Widget {
	return widgets
}
func WidgetsToInterfaceSlice(widgets ...*Widget) []interface{} {
	var interfaceSlice []interface{} = make([]interface{}, len(widgets))
	for i, d := range widgets {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}
