package skeleton

import (
	"fmt"
	"net/http"

	"github.com/merliot/dean"
	"github.com/merliot/device"
	"github.com/merliot/device/led"
)

var targets = []string{"demo", "nano-rp2040", "wioterminal"}

type Skeleton struct {
	*device.Device
	Led led.Led
}

type MsgClick struct {
	dean.ThingMsg
	State bool
}

func New(id, model, name string) dean.Thinger {
	fmt.Println("NEW SKELETON\r")
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
	s.Led.Gpio = s.ParamFirstValue("gpio")
}

func (s *Skeleton) configure() {
	s.Led.Configure()
}

func (s *Skeleton) Setup() {
	s.Device.Setup()
	s.parseParams()
	s.configure()
}

func (s *Skeleton) Run(i *dean.Injector) {
	select {}
}
