/* C++ Convex Hull Algorithm
 * @author Benjamin Stewart
 */

#include <vector>
#include <cmath>
#include <iostream>
using namespace std;

/**
 * Functions to inpliment:
 *
 * 1 - getVec(a,b)
 * 		return vector a -> b
 * 2 - isLeft(a,b,origin)
 * 		return True if is a is left of b
 * 3 - sortByAngle(pList,origin)
 * 		return sorted array?
 * 4 - ccw(p,c,n)
 * 		return if p,c,n is ccw
 * 5 - hull(pList):
 * 		return vector (or array) or hull points
 */

float mag(vector<float> v) {
	return sqrt(pow(v[0],2) + pow(v[1],2));
}

float cross(vector<float> v1, vector<float> v2) {
	return (v1[0]*v2[1] - v1[1]*v2[0]);
}

vector<float> getVec(vector<float> a, vector<float> b) {
	float x = b[0] - a[0];
	float y = b[1] - a[1];
	vector<float> v;
	v.push_back(x);
	v.push_back(y);
	return v;
}

int main() {
	vector<float> vec1;
	return 0;
}
