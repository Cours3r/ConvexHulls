package main

import ("fmt"
		"math"
		"os"
		"io"
		"bufio"
		"encoding/csv"
		"strconv")

/*
Functions:

1 x getVec(a,b) // return a -> b
2 x mag(v) // return magnitude
3 x cross(v1,v2) // return v1[0]*v2[1] - v1[1]*v2[0]
4 x isLeft(a,b,origin)
5 - sortByAngle(pList,origin)
6 - hull(pList)
 */

func readInData(filename string) []Point {
	csvFile, _ := os.Open(filename)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var points []Point
	for {
		line,err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
		x,_ := strconv.ParseFloat(line[0],64)
		y,_ := strconv.ParseFloat(line[1],64)
		points = append(points, Point{x,y})
	}
	return points
}


type Point struct {
	X float64
	Y float64
}

type Vec struct {
	i float64
	j float64
}

func getVec(a,b Point) Vec {
	// returns vector a -> b
	iComp := b.X - a.X
	jComp := b.Y - a.Y
	vec := Vec{iComp,jComp}
	return vec
}

func mag(v Vec) float64 {
	return math.Sqrt(math.Pow(v.i,2)+math.Pow(v.j,2))
}

func cross(v1,v2 Vec) float64 {
	return v1.i*v2.j - v1.j*v2.i
}

func isLeft(a,b,origin Point) bool {
	v1 := getVec(origin,a)
	v2 := getVec(origin,b)
	crossProd := cross(v1,v2)
	if crossProd > 0 {
		return false
	} else if crossProd < 0 {
		return true
	} else {
		if mag(v1) > mag(v2) {
			return false
		} else {
			return true
		}
	}
}



func main() {
	a := Point{1,2}
	b := Point{5,6}
	c := Point{-5,4}
	fmt.Println("Point A: ",a)
	fmt.Println("Point B: ",b)
	v := getVec(a,b)
	fmt.Println("Vector V: <",v.i,v.j,">")
	fmt.Println("Magnitude of V: ",mag(v))
	w := getVec(a,c)
	fmt.Println("Vector W: <",w.i,w.j,">")
	fmt.Println("VxW: ",cross(v,w))
	o := Point{0,0}
	fmt.Println("A is left of B around origin: ",isLeft(a,b,o))
	fn := "points.txt"
	pointArr := readInData(fn)
	for i,v := range pointArr {
		fmt.Println("pointArr [",i,"] = ",v.X,v.Y)
	}
}
