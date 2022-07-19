type FreqStack struct {
    // freeMap store the frequency of element
    freqMap map[int]int
    // mapOfStack store the map of frequnecy to the stack of element counting on that frequency
    mapOfStack map[int][]int
    // currentMaxFreq
    maxFreq int
}


func Constructor() FreqStack {
    return FreqStack{
        freqMap : make(map[int]int),
        mapOfStack : make(map[int][]int),
    }
}
/*
push 5 5 5 5 5 3 3 3 2 2 2 1
group = 5:[5] 4: [5] 3: [5,3,2] 2:[5,3,2] 1 : [5,3,21] maxfreq=5
1st pop 5:[] 4: [5] 3: [5,3,2] 2:[5,3,2] 1 : [5,3,21] maxfreq=4
2nd pop 5:[] 4: [] 3: [5,3,2] 2:[5,3,2] 1 : [5,3,21] maxfreq=3
3rd pop 5:[] 4: [] 3: [5,3] 2:[5,3,2] 1 : [5,3,21] maxfreq=3
4th pop 5:[] 4: [] 3: [5] 2:[5,3,2] 1 : [5,3,21] maxfreq=3
5th pop 5:[] 4: [] 3: [] 2:[5,3,2] 1 : [5,3,21] maxfreq=2 and so on...
*/


func (this *FreqStack) Push(val int)  {
    // increase the frequency of value
    if _ , ok := this.freqMap[val]; ok {
        this.freqMap[val]++
    } else {
        this.freqMap[val] = 1
    }
    f := this.freqMap[val]
    // if the frequency is > maxFrequency update the value of maxFreq
    if f > this.maxFreq {
        this.maxFreq = f
    }
    // append of mapOfstack at that given freqency and add the val
    this.mapOfStack[f] = append(this.mapOfStack[f], val)
}


func (this *FreqStack) Pop() int {
    // get the stack of maxFrequency
    currentstack := this.mapOfStack[this.maxFreq]
    // get the last inserted value from the stack
    val := currentstack[len(currentstack)-1]
    // update the stack and remove last element
    currentstack = currentstack[:len(currentstack)-1]
    // update the frequency count of the value
    this.freqMap[val]--
    // update the mapOfStack of maxfreq with the removed stack element
    this.mapOfStack[this.maxFreq] = currentstack
    // if the stack with maxfreq is 0 decrease the maxfreq
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