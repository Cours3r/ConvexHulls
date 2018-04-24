# @author Benjamin Stewart
# generate an array of 100 random x,y points
from random import randint
import numpy.random as npr

num = 75
filename = "points.txt"

def genPoint():
	x = randint(-100,100)
	y = randint(-100,100)
	return (x,y)

def genPointArr(size):
	pointArr = []
	while len(pointArr) < size:
		pointArr.append(genPoint())
	return pointArr

def getNormalArr(mu,sigma,size):
	a = npr.normal(mu,sigman(2,size))
	out = []
	for coord in a:
		out.append((coord[0],coord[1]))
	return out
	
with open(filename,"w") as f:
	for p in pointArr:
		f.write("{0},{1}\n".format(p[0],p[1]))



