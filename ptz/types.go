package ptz

import (
	"github.com/inspii/onvif/onvif"
	"github.com/inspii/onvif/xsd"
)

type Capabilities struct {
	EFlip                       xsd.Boolean `xml:"EFlip,attr"`
	Reverse                     xsd.Boolean `xml:"Reverse,attr"`
	GetCompatibleConfigurations xsd.Boolean `xml:"GetCompatibleConfigurations,attr"`
	MoveStatus                  xsd.Boolean `xml:"MoveStatus,attr"`
	StatusPosition              xsd.Boolean `xml:"StatusPosition,attr"`
}

//PTZ main types

type GetServiceCapabilities struct {
	XMLName string `xml:"http://www.onvif.org/ver10/ptz/wsdl GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	Capabilities Capabilities
}

type GetNodes struct {
	XMLName string `xml:"http://www.onvif.org/ver10/ptz/wsdl GetNodes"`
}

type GetNodesResponse struct {
	PTZNode onvif.PTZNode
}

type GetNode struct {
	XMLName   string               `xml:"http://www.onvif.org/ver10/ptz/wsdl GetNode"`
	NodeToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl NodeToken"`
}

type GetNodeResponse struct {
	PTZNode onvif.PTZNode
}

type GetConfiguration struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl GetConfiguration"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
}

type GetConfigurationResponse struct {
	PTZConfiguration onvif.PTZConfiguration
}

type GetConfigurations struct {
	XMLName string `xml:"http://www.onvif.org/ver10/ptz/wsdl GetConfigurations"`
}

type GetConfigurationsResponse struct {
	PTZConfiguration onvif.PTZConfiguration
}

type SetConfiguration struct {
	XMLName          string                 `xml:"http://www.onvif.org/ver10/ptz/wsdl SetConfiguration"`
	PTZConfiguration onvif.PTZConfiguration `xml:"http://www.onvif.org/ver10/ptz/wsdl PTZConfiguration"`
	ForcePersistence xsd.Boolean            `xml:"http://www.onvif.org/ver10/ptz/wsdl ForcePersistence"`
}

type SetConfigurationResponse struct {
}

type GetConfigurationOptions struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl GetConfigurationOptions"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
}

type GetConfigurationOptionsResponse struct {
	PTZConfigurationOptions onvif.PTZConfigurationOptions
}

type SendAuxiliaryCommand struct {
	XMLName       string               `xml:"http://www.onvif.org/ver10/ptz/wsdl SendAuxiliaryCommand"`
	ProfileToken  onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	AuxiliaryData onvif.AuxiliaryData  `xml:"http://www.onvif.org/ver10/ptz/wsdl AuxiliaryData"`
}

type SendAuxiliaryCommandResponse struct {
	AuxiliaryResponse onvif.AuxiliaryData
}

type GetPresets struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl GetPresets"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
}

type GetPresetsResponse struct {
	Preset onvif.PTZPreset
}

type SetPreset struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl SetPreset"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	PresetName   xsd.String           `xml:"http://www.onvif.org/ver10/ptz/wsdl PresetName"`
	PresetToken  onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl PresetToken"`
}

type SetPresetResponse struct {
	PresetToken onvif.ReferenceToken
}

type RemovePreset struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl RemovePreset"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	PresetToken  onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl PresetToken"`
}

type RemovePresetResponse struct {
}

type GotoPreset struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl GotoPreset"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	PresetToken  onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl PresetToken"`
	Speed        onvif.PTZSpeed       `xml:"http://www.onvif.org/ver10/ptz/wsdl Speed"`
}

type GotoPresetResponse struct {
}

type GotoHomePosition struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl GotoHomePosition"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	Speed        onvif.PTZSpeed       `xml:"http://www.onvif.org/ver10/ptz/wsdl Speed"`
}

type GotoHomePositionResponse struct {
}

type SetHomePosition struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl SetHomePosition"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
}

type SetHomePositionResponse struct {
}

type ContinuousMove struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl ContinuousMove"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	Velocity     onvif.PTZSpeed       `xml:"http://www.onvif.org/ver10/ptz/wsdl Velocity"`
	Timeout      xsd.Duration         `xml:"http://www.onvif.org/ver10/ptz/wsdl Timeout"`
}

type ContinuousMoveResponse struct {
}

type RelativeMove struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl RelativeMove"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	Translation  onvif.PTZVector      `xml:"http://www.onvif.org/ver10/ptz/wsdl Translation"`
	Speed        onvif.PTZSpeed       `xml:"http://www.onvif.org/ver10/ptz/wsdl Speed"`
}

type RelativeMoveResponse struct {
}

type GetStatus struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl GetStatus"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
}

type GetStatusResponse struct {
	PTZStatus onvif.PTZStatus
}

type AbsoluteMove struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl AbsoluteMove"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	Position     onvif.PTZVector      `xml:"http://www.onvif.org/ver10/ptz/wsdl Position"`
	Speed        onvif.PTZSpeed       `xml:"http://www.onvif.org/ver10/ptz/wsdl Speed"`
}

type AbsoluteMoveResponse struct {
}

type GeoMove struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl GeoMove"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	Target       onvif.GeoLocation    `xml:"http://www.onvif.org/ver10/ptz/wsdl Target"`
	Speed        onvif.PTZSpeed       `xml:"http://www.onvif.org/ver10/ptz/wsdl Speed"`
	AreaHeight   xsd.Float            `xml:"http://www.onvif.org/ver10/ptz/wsdl AreaHeight"`
	AreaWidth    xsd.Float            `xml:"http://www.onvif.org/ver10/ptz/wsdl AreaWidth"`
}

type GeoMoveResponse struct {
}

type Stop struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl Stop"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	PanTilt      xsd.Boolean          `xml:"http://www.onvif.org/ver10/ptz/wsdl PanTilt"`
	Zoom         xsd.Boolean          `xml:"http://www.onvif.org/ver10/ptz/wsdl Zoom"`
}

type StopResponse struct {
}

type GetPresetTours struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl GetPresetTours"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
}

type GetPresetToursResponse struct {
	PresetTour onvif.PresetTour
}

type GetPresetTour struct {
	XMLName         string               `xml:"http://www.onvif.org/ver10/ptz/wsdl GetPresetTour"`
	ProfileToken    onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	PresetTourToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl PresetTourToken"`
}

type GetPresetTourResponse struct {
	PresetTour onvif.PresetTour
}

type GetPresetTourOptions struct {
	XMLName         string               `xml:"http://www.onvif.org/ver10/ptz/wsdl GetPresetTourOptions"`
	ProfileToken    onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	PresetTourToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl PresetTourToken"`
}

type GetPresetTourOptionsResponse struct {
	Options onvif.PTZPresetTourOptions
}

type CreatePresetTour struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl CreatePresetTour"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
}

type CreatePresetTourResponse struct {
	PresetTourToken onvif.ReferenceToken
}

type ModifyPresetTour struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl ModifyPresetTour"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	PresetTour   onvif.PresetTour     `xml:"http://www.onvif.org/ver10/ptz/wsdl PresetTour"`
}

type ModifyPresetTourResponse struct {
}

type OperatePresetTour struct {
	XMLName         string                       `xml:"http://www.onvif.org/ver10/ptz/wsdl OperatePresetTour"`
	ProfileToken    onvif.ReferenceToken         `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	PresetTourToken onvif.ReferenceToken         `xml:"onvif:PresetTourToken"`
	Operation       onvif.PTZPresetTourOperation `xml:"onvif:Operation"`
}

type OperatePresetTourResponse struct {
}

type RemovePresetTour struct {
	XMLName         string               `xml:"http://www.onvif.org/ver10/ptz/wsdl RemovePresetTour"`
	ProfileToken    onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
	PresetTourToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl PresetTourToken"`
}

type RemovePresetTourResponse struct {
}

type GetCompatibleConfigurations struct {
	XMLName      string               `xml:"http://www.onvif.org/ver10/ptz/wsdl GetCompatibleConfigurations"`
	ProfileToken onvif.ReferenceToken `xml:"http://www.onvif.org/ver10/ptz/wsdl ProfileToken"`
}

type GetCompatibleConfigurationsResponse struct {
	PTZConfiguration onvif.PTZConfiguration
}
