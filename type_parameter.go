package gomi

import (
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

// Parameters ...
type Parameters struct {
	MeasureName       string
	Output            string
	ConfigFile        string
	UseContinuous     bool
	UseStateDependent bool
	Verbose           bool
	LogData           bool
	UseSparseMatrix   bool
	ContinuousMode    int
	K                 int
	GlobalBins        int
	Iterations        int
	WBins             []int
	SBins             []int
	ABins             []int
	GlobalFile        string
	WIndices          []int
	SIndices          []int
	AIndices          []int
	WFile             string
	SFile             string
	AFile             string
	DFile             string
	WorldMin          []float64
	SensorMin         []float64
	ActuatorMin       []float64
	WorldMax          []float64
	SensorMax         []float64
	ActuatorMax       []float64
	NormalisationMin  []float64
	NormalisationMax  []float64
}

// GenerateString ...
func (p Parameters) GenerateString(prefix string) string {
	s := fmt.Sprintf("%sMeasure: %s", prefix, p.MeasureName)
	s = fmt.Sprintf("%s\n%sOutput:                    %s", s, prefix, p.Output)
	s = fmt.Sprintf("%s\n%sConfig file:               %s", s, prefix, p.ConfigFile)
	s = fmt.Sprintf("%s\n%sUse state-dependent:       %t", s, prefix, p.UseStateDependent)
	s = fmt.Sprintf("%s\n%sUse continuous:            %t", s, prefix, p.UseContinuous)
	s = fmt.Sprintf("%s\n%sUse sparse matrix:         %t", s, prefix, p.UseSparseMatrix)
	s = fmt.Sprintf("%s\n%sVerbose:                   %t", s, prefix, p.Verbose)
	s = fmt.Sprintf("%s\n%sLog converted data:        %t", s, prefix, p.LogData)
	s = fmt.Sprintf("%s\n%sContinuous Mode:           %d", s, prefix, p.ContinuousMode)
	s = fmt.Sprintf("%s\n%sk:                         %d", s, prefix, p.K)
	s = fmt.Sprintf("%s\n%sBins:                      %d", s, prefix, p.GlobalBins)
	s = fmt.Sprintf("%s\n%sIterations:                %d", s, prefix, p.Iterations)
	s = fmt.Sprintf("%s\n%sW bins:                    %v", s, prefix, p.WBins)
	s = fmt.Sprintf("%s\n%sS bins:                    %v", s, prefix, p.SBins)
	s = fmt.Sprintf("%s\n%sA bins:                    %v", s, prefix, p.ABins)
	s = fmt.Sprintf("%s\n%sFull data set:             %s", s, prefix, p.GlobalFile)
	s = fmt.Sprintf("%s\n%sW indices:                 %v", s, prefix, p.WIndices)
	s = fmt.Sprintf("%s\n%sS indices:                 %v", s, prefix, p.SIndices)
	s = fmt.Sprintf("%s\n%sA indices:                 %v", s, prefix, p.AIndices)
	s = fmt.Sprintf("%s\n%sW data set:                %s", s, prefix, p.WFile)
	s = fmt.Sprintf("%s\n%sS data set:                %s", s, prefix, p.SFile)
	s = fmt.Sprintf("%s\n%sA data set:                %s", s, prefix, p.AFile)
	s = fmt.Sprintf("%s\n%sW domains min:             %v", s, prefix, p.WorldMin)
	s = fmt.Sprintf("%s\n%sW domains max:             %v", s, prefix, p.WorldMax)
	s = fmt.Sprintf("%s\n%sS domains min:             %v", s, prefix, p.SensorMin)
	s = fmt.Sprintf("%s\n%sS domains max:             %v", s, prefix, p.SensorMax)
	s = fmt.Sprintf("%s\n%sA domains min:             %v", s, prefix, p.ActuatorMin)
	s = fmt.Sprintf("%s\n%sA domains max:             %v", s, prefix, p.ActuatorMax)
	s = fmt.Sprintf("%s\n%sNormalisation domains min: %v", s, prefix, p.NormalisationMin)
	s = fmt.Sprintf("%s\n%sNormalisation domains max: %v", s, prefix, p.NormalisationMax)
	return s
}

// String ...
func (p Parameters) String() string {
	return p.GenerateString("")
}

// CreateParametersContainer ...
func CreateParametersContainer() Parameters {
	return Parameters{MeasureName: defaultMeasure,
		UseContinuous:     defaultUseContinuous,
		UseStateDependent: defaultUseStateDependent,
		UseSparseMatrix:   false,
		Verbose:           false,
		LogData:           false,
		K:                 defaultK,
		GlobalBins:        defaultBins,
		Iterations:        defaultIterations,
		ContinuousMode:    defaultContinuousMode,
		WBins:             []int{},
		SBins:             []int{},
		ABins:             []int{},
		WIndices:          []int{},
		SIndices:          []int{},
		AIndices:          []int{},
		WorldMin:          []float64{},
		WorldMax:          []float64{},
		SensorMin:         []float64{},
		SensorMax:         []float64{},
		ActuatorMin:       []float64{},
		ActuatorMax:       []float64{},
		Output:            defaultOutput,
		ConfigFile:        "",
		GlobalFile:        defaultFile,
		WFile:             defaultWFile,
		SFile:             defaultSFile,
		AFile:             defaultAFile}
}

// SetMeasureName set the name of the measure to use, if
func (p *Parameters) SetMeasureName(name string) {
	if name != defaultMeasure {
		p.MeasureName = name
	}
}

// SetUseStateDependent ...
func (p *Parameters) SetUseStateDependent(b bool) {
	if b != defaultUseStateDependent {
		p.UseStateDependent = b
	}
}

// SetUseContinuous ...
func (p *Parameters) SetUseContinuous(b bool) {
	if b != defaultUseContinuous {
		p.UseContinuous = b
	}
}

// SetUseSparseMatrix ...
func (p *Parameters) SetUseSparseMatrix(b bool) {
	if b != defaultUseSparse {
		p.UseSparseMatrix = b
	}
}

// SetWBins ...
func (p *Parameters) SetWBins(wBins string) {
	if wBins != "" {
		p.WBins = parseIntString(wBins)
	}
}

// SetSBins ...
func (p *Parameters) SetSBins(sBins string) {
	if sBins != "" {
		p.SBins = parseIntString(sBins)
	}
}

// SetABins ...
func (p *Parameters) SetABins(aBins string) {
	if aBins != "" {
		p.ABins = parseIntString(aBins)
	}
}

// SetWIndices ...
func (p *Parameters) SetWIndices(wIndices string) {
	if wIndices != "" {
		p.WIndices = parseIntString(wIndices)
	}
}

// SetSIndices ...
func (p *Parameters) SetSIndices(wIndices string) {
	if wIndices != "" {
		p.SIndices = parseIntString(wIndices)
	}
}

// SetAIndices ...
func (p *Parameters) SetAIndices(wIndices string) {
	if wIndices != "" {
		p.AIndices = parseIntString(wIndices)
	}
}

// SetGlobalBins ...
func (p *Parameters) SetGlobalBins(bins int) {
	if bins != defaultBins {
		p.GlobalBins = bins
	}
}

// SetGlobalFile ...
func (p *Parameters) SetGlobalFile(file string) {
	if file != defaultFile {
		p.GlobalFile = file
	}
}

// SetWFile ...
func (p *Parameters) SetWFile(file string) {
	if file != defaultWFile {
		p.WFile = file
	}
}

// SetSFile ...
func (p *Parameters) SetSFile(file string) {
	if file != defaultSFile {
		p.SFile = file
	}
}

// SetAFile ...
func (p *Parameters) SetAFile(file string) {
	if file != defaultFile {
		p.AFile = file
	}
}

// SetLogData ...
func (p *Parameters) SetLogData(log bool) {
	p.LogData = log
}

// SetIterations ...
func (p *Parameters) SetIterations(iterations int) {
	if iterations != defaultIterations {
		p.Iterations = iterations
	}
}

type domainCfg struct {
	WorldMin    []float64 `yaml:"W min"`
	WorldMax    []float64 `yaml:"W max"`
	SensorMin   []float64 `yaml:"S min"`
	SensorMax   []float64 `yaml:"S max"`
	ActuatorMin []float64 `yaml:"A min"`
	ActuatorMax []float64 `yaml:"A max"`
}

// SetDFile ...
func (p *Parameters) SetDFile(file string) {
	p.DFile = file
	if p.DFile == "" {
		return
	}

	t := domainCfg{}

	data, err := ioutil.ReadFile(p.DFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		panic(err)
	}

	p.WorldMin = t.WorldMin
	p.WorldMax = t.WorldMax

	p.SensorMin = t.SensorMin
	p.SensorMax = t.SensorMax

	p.ActuatorMin = t.ActuatorMin
	p.ActuatorMax = t.ActuatorMax
}

// SetWMinMax ...
func (p *Parameters) SetWMinMax(min []float64, max []float64) {
	p.WorldMin = min
	p.WorldMax = max
}

// SetSMinMax ...
func (p *Parameters) SetSMinMax(min []float64, max []float64) {
	p.SensorMin = min
	p.SensorMax = max
}

// SetAMinMax ...
func (p *Parameters) SetAMinMax(min []float64, max []float64) {
	p.ActuatorMin = min
	p.ActuatorMax = max
}

// SetK ...
func (p *Parameters) SetK(k int) {
	p.K = k
}

// SetOutput ...
func (p *Parameters) SetOutput(output string) {
	p.Output = output
}

// SetVerbose ...
func (p *Parameters) SetVerbose(verbose bool) {
	p.Verbose = verbose
}

// SetContinuousMode ...
func (p *Parameters) SetContinuousMode(cm int) {
	p.ContinuousMode = cm
}

// CfgT ...
type CfgT struct {
	Measure        string `yaml:"Measure"`
	Continuous     bool   `yaml:"Continuous"`
	ContinuousMode int    `yaml:"Continuous mode"`
	UseState       bool   `yaml:"State-dependent"`
	Verbose        bool   `yaml:"Verbose"`
	Bins           int    `yaml:"Bins"`
	Iterations     int    `yaml:"Iterations"`
	K              int    `yaml:"k"`
	Output         string `yaml:"Output file"`
	WBins          string `yaml:"W Bins"`
	ABins          string `yaml:"A Bins"`
	SBins          string `yaml:"S Bins"`
	WIndices       string `yaml:"W Indices"`
	AIndices       string `yaml:"A Indices"`
	SIndices       string `yaml:"S Indices"`
	File           string `yaml:"Full data file"`
	WFile          string `yaml:"W data file"`
	AFile          string `yaml:"A data file"`
	SFile          string `yaml:"S data file"`
	DFile          string `yaml:"Domain file"`
}

// SetConfigFile ..
func (p *Parameters) SetConfigFile(file string) {
	p.ConfigFile = file
	if p.ConfigFile == "" {
		return
	}

	t := CfgT{}

	data, err := ioutil.ReadFile(p.ConfigFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		panic(err)
	}

	p.SetMeasureName(t.Measure)
	p.SetUseContinuous(t.Continuous)
	p.SetContinuousMode(t.ContinuousMode)
	p.SetUseStateDependent(t.UseState)
	p.SetGlobalBins(t.Bins)
	p.SetWBins(t.WBins)
	p.SetSBins(t.SBins)
	p.SetABins(t.ABins)
	p.SetK(t.K)
	p.SetOutput(t.Output)
	p.SetVerbose(t.Verbose)
	p.SetGlobalFile(t.File)
	p.SetWIndices(t.WIndices)
	p.SetSIndices(t.SIndices)
	p.SetAIndices(t.AIndices)
	p.SetWFile(t.WFile)
	p.SetSFile(t.SFile)
	p.SetAFile(t.AFile)
	p.SetDFile(t.DFile)
	p.SetIterations(t.Iterations)

	p.Verbose = true
}

// checkFile returns true, if the file exists and false otherwise
func checkFile(filename string) bool {
	if filename == "" {
		return true
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// CheckParameters checks for the sanity of the command line parameters
func (p *Parameters) CheckParameters() {
	errorMsg := ""
	hasError := false
	var foundFile bool

	foundFile = checkFile(p.GlobalFile)
	hasError = foundFile && hasError
	if foundFile == false {
		errorMsg = fmt.Sprintf("Global file %s not found\n", p.GlobalFile)
	}

	foundFile = checkFile(p.WFile)
	hasError = foundFile && hasError
	if foundFile == true {
		errorMsg = fmt.Sprintf("%sWorld file %s not found\n", errorMsg, p.WFile)
	}

	foundFile = checkFile(p.SFile)
	hasError = foundFile && hasError
	if foundFile == true {
		errorMsg = fmt.Sprintf("%sSensor file %s not found\n", errorMsg, p.SFile)
	}

	foundFile = checkFile(p.AFile)
	hasError = foundFile && hasError
	if foundFile == true {
		errorMsg = fmt.Sprintf("%sActuator file %s not found\n", errorMsg, p.AFile)
	}

	if hasError == true {
		fmt.Println(errorMsg)
		os.Exit(-1)
	}
}
