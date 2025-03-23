package main

type MovingAverage struct {
	queue []int
	head  int
	tail  int
	c     int
	s     int
	sum   int
}

func MovingAverageConstructor(size int) MovingAverage {
	return MovingAverage{
		queue: make([]int, size, size),
		head:  -1,
		tail:  -1,
		c:     size,
		s:     0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.s == this.c {
		pop := this.queue[this.head]
		if this.head == this.tail {
			this.head = -1
			this.tail = -1
		} else {
			this.head = (this.head + 1) % this.c
		}
		this.sum -= pop
		this.s--
	}
	if this.head < 0 {
		this.head = 0
	}
	this.sum += val
	this.tail = (this.tail + 1) % this.c
	this.queue[this.tail] = val
	this.s++
	return float64(this.sum) / float64(this.s)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
