type FreqStack struct {
    freqMap map[int]int
    mapOfStack map[int][]int
    maxFreq int
}


func Constructor() FreqStack {
    return FreqStack{
        freqMap : make(map[int]int),
        mapOfStack : make(map[int][]int),
        maxFreq : 0,
    }
}


func (this *FreqStack) Push(val int)  {
    this.freqMap[val] = this.freqMap[val] + 1
    if this.freqMap[val] > this.maxFreq {
        this.maxFreq = this.freqMap[val]
    }
    if _, ok := this.mapOfStack[this.maxFreq]; !ok{
        var stack []int
        this.mapOfStack[this.maxFreq] = stack
    }
    stack := this.mapOfStack[this.freqMap[val]]
    stack = append(stack, val)
    this.mapOfStack[this.freqMap[val]] = stack
}


func (this *FreqStack) Pop() int {
    currentstack := this.mapOfStack[this.maxFreq]
    val := currentstack[len(currentstack)-1]
    currentstack = currentstack[:len(currentstack)-1]
    this.freqMap[val] = this.freqMap[val] -1    
    this.mapOfStack[this.maxFreq] = currentstack
    if len(currentstack) == 0 {
        this.maxFreq = this.maxFreq-1
        return val
    }

    return val
}


/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */