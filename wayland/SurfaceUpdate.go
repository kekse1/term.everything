package wayland

import "github.com/mmulet/term.everything/wayland/protocols"

type SurfaceUpdate struct {
	Offset *Point

	Damage []Rect

	DamageBuffer []Rect

	BufferScale *int32

	BufferTransform *protocols.WlOutputTransform_enum

	InputRegion *protocols.ObjectID[protocols.WlRegion]

	OpaqueRegion *protocols.ObjectID[protocols.WlRegion]

	Buffer *protocols.ObjectID[protocols.WlBuffer]

	/**
	 * You should unshift when adding to
	 * this array so that the objects will
	 * be added in the correct order. (ie
	 * the ones added last will be on top)
	 */
	AddSubSurface []protocols.ObjectID[protocols.WlSurface]

	XdgSurfaceWindowGeometry *XdgWindowGeometry

	/**
	 * set_child_position and z_oder_subsurfaces
	 * take place whenever the parent surface is committed,
	 * thus they are part of the SurfaceUpdate of the parent
	 */
	SetChildPosition []ChildPosition

	/**
	 * null means above or below the parent.
	 */
	ZOrderSubsurfaces []ZOrderSubsurface

	XwaylandSurfarfaceV1Serial *XWaylandSurfaceV1Serial
}

type Point struct {
	X int32
	Y int32
}

type Rect struct {
	X      int32
	Y      int32
	Width  int32
	Height int32
}

// ChildPosition mirrors { child: Object_ID<wl_surface>; x: number; y: number }
type ChildPosition struct {
	Child protocols.ObjectID[protocols.WlSurface]
	X     int32
	Y     int32
}

type ZOrderSubsurface struct {
	Type        ZOrder
	ChildToMove protocols.ObjectID[protocols.WlSurface]
	/**
	 * nil means above or below the parent.
	 */
	RelativeTo *protocols.ObjectID[protocols.WlSurface]
}

type ZOrder int

const (
	ZOrderTypeAbove ZOrder = iota
	ZOrderTypeBelow ZOrder = iota + 1
)

type XWaylandSurfaceV1Serial struct {
	Low uint32
	Hi  uint32
}
