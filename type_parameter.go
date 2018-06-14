package main

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v1"
)

type Paramters struct {
	MeasureName       string
	UseContinuous     bool
	UseStateDependent bool
	GlobalBins        int
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
	WMin              []float64
	SMin              []float64
	AMin              []float64
	WMax              []float64
	SMax              []float64
	AMax              []float64
}

func (p Paramters) String() string {
	s := ""
	s = fmt.Sprintf("%s\nMeasure:             %s", s, p.MeasureName)
	s = fmt.Sprintf("%s\nUse state-dependent: %t", s, p.UseStateDependent)
	s = fmt.Sprintf("%s\nUse continuous:      %t", s, p.UseContinuous)
	s = fmt.Sprintf("%s\nBins:                %d", s, p.GlobalBins)
	s = fmt.Sprintf("%s\nW bins:              %v", s, p.WBins)
	s = fmt.Sprintf("%s\nS bins:              %v", s, p.SBins)
	s = fmt.Sprintf("%s\nA bins:              %v", s, p.ABins)
	s = fmt.Sprintf("%s\nFull data set:       %s", s, p.GlobalFile)
	s = fmt.Sprintf("%s\nW indices:           %v", s, p.WIndices)
	s = fmt.Sprintf("%s\nS indices:           %v", s, p.SIndices)
	s = fmt.Sprintf("%s\nA indices:           %v", s, p.AIndices)
	s = fmt.Sprintf("%s\nW data set:          %s", s, p.WFile)
	s = fmt.Sprintf("%s\nS data set:          %s", s, p.SFile)
	s = fmt.Sprintf("%s\nA data set:          %s", s, p.AFile)
	s = fmt.Sprintf("%s\nW domains:           %v %v", s, p.WMin, p.WMax)
	s = fmt.Sprintf("%s\nS domains:           %v %v", s, p.SMin, p.SMax)
	s = fmt.Sprintf("%s\nA domains:           %v %v", s, p.AMin, p.AMax)
	return s
}

func CreateParamtersContainer() Paramters {
	return Paramters{MeasureName: "",
		UseContinuous:     false,
		UseStateDependent: false,
		GlobalBins:        0,
		WBins:             []int{},
		SBins:             []int{},
		ABins:             []int{},
		WIndices:          []int{},
		SIndices:          []int{},
		AIndices:          []int{},
		WMin:              []float64{},
		WMax:              []float64{},
		SMin:              []float64{},
		SMax:              []float64{},
		AMin:              []float64{},
		AMax:              []float64{},
		GlobalFile:        "",
		WFile:             "",
		SFile:             "",
		AFile:             ""}
}

func (p *Paramters) AddMeasureName(name string) {
	p.MeasureName = name
}

func (p *Paramters) SetUseStateDependent(b bool) {
	p.UseStateDependent = b
}

func (p *Paramters) SetUseContinuous(b bool) {
	p.UseContinuous = b
}

func (p *Paramters) AddWBins(wBins string) {
	p.WBins = parseIntString(wBins)
}

func (p *Paramters) AddSBins(wBins string) {
	p.SBins = parseIntString(wBins)
}

func (p *Paramters) AddABins(wBins string) {
	p.ABins = parseIntString(wBins)
}

func (p *Paramters) AddWIndices(wIndices string) {
	p.WIndices = parseIntString(wIndices)
}

func (p *Paramters) AddSIndices(wIndices string) {
	p.SIndices = parseIntString(wIndices)
}

func (p *Paramters) AddAIndices(wIndices string) {
	p.AIndices = parseIntString(wIndices)
}

func (p *Paramters) AddGlobalBins(bins int) {
	p.GlobalBins = bins
}

func (p *Paramters) AddGlobalFile(file string) {
	p.GlobalFile = file
}

func (p *Paramters) AddWFile(file string) {
	p.WFile = file
}

func (p *Paramters) AddSFile(file string) {
	p.SFile = file
}

func (p *Paramters) AddAFile(file string) {
	p.AFile = file
}

type T struct {
	Wmin []float64 `yaml:"W min"`
	Wmax []float64 `yaml:"W max"`
	Smin []float64 `yaml:"S min"`
	Smax []float64 `yaml:"S max"`
	Amin []float64 `yaml:"A min"`
	Amax []float64 `yaml:"A max"`
}

func (p *Paramters) AddDFile(file string) {
	p.DFile = file
	if p.DFile == "" {
		return
	}

	t := T{}

	data, err := ioutil.ReadFile(p.DFile)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		panic(err)
	}

	p.WMin = t.Wmin
	p.WMax = t.Wmax

	p.SMin = t.Smin
	p.SMax = t.Smax

	p.AMin = t.Amin
	p.AMax = t.Amax
}

func (p *Paramters) AddWMinMax(min []float64, max []float64) {
	p.WMin = min
	p.WMax = max
}

func (p *Paramters) AddSMinMax(min []float64, max []float64) {
	p.SMin = min
	p.SMax = max
}

func (p *Paramters) AddAMinMax(min []float64, max []float64) {
	p.AMin = min
	p.AMax = max
}
