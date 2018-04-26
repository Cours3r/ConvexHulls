# @author Benjamin Stewart

import math, time, csv, sys
import matplotlib.pyplot as plt
from random import randint
import numpy.random as npr


def getVec(a,b):
	# Returns 1D array (vector) from a -> b
	x_comp = b[0] - a[0]
	y_comp = b[1] - a[1]
	return (x_comp,y_comp)

def mag(v):
	return math.sqrt(v[0]**2 + v[1]**2)

def cross(v1,v2):
	return v1[0]*v2[1] - v1[1]*v2[0]

def getLowestY(points):
	## Get start point, lowest y and corresponding x
	start = points[0]
	ind = 0
	for i in range(1,len(points)):
		if points[i][1] < start[1]:
			start = points[i]
			ind = i
		elif points[i][1] == start[1]:
			if points[i][0] < start[0]:
				start = points[i]
				ind = i
	points.pop(ind)
	return (points, start)

def isleft(a,b,origin):
	# test if a is to the left of b
	A = getVec(origin,a)
	B = getVec(origin,b)
	crossPod = cross(A,B)
	if crossPod > 0:
		return False
	elif crossPod < 0:
		return True
	elif crossPod == 0:
		if mag(A) > mag(B):
			return False
		else:
			return True

def sortByAngle(pList,origin):
	# origin is a point
	# points is array of points
	out = [pList.pop(0)]
	for dot in pList:
		Left = False
		spot = -1
		while Left == False:
			spot += 1
			if spot > (len(out)-1):
				Left = True
			else:
				Left = isleft(dot,out[spot],origin)
		out[spot:spot] = [dot]
	return out[::-1]

def ccw(p,c,n):
	v1 = getVec(p,c)
	v2 = getVec(p,n)
	if cross(v1,v2) >= 0:
		return True
	else:
		return False

def hull(points):
	points,start = getLowestY(points)
	## Sort by polar angle
	angleSorted = sortByAngle(points,start)
	storage = angleSorted.copy()
	## Calculate the path
	path = [start,angleSorted.pop(0),angleSorted.pop(0)]
	if not ccw(path[0],path[1],path[2]):
		path.pop(1)
	while angleSorted:
		path.append(angleSorted.pop(0))
		while not ccw(path[-3],path[-2],path[-1]):
			path.pop(-2)
	return path,start,storage

def genPointArr(low,high,size):
	pointArr = []
	while len(pointArr) < size:
		pointArr.append(genPoint(low,high))
	return pointArr

def getNormalArr(mu,sigma,size):
	a = npr.normal(mu,sigma,(size,2))
	out = []
	for coord in a:
		out.append((coord[0],coord[1]))
	return out

def drawGraph(storage,outSidePath,origin):
	chart = plt.figure()
	plt.axhline(0,color="black",linewidth=.5)
	plt.axvline(0,color="black",linewidth=.5)
	pathInd = ["start"]
	for node in range(0,len(storage)):
		if storage[node] in outSidePath:
			pathInd.append(node)
		else:
			plt.plot(storage[node][0],storage[node][1],"o")
		plt.annotate(str(node),xy=(storage[node][0],storage[node][1]+.25))
		plt.plot(origin[0],origin[1],"^")
	plt.annotate("start",xy=(origin[0],origin[1]+.25))
	for i in range(0,len(outSidePath)):
		if i == len(outSidePath)-1:
			j = 0
		else:
			j = i+1
		plt.plot([outSidePath[i][0],outSidePath[j][0]],[outSidePath[i][1],outSidePath[j][1]],marker="^")
	return pathInd

def main():
	m = 0
	sd = 8
	s = 125
	#arr = getNormalArr(m,sd,s)
	if len(sys.argv) > 1:
		pointFile = sys.argv[1]
	else:
		pointFile = "points.txt"
	arr = []
	with open(pointFile,"r") as f:
		reader = csv.reader(f)
		for row in reader:
			arr.append((float(row[0]),float(row[1])))

	tic = time.time()
	outSidePath,origin,storage = hull(arr)
	## Calculate time taken
	elapsed = time.time() - tic
	## Draw the diagram
	pathInd = drawGraph(storage,outSidePath,origin)
	## Report Results
	print("PathInd: ",pathInd)
	print("Length: ",len(pathInd))
	print("Time: ",elapsed)
	plt.show()

if __name__ == "__main__":
	main()
