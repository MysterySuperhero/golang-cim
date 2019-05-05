package cim

type PowerTransformer struct {
	ConductingEquipment
	PowerTransformerEnd []string
	Terminals           []string
	EquipmentContainer  string
}

const PowerTransformerName = "PowerTransformer"

type PowerTransformerEnd struct {
	TransformerEnd
	x0               float32
	b0               float32
	g0               float32
	ratedS           float32
	b                float32
	ratedU           float32
	g                float32
	r                float32
	x                float32
	r0               float32
	connectionKind   string
	PowerTransformer string
	Terminal         string
}

const PowerTransformerEndName = "PowerTransformerEnd"

// TODO not all fields
type TransformerEnd struct {
	IdentifiedObject IdentifiedObject
	endNumber        int32
	rgRound          float32
	grounded         bool
	magBaseU         float32
	magSatFlux       float32
	bMagSat          float32
	xGround          float32
}

const TransformerEndName = "TransformerEnd"

type Line struct {
	EquipmentContainer
}

const LineName = "Line"

type Conductor struct {
	ConductingEquipment
}

const ConductorName = "Conductor"

type ACLineSegment struct {
	Conductor
	CurrentTemperatureCurves []string
}

const ACLineSegmentName = "ACLineSegment"

type CurrentTemperatureCurve struct {
	Curve
	ACLineSegments []string
}
