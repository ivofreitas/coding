package stack

type MinStack struct {
	memory  []int
	minimum []int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	if len(this.minimum) == 0 || this.minimum[len(this.minimum)-1] >= val {
		this.minimum = append(this.minimum, val)
	}
	this.memory = append(this.memory, val)
}

func (this *MinStack) Pop() {
	if len(this.memory) > 0 {
		if len(this.minimum) > 0 &&
			this.memory[len(this.memory)-1] == this.minimum[len(this.minimum)-1] {
			this.minimum = this.minimum[:len(this.minimum)-1]
		}
		this.memory = this.memory[:len(this.memory)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.memory) > 0 {
		return this.memory[len(this.memory)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.minimum) > 0 {
		return this.minimum[len(this.minimum)-1]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
