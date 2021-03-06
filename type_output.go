package gomi

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// OutputMinMax ...
type OutputMinMax struct {
	Min *[]float64 `json:"min,omitempty"`
	Max *[]float64 `json:"max,omitempty"`
}

// OutputMeasureContinuous ...
type OutputMeasureContinuous struct {
	Mode          *int          `json:"mode,omitempty"`
	K             *int          `json:"k,omitempty"`
	Normalisation *OutputMinMax `json:"normalisation,omitempty"`
}

// OutputBins ...
type OutputBins struct {
	W      *[]int `json:"w,omitempty"`
	S      *[]int `json:"s,omitempty"`
	A      *[]int `json:"a,omitempty"`
	Global *int   `json:"global,omitempty"`
}

// OutputMeasureDiscrete ...
type OutputMeasureDiscrete struct {
	Iterations *int        `json:"iterations,omitempty"`
	Bins       *OutputBins `json:"bins,omitempty"`
}

// OutputMeasure ...
type OutputMeasure struct {
	Name              *string                  `json:"name,omitempty"`
	UseContinuous     *bool                    `json:"useContinuous,omitempty"`
	UseStateDependent *bool                    `json:"stateDependent,omitempty"`
	Continuous        *OutputMeasureContinuous `json:"continuous,omitempty"`
	Discrete          *OutputMeasureDiscrete   `json:"discrete,omitempty"`
}

// OutputResult ...
type OutputResult struct {
	Average   *float64   `json:"averaged,omitempty"`
	PointWise *[]float64 `json:"point-wise,omitempty"`
}

// OutputDomainFile ...
type OutputDomainFile struct {
	Name     *string       `json:"name,omitempty"`
	World    *OutputMinMax `json:"world,omitempty"`
	Sensor   *OutputMinMax `json:"sensors,omitempty"`
	Actuator *OutputMinMax `json:"actuators,omitempty"`
}

// OutputDataFile ...
type OutputDataFile struct {
	Global *string           `json:"global,omitempty"`
	W      *string           `json:"world,omitempty"`
	S      *string           `json:"sensors,omitempty"`
	A      *string           `json:"actuators,omitempty"`
	Domain *OutputDomainFile `json:"domain,omitempty"`
}

// OutputDataIndices ...
type OutputDataIndices struct {
	W *[]int `json:"world,omitempty"`
	S *[]int `json:"sensors,omitempty"`
	A *[]int `json:"actuators,omitempty"`
}

// OutputData ...
type OutputData struct {
	File    *OutputDataFile    `json:"file,omitempty"`
	Indices *OutputDataIndices `json:"indices,omitempty"`
}

// OutputDataRawNormalised ...
type OutputDataRawNormalised struct {
	Raw        *[][]float64 `json:"raw,omitempty"`
	Normalised *[][]float64 `json:"normalised,omitempty"`
}

// Output is the JSON struct that will be exported as result of gomi
type Output struct {
	Date    *string        `json:"date,omitempty"`
	Measure *OutputMeasure `json:"measure,omitempty"`
	Result  *OutputResult  `json:"result,omitempty"`
	Data    *OutputData    `json:"data,omitempty"`

	W2W1A1   *OutputDataRawNormalised `json:"w2w1a1,omitempty"`
	W2W1S1A1 *OutputDataRawNormalised `json:"w2w1s1a1,omitempty"`
}

// CreateW2W1A1 ...
func (o *Output) CreateW2W1A1() {
	if o.W2W1A1 == nil {
		var r OutputDataRawNormalised
		o.W2W1A1 = &r
	}
}

// SetW2W1A1Raw ...
func (o *Output) SetW2W1A1Raw(data [][]float64) {
	o.CreateW2W1A1()
	o.W2W1A1.Raw = &data
}

// SetW2W1A1Normalised ...
func (o *Output) SetW2W1A1Normalised(data [][]float64) {
	o.CreateW2W1A1()
	o.W2W1A1.Normalised = &data
}

// CreateW2W1S1A1 ...
func (o *Output) CreateW2W1S1A1() {
	if o.W2W1S1A1 == nil {
		var r OutputDataRawNormalised
		o.W2W1S1A1 = &r
	}
}

// SetW2W1S1A1Raw ...
func (o *Output) SetW2W1S1A1Raw(data [][]float64) {
	o.CreateW2W1S1A1()
	o.W2W1S1A1.Raw = &data
}

// SetW2W1S1A1Normalised ...
func (o *Output) SetW2W1S1A1Normalised(data [][]float64) {
	o.CreateW2W1S1A1()
	o.W2W1S1A1.Normalised = &data
}

// ExportJSON exports to JSON
func (o Output) ExportJSON(filename string) {
	bytes, _ := json.MarshalIndent(o, "", " ")
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	f.Write(bytes)
}

// CreateResults ...
func (o *Output) CreateResults() {
	if o.Result == nil {
		var r OutputResult
		o.Result = &r
	}
}

// SetAvgResult sets the name
func (o *Output) SetAvgResult(r float64) {
	o.CreateResults()
	o.Result.Average = &r
}

// SetPointWiseResult sets the name
func (o *Output) SetPointWiseResult(r []float64) {
	o.CreateResults()
	o.Result.PointWise = &r
}

// CreateMeasure ...
func (o *Output) CreateMeasure() {
	if o.Measure == nil {
		var m OutputMeasure
		o.Measure = &m
	}
}

// SetName sets the name
func (o *Output) SetName(name string) {
	o.CreateMeasure()
	o.Measure.Name = &name
}

// SetUseContinuous sets the name
func (o *Output) SetUseContinuous(b bool) {
	o.CreateMeasure()
	o.Measure.UseContinuous = &b
}

// SetUseStateDependent sets the name
func (o *Output) SetUseStateDependent(b bool) {
	o.CreateMeasure()
	o.Measure.UseStateDependent = &b
}

// CreateBins ...
func (o *Output) CreateBins() {
	if o.Measure == nil {
		var m OutputMeasure
		o.Measure = &m
	}
	if o.Measure.Discrete == nil {
		var d OutputMeasureDiscrete
		o.Measure.Discrete = &d
	}
	if o.Measure.Discrete.Bins == nil {
		var b OutputBins
		o.Measure.Discrete.Bins = &b
	}
}

// SetABins ...
func (o *Output) SetABins(bins []int) {
	if len(bins) == 0 {
		return
	}
	o.CreateBins()
	o.Measure.Discrete.Bins.A = &bins
}

// SetWBins ...
func (o *Output) SetWBins(bins []int) {
	if len(bins) == 0 {
		return
	}
	o.CreateBins()
	o.Measure.Discrete.Bins.W = &bins
}

// SetSBins ...
func (o *Output) SetSBins(bins []int) {
	if len(bins) == 0 {
		return
	}
	o.CreateBins()
	o.Measure.Discrete.Bins.S = &bins
}

// SetGlobalBins ...
func (o *Output) SetGlobalBins(bins int) {
	if bins == 0 {
		return
	}
	o.CreateBins()
	o.Measure.Discrete.Bins.Global = &bins
}

// CreateMeasureDiscrete ...
func (o *Output) CreateMeasureDiscrete() {
	o.CreateMeasure()
	if o.Measure.Discrete == nil {
		var d OutputMeasureDiscrete
		o.Measure.Discrete = &d
	}
}

// SetIterations ...
func (o *Output) SetIterations(iterations int) {
	if iterations == 0 {
		return
	}
	o.CreateMeasureDiscrete()
	o.Measure.Discrete.Iterations = &iterations
}

// SetNormalisation ...
func (o *Output) SetNormalisation(min, max []float64) {
	if len(min) == 0 && len(max) == 0 {
		fmt.Println("end end")
		return
	}
	o.CreateMinMax()
	if len(min) > 0 {
		o.Measure.Continuous.Normalisation.Min = &min
	}
	if len(max) > 0 {
		o.Measure.Continuous.Normalisation.Max = &max
	}
}

// CreateMinMax ...
func (o *Output) CreateMinMax() {
	o.CreateContinuous()
	if o.Measure.Continuous.Normalisation == nil {
		var n OutputMinMax
		o.Measure.Continuous.Normalisation = &n
	}
}

// CreateContinuous ...
func (o *Output) CreateContinuous() {
	o.CreateMeasure()
	if o.Measure.Continuous == nil {
		c := OutputMeasureContinuous{}
		o.Measure.Continuous = &c
	}
}

// SetK ...
func (o *Output) SetK(k int) {
	if k == 0 {
		return
	}
	o.CreateContinuous()
	o.Measure.Continuous.K = &k
}

// SetContinuousMode ...
func (o *Output) SetContinuousMode(mode int) {
	o.CreateContinuous()
	o.Measure.Continuous.Mode = &mode
}

// SetGlobalFile ...
func (o *Output) SetGlobalFile(name string) {
	if name == "" {
		return
	}
	o.CreateDataFile()
	o.Data.File.Global = &name
}

// SetWFile ...
func (o *Output) SetWFile(name string) {
	if name == "" {
		return
	}
	o.CreateDataFile()
	o.Data.File.W = &name
}

// SetAFile ...
func (o *Output) SetAFile(name string) {
	if name == "" {
		return
	}
	o.CreateDataFile()
	o.Data.File.A = &name
}

// SetSFile ...
func (o *Output) SetSFile(name string) {
	if name == "" {
		return
	}
	o.CreateDataFile()
	o.Data.File.S = &name
}

// CreateData ...
func (o *Output) CreateData() {
	if o.Data == nil {
		d := OutputData{}
		o.Data = &d
	}
}

// CreateDataFile ...
func (o *Output) CreateDataFile() {
	o.CreateData()
	if o.Data.File == nil {
		f := OutputDataFile{}
		o.Data.File = &f
	}
}

// CreateDomainFile ...
func (o *Output) CreateDomainFile() {
	o.CreateDataFile()
	if o.Data.File.Domain == nil {
		d := OutputDomainFile{}
		o.Data.File.Domain = &d
	}
}

// CreateDomainWFile ...
func (o *Output) CreateDomainWFile() {
	o.CreateDomainFile()
	if o.Data.File.Domain.World == nil {
		m := OutputMinMax{}
		o.Data.File.Domain.World = &m
	}
}

// CreateDomainSFile ...
func (o *Output) CreateDomainSFile() {
	o.CreateDomainFile()
	if o.Data.File.Domain.Sensor == nil {
		m := OutputMinMax{}
		o.Data.File.Domain.Sensor = &m
	}
}

// CreateDomainAFile ...
func (o *Output) CreateDomainAFile() {
	o.CreateDomainFile()
	if o.Data.File.Domain.Actuator == nil {
		m := OutputMinMax{}
		o.Data.File.Domain.Actuator = &m
	}
}

// CreateDataIndices ...
func (o *Output) CreateDataIndices() {
	o.CreateData()
	if o.Data.Indices == nil {
		i := OutputDataIndices{}
		o.Data.Indices = &i
	}
}

// SetWIndices ...
func (o *Output) SetWIndices(indices []int) {
	if len(indices) == 0 {
		return
	}
	o.CreateDataIndices()
	o.Data.Indices.W = &indices
}

// SetAIndices ...
func (o *Output) SetAIndices(indices []int) {
	if len(indices) == 0 {
		return
	}
	o.CreateDataIndices()
	o.Data.Indices.A = &indices
}

// SetSIndices ...
func (o *Output) SetSIndices(indices []int) {
	if len(indices) == 0 {
		return
	}
	o.CreateDataIndices()
	o.Data.Indices.S = &indices
}

// SetDomainAMinMax ...
func (o *Output) SetDomainAMinMax(min, max []float64) {
	if len(min) == 0 && len(max) == 0 {
		return
	}
	o.CreateDomainAFile()
	if len(min) > 0 {
		o.Data.File.Domain.Actuator.Min = &min
	}
	if len(max) > 0 {
		o.Data.File.Domain.Actuator.Max = &max
	}
}

// SetDomainWMinMax ...
func (o *Output) SetDomainWMinMax(min, max []float64) {
	if len(min) == 0 && len(max) == 0 {
		return
	}
	o.CreateDomainWFile()
	if len(min) > 0 {
		o.Data.File.Domain.World.Min = &min
	}
	if len(max) > 0 {
		o.Data.File.Domain.World.Max = &max
	}
}

// SetDomainSMinMax ...
func (o *Output) SetDomainSMinMax(min, max []float64) {
	if len(min) == 0 && len(max) == 0 {
		return
	}
	o.CreateDomainSFile()
	if len(min) > 0 {
		o.Data.File.Domain.Sensor.Min = &min
	}
	if len(max) > 0 {
		o.Data.File.Domain.Sensor.Max = &max
	}
}

// SetDomainName ...
func (o *Output) SetDomainName(name string) {
	if name == "" {
		return
	}
	o.CreateDomainFile()
	o.Data.File.Domain.Name = &name
}

// SetDate ...
func (o *Output) SetDate() {
	o.CreateMeasure()
	s := time.Now().Format("2006-01-02 15:04:05")
	o.Date = &s
}

// SetParameters copies from Parameters
func (o *Output) SetParameters(p Parameters) {

	o.SetName(p.MeasureName)
	o.SetABins(p.ABins)
	o.SetWBins(p.WBins)
	o.SetSBins(p.SBins)
	o.SetGlobalBins(p.GlobalBins)
	o.SetIterations(p.Iterations)
	o.SetNormalisation(p.NormalisationMin, p.NormalisationMax)
	if p.UseContinuous {
		o.SetK(p.K)
		o.SetContinuousMode(p.ContinuousMode)
	}

	o.SetUseContinuous(p.UseContinuous)
	o.SetUseStateDependent(p.UseStateDependent)
	o.SetGlobalFile(p.GlobalFile)
	o.SetWFile(p.WFile)
	o.SetAFile(p.AFile)
	o.SetSFile(p.SFile)

	o.SetWIndices(p.WIndices)
	o.SetAIndices(p.AIndices)
	o.SetSIndices(p.SIndices)

	o.SetDomainWMinMax(p.WorldMin, p.WorldMax)
	o.SetDomainAMinMax(p.ActuatorMin, p.ActuatorMax)
	o.SetDomainSMinMax(p.SensorMin, p.SensorMax)
	o.SetDomainName(p.DFile)

}
