package main

// 结构体

type Hero struct {
	Name  string
	Ad    int
	Leval int
}

func (this *Hero) GetName() string {
	return this.Name
}
func (this *Hero) SetName(newName string) {
	this.Name = newName

}
func main() {

}
