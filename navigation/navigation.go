package navigation

// Navigation - menu navigation struct
type Navigation struct {
	ID       string       `json:"id"`
	Label    string       `json:"label"`
	Href     string       `json:"href"`
	IsRoot   bool         `json:"isRoot"`
	IsLogged bool         `json:"isLogged"`
	Children []Navigation `json:"children"`
}
