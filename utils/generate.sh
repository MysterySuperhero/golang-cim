#!/bin/sh

cim_path=${GOPATH}/src/golang-cim
structs="TargetType=Analog,AnalogValue,ConnectivityNode,Discrete,DiscreteValue,MeasurementValue,PowerTransformer,PowerTransformerEnd,Line,ACLineSegment,OperationalLimitSet,CurrentLimit,OperationalLimit,CurrentTemperatureCurve,CurveData,Curve,CheckPoint,VoltageCheckPoint,VoltageTimeCurve,VoltageLimit"

# generate code
cat unmarshall_template.go | genny gen ${structs} > ${cim_path}/unmarshall.go
