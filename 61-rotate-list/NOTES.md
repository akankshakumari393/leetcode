```
take a node current and make it point to head
​
calculate the length of linked list
​
Now current is pointing to the last element
​
make a circular linkedlist by cur.Next = head
​
iterate with current for len-k number of time
​
set the head to cur.next and
set cur.next to nil, to break the list in the mid
```