type TextEditor struct {
    text *list.List
    cursor *list.Element
}


func Constructor() TextEditor {
    l := list.New()
    l.PushBack('*')
    return TextEditor {
        text: l,
        cursor: l.Front(),
    }
}


func (this *TextEditor) AddText(text string)  {
    for _, c := range text {
        this.text.InsertAfter(c, this.cursor)
        this.cursor = this.cursor.Next()
    }
}


func (this *TextEditor) DeleteText(k int) int {
    count := 0
    for this.cursor != this.text.Front() && count < k {
        count++
        b := this.cursor.Prev()
        this.text.Remove(this.cursor)
        this.cursor = b
    }
    return count
}


func (this *TextEditor) CursorLeft(k int) string {
    count := 0
    for this.cursor != this.text.Front() && count < k {
        this.cursor = this.cursor.Prev()
        count++
    }
    return this.GetPrev(10)
}


func (this *TextEditor) CursorRight(k int) string {
    count := 0
    for this.cursor != this.text.Back() && count < k {
        this.cursor = this.cursor.Next()
        count++
    }
    return this.GetPrev(10)
}

func (this *TextEditor) GetPrev(k int) string {
    count, itr := 0, this.cursor
    for itr != this.text.Front() && count < k {
        itr = itr.Prev()
        count++
    }
    itr = itr.Next()
    
    sb := strings.Builder{}
    for ; itr != this.cursor.Next(); itr = itr.Next() {
        sb.WriteRune(itr.Value.(rune))
    }
    return sb.String()
}


/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */