**gomi** is a software package for the evaluation of morphological intelligence measures on data. gomi includes all currently available morphological intelligence measures. This includes measures with continuous estimators as well as measures that use a discrete estimator.

The following sections will only provide minimal information about the measures. The goal of this post to provide information about the application of the measures, not an introduction of the measures themselves. For this purpose, please visit the other posts on these pages:

**gomi** is written in [Go](https://golang.org). For the installation of Go, please read the
installation documentation provided [here](https://golang.org/doc/install). Pre-compiled binaries of gomi for
Windows, Linux, and macOS are available in the release files [here](https://github.com/kzahedi/gomi/releases).

Once Go is installed, gomi can easily be installed using the following commands:


```shell
go get github.com/kzahedi/gomi
```
The following two packages are required and might have to be installed manually:


```shell
go get github.com/kzahedi/goent
go get github.com/kzahedi/utils
```
A zip (and tarball) of stable releases can be downloaded [here].

## Using the gomi binary

To calculate MI_W with the binary, use the following command line parameters:


```shell
gomi -mi MI_W -file musfib.csv -wi 1,2,3 -ai 9 -v -bins 300 -o MI_W.csv
```
The file musfib was used in [1] and can be downloaded [here](http://github.com/kzahedi/entropy).

Explanation of the command line options used in the example above:

| Option | Explanation |
|---|---|
| -mi MI_W	| Chooses MI_W as the measure|
|-file musfib.csv|	Data is provided in the file musfib.csv|
|-wi 1,2,3	| Columns 1,2,3 of the data provided in musfib.csv define the world state (counting starts with 0)|
| -ai 9	| Column 9 of the data provided in musfib.csv defines the action state (counting starts with 0)|
| -v |	gomi will print useful information while it is running and it will print the result|
|-bins 300  | Global definition of the binning. This value will be used for each of the four columns |
| -o MI_W.csv|  The result and the specified parameters will be written to MI_W.csv|

## Using gomi as a library

Using gomi as a library
The measures implemented in gomi can also be used as a library. The following code snippet gives an example:

```go
package main

import (
	"fmt"
 
	"github.com/kzahedi/goent/dh"
	goent "github.com/kzahedi/goent/discrete"
	mc "github.com/kzahedi/gomi/discrete"
)
 
func main() {
	// W and A are just examples for data. These would usually be read from some data file
	// For this example to work, we provide some dummy data
	w := [][]float64{{0.0, 1.0},
		{0.1, 1.1},
		{0.2, 1.2},
		{0.3, 1.3},
		{0.4, 1.4}}
 
	a := [][]float64{{0.0, 1.0, 2.0},
		{0.1, 1.1, 2.1},
		{0.2, 1.2, 2.2},
		{0.3, 1.3, 2.3},
		{0.4, 1.4, 2.4}}
 
	// discretising data. Discrestise(data, bins for each column, min values for each column, max values for each column)
	wDiscretised := dh.Discretise(w, []int{10, 10}, []float64{0.0, 0.0}, []float64{1.0, 2.0})
	aDiscretised := dh.Discretise(a, []int{10, 10, 10}, []float64{0.0, 0.0, 0.0}, []float64{1.0, 2.0, 3.0})
 
	// univariate variables
	wUnivariate := dh.MakeUnivariateRelabelled(wDiscretised, []int{10, 10})
	aUnivariate := dh.MakeUnivariateRelabelled(aDiscretised, []int{10, 10, 10})
 
	// creating w', w, a data
	w2w1a1 := make([][]int, len(w)-1, len(w)-1)
 
	for i := 0; i < len(w)-1; i++ {
		w2w1a1[i] = make([]int, 3, 3)
		w2w1a1[i][0] = wUnivariate[i+1]
		w2w1a1[i][1] = wUnivariate[i]
		w2w1a1[i][2] = aUnivariate[i]
	}
 
	// calculating p(w',w,a)
	pw2w1a1 := goent.Emperical3D(w2w1a1)
 
	// calculating MI_W
	result := mc.MorphologicalComputationW(pw2w1a1)
 
	fmt.Println(result)
}
```


A complete reference can be found at
[here](http://keyan.ghazi-zahedi.eu/gomi).


References:
-  K. Ghazi-Zahedi, C. Langer, and N. Ay. Morphological computation: Synergy of body and brain. Entropy, 19(9), 2017.
- K. Ghazi-Zahedi, R. Deimel, G. Montufar, V. Wall, and O. Brock. Morphological computation: The good, the bad, and the ugly. In IROS 2017, 2017.
-  K. Ghazi-Zahedi, D. F. Haeufle, G. F. Montufar, S. Schmitt, and N. Ay. Evaluating morphological computation in muscle and dc-motor driven models of hopping movements. Frontiers in Robotics and AI, 3(42), 2016.
- K. Ghazi-Zahedi and J. Rauh. Quantifying morphological computation based on an information decomposition of the sensorimotor loop. In Proceedings of the 13th European Conference on Artificial Life (ECAL 2015), pages 70—77, July 2015.
- K. Zahedi and N. Ay. Quantifying morphological computation. Entropy, 15(5):1887–1915, 2013.

