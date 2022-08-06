type DetectSquares struct {
    Counts    map[[2]int]int
    Points         [][]int
}

func Constructor() DetectSquares {
    return DetectSquares{
        Counts: map[[2]int]int{},
        Points: [][]int{},
    }
}

func (this *DetectSquares) Add(point []int)  {
    key:=[2]int{point[0],point[1]}
    this.Counts[key]+=1
    this.Points=append(this.Points,point)
}

func (this *DetectSquares) Count(point []int) int {
    res:=0
    px,py:=point[0],point[1]
    for _,p := range this.Points{
        if math.Abs(float64(px-p[0]))!=math.Abs(float64(py-p[1])) || px==p[0] || py==p[1] {
            continue
        }
        res+=this.Counts[[2]int{px,p[1]}]*this.Counts[[2]int{p[0],py}]
    }
    return res
}


/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */