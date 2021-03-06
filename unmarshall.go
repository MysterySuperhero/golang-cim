// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package cim

import "encoding/json"

func UnMarshallAnalog(rawValue json.RawMessage) (*Analog, error) {
	var result Analog
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallAnalogValue(rawValue json.RawMessage) (*AnalogValue, error) {
	var result AnalogValue
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallConnectivityNode(rawValue json.RawMessage) (*ConnectivityNode, error) {
	var result ConnectivityNode
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallDiscrete(rawValue json.RawMessage) (*Discrete, error) {
	var result Discrete
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallDiscreteValue(rawValue json.RawMessage) (*DiscreteValue, error) {
	var result DiscreteValue
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallMeasurementValue(rawValue json.RawMessage) (*MeasurementValue, error) {
	var result MeasurementValue
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallPowerTransformer(rawValue json.RawMessage) (*PowerTransformer, error) {
	var result PowerTransformer
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallPowerTransformerEnd(rawValue json.RawMessage) (*PowerTransformerEnd, error) {
	var result PowerTransformerEnd
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallLine(rawValue json.RawMessage) (*Line, error) {
	var result Line
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallACLineSegment(rawValue json.RawMessage) (*ACLineSegment, error) {
	var result ACLineSegment
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallOperationalLimitSet(rawValue json.RawMessage) (*OperationalLimitSet, error) {
	var result OperationalLimitSet
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallCurrentLimit(rawValue json.RawMessage) (*CurrentLimit, error) {
	var result CurrentLimit
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallOperationalLimit(rawValue json.RawMessage) (*OperationalLimit, error) {
	var result OperationalLimit
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallCurrentTemperatureCurve(rawValue json.RawMessage) (*CurrentTemperatureCurve, error) {
	var result CurrentTemperatureCurve
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallCurveData(rawValue json.RawMessage) (*CurveData, error) {
	var result CurveData
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallCurve(rawValue json.RawMessage) (*Curve, error) {
	var result Curve
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallCheckPoint(rawValue json.RawMessage) (*CheckPoint, error) {
	var result CheckPoint
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallVoltageCheckPoint(rawValue json.RawMessage) (*VoltageCheckPoint, error) {
	var result VoltageCheckPoint
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallVoltageTimeCurve(rawValue json.RawMessage) (*VoltageTimeCurve, error) {
	var result VoltageTimeCurve
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}

func UnMarshallVoltageLimit(rawValue json.RawMessage) (*VoltageLimit, error) {
	var result VoltageLimit
	err := json.Unmarshal(rawValue, &result)
	return &result, err
}
