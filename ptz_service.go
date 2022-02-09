package onvif

import (
	"github.com/inspii/onvif/onvif"
	"github.com/inspii/onvif/ptz"
	"github.com/inspii/onvif/xsd"
	"time"
)

type PTZService struct {
	Client onvifCaller
}

func NewPTZService(endpoint string, onvifDevice *OnvifDevice) *PTZService {
	return &PTZService{
		Client: NewOnvifClient(endpoint, &onvifDevice.auth),
	}
}

func (s *PTZService) WithoutAuth() *PTZService {
	return &PTZService{
		Client: s.Client.WithoutAuth(),
	}
}

func (s *PTZService) ContinuousMove(profileToken string, x, y, zoom float64, timeout time.Duration) (res ptz.ContinuousMoveResponse, err error) {
	req := ptz.ContinuousMove{
		ProfileToken: onvif.ReferenceToken(profileToken),
		Velocity: onvif.PTZSpeed{
			PanTilt: onvif.Vector2D{
				X:     x,
				Y:     y,
				Space: "http://www.onvif.org/ver10/tptz/PanTiltSpaces/VelocityGenericSpace",
			},
			Zoom: onvif.Vector1D{
				X:     zoom,
				Space: "http://www.onvif.org/ver10/tptz/ZoomSpaces/VelocityGenericSpace",
			},
		},
		Timeout: xsd.Duration(timeout),
	}
	err = s.Client.Call(req, &res)
	return
}

func (s *PTZService) Stop(profileToken string) (res ptz.StopResponse, err error) {
	req := ptz.Stop{
		ProfileToken: onvif.ReferenceToken(profileToken),
	}
	err = s.Client.Call(req, &res)
	return
}
