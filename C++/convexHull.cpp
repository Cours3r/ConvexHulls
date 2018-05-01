/*
 * C++ Convex Hull Algorithm
 * @author Benjamin Stewart
 */

#include <vector>
#include <cmath>
#include <iostream>
using namespace std;

struct point {
	float x;
	float y;
};

struct vec {
	float i;
	float j;
};

float mag(vec v) {
	return sqrt(pow(v.i,2) + pow(v.j,2));
}

float cross(vec v1, vec v2) {
	return (v1.i*v2.j - v1.j*v2.i);
}

vec getVec(point a, point b) {
	float icomp = b.x - a.x;
	float jcomp = b.y - a.y;
	vec v = {icomp,jcomp};
	return v;
}

bool isLeft(point a, point b, point origin) {
	vec v1 = getVec(origin,a);
	vec v2 = getVec(origin,b);
	float crossProd = cross(v1,v2);
	if (crossProd > 0) {
		return false;
	} else if (crossProd < 0) {
		return true;
	} else {
		if (mag(v1) > mag(v2)) {
			return false;
		} else {
			return true;
		}
	}
}



int main() {
	point A = {1.0,2.0};
	point B = {4.0,9.0};
	point org = {0.0,0.0};
	cout << "A: <" << A.x << ", " << A.y << ">" << endl;
	cout << "B: <" << B.x << ", " << B.y << ">" << endl;
	vec v1 = getVec(A,B);
	cout << "v1: " << v1.i << ", " << v1.j << endl;
	cout << "A is left of B: ";
	if ((isLeft(A,B,org))) {
		cout << "True" << endl;
	} else {
		cout << "False" << endl;
	}
	return 0;
}
