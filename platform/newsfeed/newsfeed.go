package newsfeed

// Getter items
type Getter interface {
	GetAll() []Item
}

// Adder items
type Adder interface {
	Add(item Item)
}

// Item json struct
type Item struct {
	Title string `json: "title"`
	Post  string `json: "post"`
}

// Repo news
type Repo struct {
	Items []Item
}

// New instance
func New() *Repo {
	return &Repo{}
}

// Add (item) to Items
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

// GetAll Items
func (r *Repo) GetAll() []Item {
	return r.Items
}
