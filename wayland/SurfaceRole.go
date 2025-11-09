package wayland

import "github.com/mmulet/term.everything/wayland/protocols"

/**
 * Surface roles can be thought of as type.
 * [source ](https://wayland.app/protocols/wayland#wl_surface)
 * Things you may do:
 * 1. if unset you may assign it a role
 * 2. if assigned a role you may *not* change it to another role
 * 3. if assigned a role you may *may* assign it the *same* role
 * 4. The role may destroyed, allowing you assign it to the role again (maybe with different data), but not a different role.
 * 5. If the surface is destroyed before the role is destroyed, that is an error.
 */
type SurfaceRole interface {
	surface_role()
	HasData() bool
	ClearData()
}

type SurfaceRoleXdgPopup struct {
	Data *protocols.ObjectID[protocols.XdgPopup]
}

func (r *SurfaceRoleXdgPopup) surface_role() {}
func (r *SurfaceRoleXdgPopup) HasData() bool {
	return r.Data != nil
}

func (r *SurfaceRoleXdgPopup) ClearData() {
	r.Data = nil
}

type CursorHotspot struct {
	X, Y int32
}

type SurfaceRoleCursorData struct {
	Hotspot CursorHotspot
}

type SurfaceRoleCursor struct {
	Data *SurfaceRoleCursorData
}

func (r *SurfaceRoleCursor) surface_role() {}
func (r *SurfaceRoleCursor) HasData() bool {
	return r.Data != nil
}
func (r *SurfaceRoleCursor) ClearData() {
	r.Data = nil
}

type ToplevelPendingState struct {
	MinSize *Size
	MaxSize *Size
}

type SurfaceRoleXdgToplevel struct {
	Data *protocols.ObjectID[protocols.XdgToplevel]
}

func (r *SurfaceRoleXdgToplevel) surface_role() {}
func (r *SurfaceRoleXdgToplevel) HasData() bool {
	return r.Data != nil
}

func (r *SurfaceRoleXdgToplevel) ClearData() {
	r.Data = nil
}

type SurfaceRoleWaylandSurfaceData struct {
	Serial *XWaylandSurfaceV1Serial
}

type SurfaceRoleXWaylandSurface struct {
	Data *SurfaceRoleWaylandSurfaceData
}

func (r *SurfaceRoleXWaylandSurface) surface_role() {}
func (r *SurfaceRoleXWaylandSurface) HasData() bool {
	return r.Data != nil
}

func (r *SurfaceRoleXWaylandSurface) ClearData() {
	r.Data = nil
}

type SurfaceRoleSubSurface struct {
	Data *protocols.ObjectID[protocols.WlSubsurface]
}

func (r *SurfaceRoleSubSurface) surface_role() {}
func (r *SurfaceRoleSubSurface) HasData() bool {
	return r.Data != nil
}
func (r *SurfaceRoleSubSurface) ClearData() {
	r.Data = nil
}
