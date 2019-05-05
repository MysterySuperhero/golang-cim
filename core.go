package cim

type ACDCTerminal struct {
	IdentifiedObject
	connected      bool
	sequenceNumber int32
}

const ACDCTerminalName = "ACDCTerminal"

type ConductingEquipment struct {
	Equipment
}

const ConductingEquipmentName = "ConductingEquipment"

type ConnectivityNode struct {
	IdentifiedObject
	ConnectivityNodeContainer string
	Terminals                 []string
}

const ConnectivityNodeName = "ConnectivityNode"

type Element struct {
	UUID string
}

const ElementName = "Element"

type Equipment struct {
	PowerSystemResource
	NormallyInService   bool
	Aggregate           bool
	OperationalLimitSet []string
	EquipmentContainer  string
}

const EquipmentName = "Equipment"

type IdentifiedObject struct {
	Element
	MRID      string
	AliasName string
	Name      string
	Names     []string
}

const IdentifiedObjectName = "IdentifiedObject"

type Name struct {
	Name string
}

type PowerSystemResource struct {
	IdentifiedObject
	Measurements []string
}

const PowerSystemResourceName = "PowerSystemResource"

type ConnectivityNodeContainer struct {
	PowerSystemResource
}

const ConnectivityNodeContainerName = "ConnectivityNodeContainer"

type EquipmentContainer struct {
	ConnectivityNodeContainer
	Equipments []string
}

const EquipmentContainerName = "EquipmentContainer"

type Curve struct {
	IdentifiedObject
	XUnit      string
	CurveDatas []string
}

const CurveName = "Curve"

type CurveData struct {
	Element
	XValue  float64
	Y1Value float64
	Y2Value float64
	Y3Value float64
	Curve   string
}

const CurveDataName = "CurveData"
