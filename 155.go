package main

type MinStack struct {
	stack       []int
	s           int
	tracerStack []Tracer
	n           int
}

type Tracer struct {
	val   int
	count int
}

func ConstructorMinStack() MinStack {
	return MinStack{
		stack: make([]int, 0), s: -1,
		tracerStack: make([]Tracer, 0), n: -1,
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	this.s++
	if this.n < 0 || val < this.tracerStack[this.n].val {
		this.tracerStack = append(this.tracerStack, Tracer{val: val, count: 1})
		this.n++
	} else if val == this.tracerStack[this.n].val {
		this.tracerStack[this.n].count++
	}
	// fmt.Println(this)
}

func (this *MinStack) Pop() {
	top := this.stack[this.s]
	this.stack = this.stack[:this.s]
	this.s--
	tracerTop := &this.tracerStack[this.n]
	if top == tracerTop.val {
		if tracerTop.count < 2 {
			this.tracerStack = this.tracerStack[:this.n]
			this.n--
		} else {
			tracerTop.count--
		}
	}
	// fmt.Println("pop", this)
}

func (this *MinStack) Top() int {
	return this.stack[this.s]
}

func (this *MinStack) GetMin() int {
	return this.tracerStack[this.n].val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
