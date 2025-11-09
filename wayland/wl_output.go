package wayland

import (
	"github.com/mmulet/term.everything/wayland/protocols"
)

type WlOutput struct {
	Version uint32
}

func (o *WlOutput) WlOutput_release(s protocols.ClientState, _ protocols.ObjectID[protocols.WlOutput]) bool {
	return true
}

func (o *WlOutput) WlOutput_on_bind(
	s protocols.ClientState,
	_ protocols.ObjectID[protocols.WlOutput],
	_ string,
	newID protocols.ObjectID[protocols.WlOutput],
	version uint32,
) {
	o.Version = version

	protocols.WlOutput_scale(s, o.Version, newID, 1)

	protocols.WlOutput_name(s, o.Version, newID, "term.everything Virtual Monitor")
	protocols.WlOutput_description(s, o.Version, newID, "The best monitor")

	protocols.WlOutput_geometry(
		s,
		newID,
		0,
		0,
		int32(VirtualMonitorSize.Width),
		int32(VirtualMonitorSize.Height),
		int32(protocols.WlOutputSubpixel_enum_unknown),
		"Very Good",
		"The best model",
		int32(protocols.WlOutputTransform_enum_normal),
	)

	protocols.WlOutput_mode(
		s,
		newID,
		protocols.WlOutputMode_enum_current,
		int32(VirtualMonitorSize.Width),
		int32(VirtualMonitorSize.Height),
		60_000,
	)

	protocols.WlOutput_done(s, version, newID)
}

func MakeWlOutput() *protocols.WlOutput {
	return &protocols.WlOutput{
		Delegate: &WlOutput{
			Version: 1,
		},
	}
}
