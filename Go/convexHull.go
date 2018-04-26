package main

import ("fmt"
		"math"
		"os"
		"io"
		"bufio"
		"encoding/csv"
		"strconv"
		"time")

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

func Ccw(p,c,n Point) bool {
	v1 := GetVec(p,c)
	v2 := GetVec(p,n)
	if Cross(v1,v2) >= 0 {
		return true
	} else {
		return false
	}
}

func HullPath(points []Point, start Point) ([]Point,[]int) {
	path := []Point{start,points[0],points[1]}
	out := []int{0,1}
	points = points[2:]
	if !Ccw(path[0],path[1],path[2]) {
		path = path[:2]
		out = out[:1]
	}
	for i,v := range points {
		path = append(path,v)
		out = append(out,i+2)
		for !Ccw(path[len(path)-3],path[len(path)-2],path[len(path)-1]) {
			path = append(path[:len(path)-2],path[len(path)-1])
			out = append(out[:len(out)-2],out[len(out)-1])
		}
	}
	return path, out
}

func main() {
	fn := "points.txt" //Default filename
	pointArr := ReadInData(fn) //get data from file
	pointArr,base := GetLowestY(pointArr) //find starting point
	sortedPointArr := SortByAngle(pointArr,base) //sort data
	tic := time.Now()
	_, intPath := HullPath(sortedPointArr,base) // get the hull path and the number path
	toc := time.Now()
	diff := toc.Sub(tic)
	//fmt.Println(complexHull) //Print complex hull path points
	fmt.Println("Path: start +",intPath) //print number of the points on the path
	fmt.Println("Length: ",len(intPath)+1)
	fmt.Println("Time (s): ",diff.Seconds())
	fmt.Println("Time (ns): ",diff.Nanoseconds())
}
