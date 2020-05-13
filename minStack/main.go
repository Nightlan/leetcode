package main

type minStackData struct {
	val    int
	curMin int
}

type MinStack struct {
	buf []*minStackData
	top int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		top: -1,
	}
}

func (this *MinStack) Push(x int) {
	if this.top == -1 {
		this.min = x
	} else {
		if this.min > x {
			this.min = x
		}
	}
	data := &minStackData{
		val:    x,
		curMin: this.min,
	}
	this.top++
	if this.top < len(this.buf) {
		this.buf[this.top] = data
	} else {
		this.buf = append(this.buf, data)
	}
}

func (this *MinStack) Pop() {
	this.buf[this.top] = nil
	this.top--
	if this.top >= 0 {
		data := this.buf[this.top]
		this.min = data.curMin
	}
}

func (this *MinStack) Top() int {
	data := this.buf[this.top]
	return data.val
}

func (this *MinStack) GetMin() int {
	return this.min
}

func main() {

}
