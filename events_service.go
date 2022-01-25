package onvif

import (
	tev "github.com/inspii/onvif/events"
)

type EventsService struct {
	Client   onvifCaller
	endpoint string
}

func NewEventsService(endpoint string, onvifDevice *OnvifDevice) *EventsService {
	return &EventsService{
		Client:   NewOnvifClient(endpoint, &onvifDevice.auth),
		endpoint: endpoint,
	}
}

func (s *EventsService) makeAddressingHeaders(action string) []interface{} {
	return tev.MakeAnonymousAddressingHeaders(action, s.endpoint)
}

// GetServiceCapabilities returns the capabilities of the event service.
func (s *EventsService) GetServiceCapabilities() (res tev.GetServiceCapabilitiesResponse, err error) {
	headers := s.makeAddressingHeaders("http://www.onvif.org/ver10/events/wsdl/EventPortType/GetServiceCapabilitiesRequest")
	err = s.Client.Call(tev.GetServiceCapabilities{}, &res, headers...)
	return
}

/*
GetEventProperties returns information about the FilterDialects, Schema files and
topics supported by the device.

The WS-BaseNotification specification defines a set of OPTIONAL WS-ResouceProperties.
This specification does not require the implementation of the WS-ResourceProperty interface.
Instead, the subsequent direct interface shall be implemented by an ONVIF compliant device
in order to provide information about the FilterDialects, Schema files and topics supported by
the device.
*/
func (s *EventsService) GetEventProperties() (res tev.GetEventPropertiesResponse, err error) {
	headers := s.makeAddressingHeaders("http://www.onvif.org/ver10/events/wsdl/EventPortType/GetEventPropertiesRequest")
	err = s.Client.Call(tev.GetEventProperties{}, &res, headers...)
	return
}

/*
CreatePullPointSubscription returns a PullPointSubscription that can be polled using PullMessages.
This message contains the same elements as the SubscriptionRequest of the WS-BaseNotification
without the ConsumerReference.

If no Filter is specified the pullpoint notifies all occurring events to the client.
*/
func (s *EventsService) CreatePullPointSubscription(filter string, changeOnly bool, initialTerminationTime *tev.AbsoluteOrRelativeTimeType) (res tev.CreatePullPointSubscriptionResponse, err error) {
	headers := s.makeAddressingHeaders("http://www.onvif.org/ver10/events/wsdl/EventPortType/CreatePullPointSubscriptionRequest")
	err = s.Client.Call(tev.CreatePullPointSubscription{
		Filter: tev.FilterType(filter),
		SubscriptionPolicy: &tev.SubscriptionPolicy{
			ChangedOnly: changeOnly,
		},
		InitialTerminationTime: initialTerminationTime,
	}, &res, headers...)

	/*
		<wsdl:fault name="ResourceUnknownFault" message="wsrf-rw:ResourceUnknownFault"/>
		<wsdl:fault name="InvalidFilterFault" message="wsntw:InvalidFilterFault"/>
		<wsdl:fault name="TopicExpressionDialectUnknownFault" message="wsntw:TopicExpressionDialectUnknownFault"/>
		<wsdl:fault name="InvalidTopicExpressionFault" message="wsntw:InvalidTopicExpressionFault"/>
		<wsdl:fault name="TopicNotSupportedFault" message="wsntw:TopicNotSupportedFault"/>
		<wsdl:fault name="InvalidProducerPropertiesExpressionFault" message="wsntw:InvalidProducerPropertiesExpressionFault"/>
		<wsdl:fault name="InvalidMessageContentExpressionFault" message="wsntw:InvalidMessageContentExpressionFault"/>
		<wsdl:fault name="UnacceptableInitialTerminationTimeFault" message="wsntw:UnacceptableInitialTerminationTimeFault"/>
		<wsdl:fault name="UnrecognizedPolicyRequestFault" message="wsntw:UnrecognizedPolicyRequestFault"/>
		<wsdl:fault name="UnsupportedPolicyRequestFault" message="wsntw:UnsupportedPolicyRequestFault"/>
		<wsdl:fault name="NotifyMessageNotSupportedFault" message="wsntw:NotifyMessageNotSupportedFault"/>
		<wsdl:fault name="SubscribeCreationFailedFault" message="wsntw:SubscribeCreationFailedFault"/>
	*/

	return
}
