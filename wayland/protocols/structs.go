package protocols

type ObjectID[T any] uint32

type AnyObjectID = ObjectID[any]

type WaylandObject[T any] interface {
	Delegate() T
	OnRequest(s FileDescriptorClaimClientState, message DecodeState)
}

// Don't use these functions directly; use the ones in wayland/types.go
type ClientState interface {
	RemoveObject(AnyObjectID)
	RemoveGlobalBind(GlobalID, AnyObjectID)
	AddObject(AnyObjectID, any)
	SetCompositorVersion(uint32)
	GetCompositorVersion() uint32
	GetObject(AnyObjectID) any

	RegisterRoleToSurface(AnyObjectID, ObjectID[WlSurface])
	UnregisterRoleToSurface(AnyObjectID)
	Send(OutgoingEvent)
	SendError(AnyObjectID, uint32, string)

	DrawableSurfaces() map[ObjectID[WlSurface]]bool
	TopLevelSurfaces() map[ObjectID[XdgToplevel]]bool
	AddFrameDrawRequest(ObjectID[WlCallback])

	GetSurfaceIDFromRole(AnyObjectID) *AnyObjectID

	GetSurfaceFromRole(AnyObjectID) any

	GetGlobalBinds(GlobalID) any
}

type OutgoingEvent struct {
	ObjectID       AnyObjectID
	Opcode         uint16
	Data           []byte
	FileDescriptor *FileDescriptor
}

type FileDescriptorClaimClientState interface {
	ClientState
	ClaimFileDescriptor() *FileDescriptor
}
type FileDescriptor int

type Sender interface {
	Send(OutgoingEvent)
}

type DecodeStateType int

type DecodeState struct {
	Phase    DecodeStateType
	I        uint // bit offset within current field (0,8,16,24)
	ObjectID AnyObjectID
	Opcode   uint16
	Size     uint16
	Data     []byte
}

type Fixed = float64
