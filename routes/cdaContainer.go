package routes

import (
	h "../handlers"
	m "../models"
)

//CdaContainerRoutes cdas routes.
var CdaContainerRoutes = m.Routes{
	m.Route{
		"CdaContainerIndex",
		"GET",
		"/cda",
		h.CdaContainerIndex,
	},
	m.Route{
		"CdaContainerSave",
		"POST",
		"/cda",
		h.CdaContainerSave,
	},
	m.Route{
		"CdaContainerFindByID",
		"GET",
		"/cda/{cdaID}",
		h.CdaContainerFindByID,
	},
	m.Route{
		"CdaContainerUpdate",
		"PUT",
		"/cda/{cdaID}",
		h.CdaContainerUpdate,
	},
	m.Route{
		"CdaContainerDelete",
		"DELETE",
		"/cda/{cdaID}",
		h.CdaContainerDelete,
	},
}
