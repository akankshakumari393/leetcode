type MyHashSet struct {
    Set [10000][]int
    Size int
}


func Constructor() MyHashSet {
    var hashset MyHashSet
    hashset.Size = 10000
    return hashset
}


func (this *MyHashSet) Add(key int)  {
    hash := key % this.Size
    this.Set[hash] = append(this.Set[hash], key)
}


func (this *MyHashSet) Remove(key int)  {
    hash := key % this.Size
    for idx, ele := range this.Set[hash] {
        if ele == key {
            this.Set[hash][idx] = -1
        }
    }
}

func (this *MyHashSet) Contains(key int) bool {
    hash := key % this.Size
    for _, ele := range this.Set[hash] {
        if ele == key {
            return true
        }
    }
    return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */