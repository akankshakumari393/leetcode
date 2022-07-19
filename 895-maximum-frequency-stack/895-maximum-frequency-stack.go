type FreqStack struct {
    freqMap map[int]int
    mapOfStack map[int][]int
    maxFreq int
}


func Constructor() FreqStack {
    return FreqStack{
        freqMap : make(map[int]int),
        mapOfStack : make(map[int][]int),
    }
}


func (this *FreqStack) Push(val int)  {
    if _ , ok := this.freqMap[val]; ok {
        this.freqMap[val]++
    } else {
        this.freqMap[val] = 1
    }
    f := this.freqMap[val]
    if f > this.maxFreq {
        this.maxFreq = f
    }
    if _, ok := this.mapOfStack[this.maxFreq]; !ok{
        var stack []int
        this.mapOfStack[this.maxFreq] = stack
    }
    this.mapOfStack[f] = append(this.mapOfStack[f], val)
}


func (this *FreqStack) Pop() int {
    currentstack := this.mapOfStack[this.maxFreq]
    val := currentstack[len(currentstack)-1]
    currentstack = currentstack[:len(currentstack)-1]
    this.freqMap[val]--    
    this.mapOfStack[this.maxFreq] = currentstack
    if len(currentstack) == 0 {
        this.maxFreq--
    }
    return val
}


/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */