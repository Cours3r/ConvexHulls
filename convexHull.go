package main

import ("fmt")

/*
Functions:

1 - getVec(a,b) // return a -> b
2 - mag(v) // return magnitude
3 - cross(v1,v2) // return v1[0]*v2[1] - v1[1]*v2[0]
4 - isLeft(a,b,origin)
5 - sortByAngle(pList,origin)
6 - hull(pList
 */

/*
Point struct?
 */

type Point struct {
	X int
	Y int
}

func add(x,y int) int {
	return x + y
}

func main() {
	fmt.Println(add(12,11))
}
