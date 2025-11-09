package protocols

type GlobalID AnyObjectID
type Version uint32

type GlobalWlDisplayID GlobalID
type GlobalWlCompositorID GlobalID
type GlobalWlSubcompositorID GlobalID
type GlobalWlOutputID GlobalID
type GlobalWlSeatID GlobalID
type GlobalWlShmID GlobalID
type GlobalXdgWmBaseID GlobalID
type GlobalWlDataDeviceManagerID GlobalID
type GlobalWlKeyboardID GlobalID
type GlobalWlPointerID GlobalID
type GlobalZwpXwaylandKeyboardGrabManagerV1ID GlobalID
type GlobalXwaylandShellV1ID GlobalID
type GlobalWlDataDeviceID GlobalID
type GlobalWlTouchID GlobalID
type GlobalZxdgDecorationManagerV1ID GlobalID

const (
	GlobalID_WlDisplay                        GlobalWlDisplayID                        = 1
	GlobalID_WlCompositor                     GlobalWlCompositorID                     = 0xff00000
	GlobalID_WlSubcompositor                  GlobalWlSubcompositorID                  = 0xff00001
	GlobalID_WlOutput                         GlobalWlOutputID                         = 0xff00002
	GlobalID_WlSeat                           GlobalWlSeatID                           = 0xff00003
	GlobalID_WlShm                            GlobalWlShmID                            = 0xff00004
	GlobalID_XdgWmBase                        GlobalXdgWmBaseID                        = 0xff00005
	GlobalID_WlDataDeviceManager              GlobalWlDataDeviceManagerID              = 0xff00006
	GlobalID_WlKeyboard                       GlobalWlKeyboardID                       = 0xff00007
	GlobalID_WlPointer                        GlobalWlPointerID                        = 0xff00008
	GlobalID_ZwpXwaylandKeyboardGrabManagerV1 GlobalZwpXwaylandKeyboardGrabManagerV1ID = 0xff00009
	GlobalID_XwaylandShellV1                  GlobalXwaylandShellV1ID                  = 0xff00011
	GlobalID_WlDataDevice                     GlobalWlDataDeviceID                     = 0xff00012
	GlobalID_WlTouch                          GlobalWlTouchID                          = 0xff00013
	GlobalID_ZxdgDecorationManagerV1          GlobalZxdgDecorationManagerV1ID          = 0xff00014
)

func GetGlobalWlDisplayBinds(cs ClientState) *map[ObjectID[WlDisplay]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_WlDisplay))
	m := v.(*map[ObjectID[WlDisplay]]Version)
	return m
}

func GetGlobalWlCompositorBinds(cs ClientState) *map[ObjectID[WlCompositor]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_WlCompositor))
	m := v.(*map[ObjectID[WlCompositor]]Version)
	return m
}

func GetGlobalWlSubcompositorBinds(cs ClientState) *map[ObjectID[WlSubcompositor]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_WlSubcompositor))
	m := v.(*map[ObjectID[WlSubcompositor]]Version)
	return m
}

func GetGlobalWlOutputBinds(cs ClientState) *map[ObjectID[WlOutput]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_WlOutput))
	m := v.(*map[ObjectID[WlOutput]]Version)
	return m
}

func GetGlobalWlSeatBinds(cs ClientState) *map[ObjectID[WlSeat]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_WlSeat))
	m := v.(*map[ObjectID[WlSeat]]Version)
	return m
}

func GetGlobalWlShmBinds(cs ClientState) *map[ObjectID[WlShm]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_WlShm))
	m := v.(*map[ObjectID[WlShm]]Version)
	return m
}

func GetGlobalXdgWmBaseBinds(cs ClientState) *map[ObjectID[XdgWmBase]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_XdgWmBase))
	m := v.(*map[ObjectID[XdgWmBase]]Version)
	return m
}

func GetGlobalWlDataDeviceManagerBinds(cs ClientState) *map[ObjectID[WlDataDeviceManager]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_WlDataDeviceManager))
	m := v.(*map[ObjectID[WlDataDeviceManager]]Version)
	return m
}

func GetGlobalWlKeyboardBinds(cs ClientState) *map[ObjectID[WlKeyboard]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_WlKeyboard))
	m := v.(*map[ObjectID[WlKeyboard]]Version)
	return m
}

func GetGlobalWlPointerBinds(cs ClientState) *map[ObjectID[WlPointer]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_WlPointer))
	m := v.(*map[ObjectID[WlPointer]]Version)
	return m
}

func GetGlobalZwpXwaylandKeyboardGrabManagerV1Binds(cs ClientState) *map[ObjectID[ZwpXwaylandKeyboardGrabManagerV1]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_ZwpXwaylandKeyboardGrabManagerV1))
	m := v.(*map[ObjectID[ZwpXwaylandKeyboardGrabManagerV1]]Version)
	return m
}

func GetGlobalXwaylandShellV1Binds(cs ClientState) *map[ObjectID[XwaylandShellV1]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_XwaylandShellV1))
	m := v.(*map[ObjectID[XwaylandShellV1]]Version)
	return m
}

func GetGlobalWlDataDeviceBinds(cs ClientState) *map[ObjectID[WlDataDevice]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_WlDataDevice))
	m := v.(*map[ObjectID[WlDataDevice]]Version)
	return m
}

func GetGlobalWlTouchBinds(cs ClientState) *map[ObjectID[WlTouch]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_WlTouch))
	m := v.(*map[ObjectID[WlTouch]]Version)
	return m
}

func GetGlobalZxdgDecorationManagerV1Binds(cs ClientState) *map[ObjectID[ZxdgDecorationManagerV1]]Version {

	v := cs.GetGlobalBinds(GlobalID(GlobalID_ZxdgDecorationManagerV1))
	m := v.(*map[ObjectID[ZxdgDecorationManagerV1]]Version)
	return m
}
