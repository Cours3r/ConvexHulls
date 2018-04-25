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

1 x GetVec(a,b) // return a -> b
2 x Mag(v) // return magnitude
3 x Cross(v1,v2) // return v1[0]*v2[1] - v1[1]*v2[0]
4 x IsLeft(a,b,origin)
5 x ReadInData(filename)
6 - sortByAngle(pList,origin)
7 - hull(pList)
8 - ccw(a,b,c)
 */

type Point struct {
	X float64
	Y float64
}

type Vec struct {
	i float64
	j float64
}

func ReadInData(filename string) []Point {
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

func GetVec(a,b Point) Vec {
	// returns vector a -> b
	iComp := b.X - a.X
	jComp := b.Y - a.Y
	vec := Vec{iComp,jComp}
	return vec
}

func Mag(v Vec) float64 {
	return math.Sqrt(math.Pow(v.i,2)+math.Pow(v.j,2))
}

func Cross(v1,v2 Vec) float64 {
	return v1.i*v2.j - v1.j*v2.i
}

func IsLeft(a,b,origin Point) bool {
	v1 := GetVec(origin,a)
	v2 := GetVec(origin,b)
	crossProd := Cross(v1,v2)
	if crossProd > 0 {
		return false
	} else if crossProd < 0 {
		return true
	} else {
		if Mag(v1) > Mag(v2) {
			return false
		} else {
			return true
		}
	}
}

func GetLowestY(a []Point) ([]Point,Point) {
	start := a[0]
	ind := 0
	for i,p := range a[1:] {
		if p.Y < start.Y {
			start = p
			ind = i+1
		} else if p.Y == start.Y {
			if p.X < start.X {
				start = p
				ind = i+1
			}
		}
	}
	a = append(a[:ind], a[ind+1:]...)
	return a,start
}

func SortByAngle(points []Point, origin Point) []Point {
	var sortedPoints []Point
	sortedPoints = append(sortedPoints,points[0])
	points = points[1:]
	tmp := Point{0,0}
	for _,v := range points {
		left := false
		i := -1
		for !left {
			i++
			if i > len(sortedPoints)-1 {
				left = true
			} else {
				left = IsLeft(v,sortedPoints[i],origin)
			}
		}
		sortedPoints = append(sortedPoints,tmp)
		copy(sortedPoints[i+1:],sortedPoints[i:])
		sortedPoints[i] = v
	}
	for i := len(sortedPoints)/2-1; i >= 0; i-- {
		opp := len(sortedPoints)-1-i
		sortedPoints[i], sortedPoints[opp] = sortedPoints[opp], sortedPoints[i]
	}
	return sortedPoints
}

/*
		left := false
		spot := -1
		for left != true {
			spot++
			if spot > len(sortedPoints)-1 {
				*left = true
			} else {
				*left = IsLeft(v,sortedPoints[spot],origin)
			}
		sortedPoints = append(sortedPoints,tmp)
		copy(sortedPoints[spot+1:],sortedPoints[spot:])
		sortedPoints[spot] = v
		}
 */

func main() {
	a := Point{1,2}
	b := Point{5,6}
	c := Point{-5,4}
	//fmt.Println("Point A: ",a)
	//fmt.Println("Point B: ",b)
	v := GetVec(a,b)
	//fmt.Println("Vector V: <",v.i,v.j,">")
	//fmt.Println("Magnitude of V: ",Mag(v))
	w := GetVec(a,c)
	//fmt.Println("Vector W: <",w.i,w.j,">")
	fmt.Println("VxW: ",Cross(v,w))
	//o := Point{0,0}
	//fmt.Println("A is left of B around origin: ",IsLeft(a,b,o))
	fn := "points.txt"
	pointArr := ReadInData(fn)
	for i,v := range pointArr {
		fmt.Println("pointArr [",i,"] = ",v.X,v.Y)
	}
	pointArr,base := GetLowestY(pointArr)
	sortedPointArr := SortByAngle(pointArr,base)
	fmt.Println("Start = ",base.X,base.Y)
	for j,w := range sortedPointArr {
		fmt.Println("pointArr [",j,"] = ",w.X,w.Y)
	}
}
