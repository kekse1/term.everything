package wayland

import (
	"github.com/mmulet/term.everything/wayland/pointerslices"
	"github.com/mmulet/term.everything/wayland/protocols"
)

type PendingBufferUpdates struct {
	Surface protocols.ObjectID[protocols.WlSurface]
	Buffer  *protocols.ObjectID[protocols.WlBuffer]
	ZIndex  int
}

func ApplyWlSurfaceDoubleBufferedState(
	s protocols.ClientState,
	surfaceObjectID protocols.ObjectID[protocols.WlSurface],
	syncSetByParent bool,
	accumulator []PendingBufferUpdates,
	zIndex int,
) []PendingBufferUpdates {
	/**
	 * Could be a child surface
	 */
	surface := GetWlSurfaceObject(s, surfaceObjectID)
	if surface == nil {
		return accumulator
	}

	update := &surface.PendingUpdate
	if update.Buffer != nil {
		accumulator = append(accumulator, PendingBufferUpdates{
			Surface: surfaceObjectID,
			Buffer:  update.Buffer,
			ZIndex:  zIndex,
		})
	}

	if update.BufferScale != nil {
		surface.BufferScale = *update.BufferScale
	}

	if update.BufferTransform != nil {
		surface.BufferTransform = *update.BufferTransform
	}

	if update.Damage != nil || update.DamageBuffer != nil {
		surface.Damaged = true
	} else {
		surface.Damaged = false
	}

	// offset: add to current offset (doc semantics)
	if update.Offset != nil {
		/**
		 * @TODO Docs say:
		 * The x and y arguments specify the location of the new pending
		 * buffer's upper left corner,
		 * relative to the current buffer's upper left corner,
		 * in surface-local coordinates.
		 * In other words, the x and y,
		 * combined with the new surface size define in
		 * which directions the surface's size changes.
		 *
		 * So I think this means I should add the offset to the current offset
		 * of the surface, not just set it to the offset.
		 */
		surface.Offset.X += update.Offset.X
		surface.Offset.Y += update.Offset.Y

		/**
		 * From the docs:
		 * On wl_surface.offset requests to the pointer surface, hotspot_x and hotspot_y are decremented by the x and y parameters passed to the request. The offset must be applied by wl_surface.commit as usual.
		 */
		// if (surface.role?.type === "cursor" && surface.role.data) {
		//   surface.role.data.hotspot.x -= update.offset.x;
		//   surface.role.data.hotspot.y -= update.offset.y;
		// }
	}

	if update.InputRegion != nil {
		if surface.InputRegion != nil && !AreSame(surface.InputRegion, update.InputRegion) {
			RemoveObject(s, *surface.InputRegion)
		}
		surface.InputRegion = update.InputRegion
	}

	if update.OpaqueRegion != nil {
		if surface.OpaqueRegion != nil && !AreSame(surface.OpaqueRegion, update.OpaqueRegion) {
			RemoveObject(s, *surface.OpaqueRegion)
		}
		surface.OpaqueRegion = update.OpaqueRegion
	}

	if update.AddSubSurface != nil {
		for _, subID := range update.AddSubSurface {
			surface.ChildrenInDrawOrder = append(
				[]*protocols.ObjectID[protocols.WlSurface]{&subID},
				surface.ChildrenInDrawOrder...,
			)
		}
	}

	if update.SetChildPosition != nil {
		for _, childPosition := range update.SetChildPosition {
			if !pointerslices.Contains(surface.ChildrenInDrawOrder, childPosition.Child) {
				continue
			}
			childSurface := GetWlSurfaceObject(s, childPosition.Child)
			if childSurface == nil {
				continue
			}
			role, ok := childSurface.Role.(*SurfaceRoleSubSurface)
			if !ok || role.Data == nil {
				continue
			}
			sub := GetWlSubsurfaceObject(s, *role.Data)
			if sub == nil {
				continue
			}
			sub.Position = Point{
				X: childPosition.X,
				Y: childPosition.Y,
			}
		}
	}

	if update.ZOrderSubsurfaces != nil {
		for _, zUpdate := range update.ZOrderSubsurfaces {
			index_of_child := pointerslices.Index(surface.ChildrenInDrawOrder, zUpdate.ChildToMove)
			if index_of_child == -1 {
				continue
			}
			index_of_relative_to := pointerslices.IndexOfItemOrNil(surface.ChildrenInDrawOrder, zUpdate.RelativeTo)
			if index_of_relative_to == -1 {
				continue
			}

			/**
			* Remove the child from the list
			* then reinsert it at the correct index
			* either above or below the relative_to child
			* Since it is drawn in order, above means it will
			* be added to the array after the relative_to child
			* and below means it will be added before the relative_to child
			 */
			surface.ChildrenInDrawOrder = pointerslices.Delete(surface.ChildrenInDrawOrder, index_of_child, index_of_child+1)

			var offset int
			if zUpdate.Type == ZOrderTypeAbove {
				offset = 1
			} else {
				offset = 0
			}
			surface.ChildrenInDrawOrder = pointerslices.Insert(surface.ChildrenInDrawOrder, index_of_relative_to+offset, &zUpdate.ChildToMove)
		}
	}

	if update.XdgSurfaceWindowGeometry != nil {
		if xdg_surface_state_id := surface.XdgSurfaceState; xdg_surface_state_id != nil {
			if xdg_surface_state := GetXdgSurfaceObject(s, *xdg_surface_state_id); xdg_surface_state != nil {
				xdg_surface_state.WindowGeometry = *update.XdgSurfaceWindowGeometry
			}
		}
	}

	if role, ok := surface.Role.(*SurfaceRoleXdgToplevel); ok && role.Data != nil {
		top := GetXdgToplevelObject(s, *role.Data)
		if top != nil && top.PendingState != nil {

			if top.PendingState.MaxSize != nil {
				top.MaxSize = top.PendingState.MaxSize
			}

			if top.PendingState.MinSize != nil {
				top.MinSize = top.PendingState.MinSize
			}
			top.PendingState = nil
		}
	}

	// if (
	//   surface.has_role_data_of_type("xdg_toplevel") &&
	//   surface.role.data.pending_state
	// ) {
	//   if (surface.role.data.pending_state.max_size) {
	//     surface.role.data.max_size = surface.role.data.pending_state.max_size;
	//   }
	//   if (surface.role.data.pending_state.min_size) {
	//     surface.role.data.min_size = surface.role.data.pending_state.min_size;
	//   }
	//   delete surface.role.data.pending_state;
	// }

	if update.XwaylandSurfarfaceV1Serial != nil {
		if role, ok := surface.Role.(*SurfaceRoleXWaylandSurface); ok {
			if role.Data == nil {
				role.Data = &SurfaceRoleWaylandSurfaceData{}
			}
			(*role.Data).Serial = update.XwaylandSurfarfaceV1Serial
		}
	}

	surface.ResetPendingUpdate()

	for _, childSurfaceObjectID := range surface.ChildrenInDrawOrder {

		if childSurfaceObjectID == nil {
			continue
		}

		if syncSetByParent {
			accumulator = ApplyWlSurfaceDoubleBufferedState(
				s,
				*childSurfaceObjectID,
				syncSetByParent,
				accumulator,
				zIndex+1,
			)
			continue
		}

		childSurface := GetWlSurfaceObject(s, *childSurfaceObjectID)
		if childSurface == nil {
			continue
		}
		role, ok := childSurface.Role.(*SurfaceRoleSubSurface)
		if !ok || role.Data == nil {
			continue
		}
		sub := GetWlSubsurfaceObject(s, *role.Data)
		if sub == nil {
			continue
		}
		// if (
		//   child_surface.role.type !== "sub_surface" ||
		//   child_surface.role.data === null
		// ) {
		//   continue;
		// }
		// if (!child_surface.role.data.sync) {
		if !sub.Sync {
			/**
			 * The child is not set to sync with the parent
			 * so do not apply state changes now
			 */
			continue
		}
		accumulator = ApplyWlSurfaceDoubleBufferedState(
			s,
			*childSurfaceObjectID,
			true,
			accumulator,
			zIndex+1,
		)
	}
	return accumulator
}
