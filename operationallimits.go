package cim

type OperationalLimitSet struct {
	IdentifiedObject
	Terminal              string
	Equipment             string
	OperationalLimitValue []string
}

const OperationalLimitSetName = "OperationalLimitSet"

type OperationalLimit struct {
	IdentifiedObject
	OperationalLimitSet  string
	OperationalLimitType string
}

const OperationalLimitName = "OperationalLimit"

type CurrentLimit struct {
	OperationalLimit
	Value float64
}

const CurrentLimitName = "CurrentLimit"

type CheckPoint struct {
	IdentifiedObject
	Equipment string
}

const CheckPointName = "CurrentLimit"

type VoltageCheckPoint struct {
	CheckPoint
	LowerLimitsCurve string
	UpperLimitsCurve string
	VoltageLimitSet  string
}

const VoltageCheckPointName = "VoltageCheckPoint"

type VoltageTimeCurve struct {
	Curve
}

const VoltageTimeCurveName = "VoltageTimeCurve"

type VoltageLimit struct {
	OperationalLimit
	Value float64
}

const VoltageLimitName = "VoltageLimit"
