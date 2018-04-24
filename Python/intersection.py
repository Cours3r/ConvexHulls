# @author Benjamin Stewart
# Python code to check if two lines intersect
#from collections import namedtuple
import numpy as np

def getVec(a,b):
	# Returns 1D array (vector) from a -> b
	x_comp = b[0] - a[0]
	y_comp = b[1] - a[1]
	return np.array([x_comp,y_comp])

def intersection(p1a,p1b,p2a,p2b):
	r = getVec(p1a,p1b)
	s = getVec(p2a,p2b)
	diff = np.array([p2a[0]-p1a[0],p2a[1]-p1a[1]])
	#print(r)
	#print(s)
	#print(diff)
	if np.cross(r,s) == 0:
		if np.cross(diff,r) == 0:
			#collinear
			if abs(diff) <= r:
				#close enough to intersect
				if np.dot(diff,r) > 0:
					# Point in the same direction
					return 10
				else:
					# Point in opposite direction
					return 1
			else:
				# Not close enough to intersect
				return 2
		else:
			# Paralell
			return 3
	else:
		t = np.cross(diff,s) / np.cross(r,s)
		u = np.cross(diff,r) / np.cross(r,s)
		#print(t)
		#print(u)
		if (t <= 1 and t >= 0) and (u <= 1 and u >= 0):
			# Intersects at p + t*r
			return 11
		else:
			# Not paralell and not intersecting
			return 4
		

a = (2,2)
b = (7,5)
c = (4,7)
d = (-2,3)
e = (-3,7)
f = (-8,3)
g = (-2,-1)
h = (-2,-4)
i = (-6,-3)
j = (-10,-3)

segments = [(a,c),(b,d),(c,f),(e,f),(e,d),(h,f),(i,h),(g,i),(j,h)]
names = {1:"A",2:"B",3:"C",4:"D",5:"E",6:"F",7:"G",8:"H",9:"I"}

for x in range(0,len(segments)):
	tmp = []
	for y in range(0,len(segments)):
		if x == y:
			continue
		if intersection(segments[x][0],segments[x][1],segments[y][0],segments[y][1]) > 9:
			tmp.append(names[y+1])
	print(names[x+1] + " intersects with ",tmp)



