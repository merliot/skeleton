package skeleton

import (
	"embed"
	"net/http"

	"github.com/merliot/dean"
	"github.com/merliot/device"
	"github.com/merliot/device/led"
)

//go:embed css go.mod html images js template
var fs embed.FS

var targets = []string{"demo", "nano-rp2040"}

type Skeleton struct {
	*device.Device
	Led led.Led
}

type MsgClick struct {
	dean.ThingMsg
	State bool
}

func New(id, model, name string) dean.Thinger {
	println("NEW SKELETON")
	return &Skeleton{
		Device: device.New(id, model, name, fs, targets).(*device.Device),
	}
}

func (s *Skeleton) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.API(w, r, s)
}

func (s *Skeleton) save(msg *dean.Msg) {
	msg.Unmarshal(s).Broadcast()
}

func (s *Skeleton) getState(msg *dean.Msg) {
	s.Path = "state"
	msg.Marshal(s).Reply()
}

func (s *Skeleton) click(msg *dean.Msg) {
	msg.Unmarshal(&s.Led)
	if s.IsMetal() {
		s.Led.Set(s.Led.State)
	}
	msg.Broadcast()
}

func (s *Skeleton) Subscribers() dean.Subscribers {
	return dean.Subscribers{
		"state":     s.save,
		"get/state": s.getState,
		"click":     s.click,
	}
}

func (s *Skeleton) parseParams() {
	values := s.ParseDeployParams()
	s.Led.Gpio = s.ParamFirstValue(values, "gpio")
	s.Led.Configure()
}

func (s *Skeleton) Setup() {
	s.Device.Setup()
	s.parseParams()
}

func (s *Skeleton) Run(i *dean.Injector) {
	select {}
}
