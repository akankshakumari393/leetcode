type MyCalendar struct {
    root *Event
}


type Event struct {
    start, end int
    left, right *Event
}


func Constructor() MyCalendar {
    return MyCalendar{
        root: nil,
    }
}


func (this *MyCalendar) Book(start int, end int) bool {
    if end<start || start <0 || end<0 {
        return false
    }
    //do tree insert
    return Helper(start, end, &this.root) 
}

func Helper(start int, end int , root **Event) bool {
    if *root == nil {
        *root = &Event{
            start : start,
            end : end,
            left : nil,
            right : nil,
        }
        return true
    }
    if (*root).start <= start && (*root).end >=end {
        return false
    }
    
    if (*root).start >= end {
        return Helper (start, end, &(*root).right)
    }
    if (*root).end <= start {
        return Helper (start, end, &(*root).left)
    }
    return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */