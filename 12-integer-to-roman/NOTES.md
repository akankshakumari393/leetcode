# convert slice of string to string using strings.Join
​
# improvement -- can be done using while loop
​
​
```
class Solution {
public String intToRoman(int num) {
String roman[] = {"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"};
int integer[] =  {1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1};
String result="";
int i=0;
while(num>0 && i<integer.length) {
if(num>=integer[i]) {
int n= num/integer[i];
num= num%integer[i];
while(n-->0) result+=roman[i];
}
i++;
}
return result;
}
}
```