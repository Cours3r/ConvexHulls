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
	X float64
	Y float64
}

type Vec struct {
	X float64
	Y float64
}

func getVec(a,b Point) Vec {
	// returns vector a -> b
	iComp := b.X - a.X
	jComp := b.Y - a.Y
	vec := Vec{iComp,jComp}
	return vec
}

func main() {
	a := Point{1,2}
	b := Point{5,6}
	fmt.Println(a)
	fmt.Println(b)
	v := getVec(a,b)
	fmt.Println(v)
}
