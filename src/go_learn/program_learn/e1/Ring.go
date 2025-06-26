package main

// 循环链表
type Ring struct {
	next, prev *Ring
	Value      interface{}
}

func (r *Ring) init() *Ring {
	r.next = r
	r.prev = r
	return r
}

func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	p := r
	for i := 1; i < n; i++ {
		p.next = &Ring{prev: p}
		p = p.next
	}
	p.next = r
	r.prev = p
	return r
}
func main() {

	r := new(Ring)
	r.init()

}
