// +build !js

package webrtc

import (
	"github.com/pion/interceptor"
	"github.com/pion/rtp"
)

// RegisterDefaultInterceptors will register some useful interceptors. If you want to customize which interceptors are loaded,
// you should copy the code from this method and remove unwanted interceptors.
func RegisterDefaultInterceptors(mediaEngine *MediaEngine, interceptorRegistry *interceptor.Registry) error {
	err := ConfigureNack(mediaEngine, interceptorRegistry)
	if err != nil {
		return err
	}

	return nil
}

// ConfigureNack will setup everything necessary for handling generating/responding to nack messages.
func ConfigureNack(mediaEngine *MediaEngine, interceptorRegistry *interceptor.Registry) error {
	mediaEngine.RegisterFeedback(RTCPFeedback{Type: "nack"}, RTPCodecTypeVideo)
	mediaEngine.RegisterFeedback(RTCPFeedback{Type: "nack", Parameter: "pli"}, RTPCodecTypeVideo)
	interceptorRegistry.Add(&interceptor.NACK{})
	return nil
}

// TODO
type interceptorStreamAdapter struct{ interceptor interceptor.RTPWriter }

func (i *interceptorStreamAdapter) WriteRTP(header *rtp.Header, payload []byte) (int, error) {
	return 0, nil
}
func (i *interceptorStreamAdapter) Write(b []byte) (int, error) { return 0, nil }
