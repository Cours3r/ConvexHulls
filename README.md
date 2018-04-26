# ConvexHulls

Convex hull algorithms written in different languages. This is a little project of mine that I tackle when learning new programming languages, as the complex hull challenge covers a lot of the basics for each language. I do not claim that any of this is by any means optimized, but to my knowledge the completed languages work.

## Status
- [x] Python
- [x] Golang
- [ ] C++

## Notes

The python script is not nearly as fast as the Go program, but it does plot all the points and the hull, which is very useful. The plot displayed using matplotlib, and is kinda slow. The calling the python script followed by the file name, `python hull.py points.txt`, one can import data points into the script.

The golang program is much faster than the python script, and analyzes a file. The default is points.txt. It can also take a command line arguement to choose the file; `goHull points.txt`. 

There is a python scipt included called pointGen.py which can be used to generate a file with a variable amount of points easily. `python pointGen.py 500` will generater a file of 500 random points. As a default the script buids a file with 50 points. There are plans to further build out this script to allow easy adjustment of mean, standard distrobution, and file name.

## Resources

[Wikipedia](https://en.wikipedia.org/wiki/Convex_hull_algorithms) -> Good basic reading for what the algorithm is and should do.

[Geomalgorithms](⎋http://www.geomalgorithms.com/a10-_hull-1.html) -> Good walkthrough of the math and basic principle to actually coding the algorithm
