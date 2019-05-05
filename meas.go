package cim

type Analog struct {
	Measurement  Measurement
	AnalogValues []AnalogValue
}

const AnalogName = "Analog"

type AnalogValue struct {
	MeasurementValue string
	Analog           string
}

const AnalogValueName = "AnalogValue"

type Discrete struct {
	Measurement    Measurement
	DiscreteValues []DiscreteValue
}

const DiscreteName = "Discrete"


type DiscreteValue struct {
	MeasurementValue MeasurementValue
	Discrete         Discrete
}

const DiscreteValueName = "DiscreteValue"

type Measurement struct {
	IdentifiedObject IdentifiedObject
}

const MeasurementName = "Measurement"


type MeasurementValue struct {
	IdentifiedObject IdentifiedObject
	sensorAccuracy   float64
	timeStamp        int64
}

const MeasurementValueName = "MeasurementValue"
