# @author Benjamin Stewart
# generate an array of 100 random x,y points
from random import randint
import numpy.random as npr

num = 75
filename = "points.txt"

def genPointArr(size):
	pointArr = []
	while len(pointArr) < size:
		pointArr.append(genPoint())
	return pointArr

def getNormalArr(mu,sigma,size):
	a = npr.normal(mu,sigma,(size,2))
	out = []
	for coord in a:
		out.append((coord[0],coord[1]))
	return out

def main():
	m = 0
	sd = 8
	s = 5
	arr = getNormalArr(m,sd,s)
	with open(filename,"w") as f:
		for p in arr:
			f.write("{0},{1}\n".format(p[0],p[1]))

if __name__ == "__main__":
	main()
