# Icecream Parlor

## Problem Statement

Sunny and Johnny like to pool their money and go to the ice cream parlor. Johnny never buys the same flavor that Sunny does. The only other rule they have is that they spend all of their money.

Given a list of prices for the flavors of ice cream, select the two that will cost all of the money they have.

Complete the function icecreamParlor below to return an array containing the indexes of the prices of the two flavors they buy. The returned array must be sorted ascending.

### Input Format

The first line contains an integer, `t` , denoting the number of trips to the ice cream parlor. The next `t` sets of lines each describe a visit. Each trip is described as follows:


1. The integer `m` , the amount of money they have pooled.
2. The integer `n`, the number of flavors offered at the time.
3. `n` space-separated integers denoting the cost of each flavor: `cost[c1,c2,c3...cn]`.

**Note**: The index within the cost array represents the flavor of the ice cream purchased.

### Output Format

For each test case, print two space-separated integers denoting the indexes of the two flavors purchased, in ascending order. 

**Sample Input**

```
2
4
5
1 4 5 3 2
4
4
2 2 4 3
```


**Sample Output**

```
1 4
1 2
```

[Link to problem](https://www.hackerrank.com/challenges/icecream-parlor/problem)

### To Run

```
$ cat input.stdin | go run main.go
```
