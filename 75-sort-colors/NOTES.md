​
​
Dutch national flag algorithm
​
Rule
[0 .. low-1] -> 0
​
[high + 1 ..... n ] -> 2
​
Initially
high == n-1
mid == 0
low == 0
​
​
Now mid can be 0, 1, 2
​
if mid == 0 swap (low, mid)
if mid == 1 mid ++
if mid == 2 swap( high, mid), high --
​
​