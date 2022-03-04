package basic

type Person struct {
	Name string `json:"username"`
	Age  uint   `json:"age"`
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	date   string
	id     uint
}
