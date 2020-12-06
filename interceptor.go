// +build !js

package webrtc

import (
	"github.com/pion/interceptor"
	"github.com/pion/interceptor/pkg/nack"
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
	generator, err := nack.NewGeneratorInterceptor()
	if err != nil {
		return err
	}
	interceptorRegistry.Add(generator)
	responder, err := nack.NewResponderInterceptor()
	if err != nil {
		return err
	}
	interceptorRegistry.Add(responder)
	return nil
}
