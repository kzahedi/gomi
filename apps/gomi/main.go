package main

import (
	"flag"
	"os"

	"github.com/kzahedi/gomi"
)

func main() {

	helpPtr := flag.Bool("h", false, "help")
	verbosePtr := flag.Bool("v", false, "verbose")
	logPtr := flag.Bool("log", false, "log coverted data")
	cfgPtr := flag.String("cfg", "", "Config file. If present, other command line parameters will be ignored.")
	measurePtr := flag.String("mi", "MI_W", "available quantifications are: MI_W, MI_A, MI_A_Prime, MI_MI, MI_SY, MI_SY_NID, MI_CA, MI_WA, MI_WS, MI_Wp, CA, UI, CI, MI_IN")
	continuousPtr := flag.Bool("c", false, "Use continuous measure.")
	continuousModePtr := flag.Int("cm", 1, "Only required if KSG Estimator is involved. 1 = First KSG MI Estimator, 2 = Second KSG MI Estimator.")
	stateDependentPtr := flag.Bool("s", false, "Use state-dependent measure.")
	binsPtr := flag.Int("bins", 0, "Optional. Only used for discrete measures. Input is single value that is used for all random variables.")
	iterationsPtr := flag.Int("i", 0, "Optional. Iterations, e.g. for Iterative Scaling used for MI_SY.")
	outputPtr := flag.String("o", "out.txt", "Output file.")
	wBinsPtr := flag.String("wbins", "", "Only used for discrete measures. Input is single value that is used for all random variables that make up W. Input can also be a list of values. In this case there must a value for each variable in W.")
	aBinsPtr := flag.String("abins", "", "Only used for discrete measures. Input is single value that is used for all random variables that make up A. Input can also be a list of values. In this case there must a value for each variable in A.")
	sBinsPtr := flag.String("sbins", "", "Only used for discrete measures. Input is single value that is used for all random variables that make up S. Input can also be a list of values. In this case there must a value for each variable in S.")
	filePtr := flag.String("file", "", "File that contains full data set.")
	wIndicesPtr := flag.String("wi", "", "Indices of the W columns in the file given by -file.")
	aIndicesPtr := flag.String("ai", "", "Indices of the A columns in the file given by -file.")
	sIndicesPtr := flag.String("si", "", "Indices of the S columns in the file given by -file.")
	wFilePtr := flag.String("wfile", "", "File that contains W data set.")
	aFilePtr := flag.String("afile", "", "File that contains A data set.")
	sFilePtr := flag.String("sfile", "", "File that contains S data set.")
	dFilePtr := flag.String("dfile", "", "File (yaml) that contains all min, max values for W, S, A (optional)")
	sparsePtr := flag.Bool("sparse", false, "Use Sparse Matrix Implementation")
	knnPtr := flag.Int("k", 30, "k used for KSG and FP estimators")
	flag.Parse()

	if *helpPtr == true {
		flag.PrintDefaults()
		os.Exit(0)
	}

	p := gomi.CreateParametersContainer()

	if *cfgPtr != "" {
		p.SetConfigFile(*cfgPtr)
	}

	p.SetMeasureName(*measurePtr)
	p.SetUseContinuous(*continuousPtr)
	p.SetContinuousMode(*continuousModePtr)
	p.SetUseStateDependent(*stateDependentPtr)
	p.SetGlobalBins(*binsPtr)
	p.SetWBins(*wBinsPtr)
	p.SetSBins(*sBinsPtr)
	p.SetABins(*aBinsPtr)
	p.SetK(*knnPtr)
	p.SetOutput(*outputPtr)
	p.SetVerbose(*verbosePtr)
	p.SetGlobalFile(*filePtr)
	p.SetWIndices(*wIndicesPtr)
	p.SetSIndices(*sIndicesPtr)
	p.SetAIndices(*aIndicesPtr)
	p.SetWFile(*wFilePtr)
	p.SetSFile(*sFilePtr)
	p.SetAFile(*aFilePtr)
	p.SetDFile(*dFilePtr)
	p.SetIterations(*iterationsPtr)
	p.SetLogData(*logPtr)
	p.SetUseSparseMatrix(*sparsePtr)

	p.CheckParameters()

	var data gomi.Data

	data.Read(p)

	if p.UseContinuous == true && p.UseStateDependent == true {
		gomi.ContinuousSDCalculations(p, data)
	}
	if p.UseContinuous == true && p.UseStateDependent == false {
		gomi.ContinuousAvgCalculations(p, data)
	}

	if p.UseContinuous == false && p.UseSparseMatrix == true && p.UseStateDependent == false {
		gomi.DiscreteAvgCalculationsSparse(p, data)
	}
	if p.UseContinuous == false && p.UseSparseMatrix == true && p.UseStateDependent == true {
		gomi.DiscreteSDCalculationsSparse(p, data)
	}

	if p.UseContinuous == false && p.UseSparseMatrix == false && p.UseStateDependent == false {
		gomi.DiscreteAvgCalculations(p, data)
	}
	if p.UseContinuous == false && p.UseSparseMatrix == false && p.UseStateDependent == true {
		gomi.DiscreteSDCalculations(p, data)
	}
}
