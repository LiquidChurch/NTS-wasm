package modal

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Handler struct {
	app.Compo

	targetModal string
}

func (h *Handler) OnMount(ctx app.Context) {
	ctx.Handle("modal", h.openModal)
}

func (h *Handler) openModal(ctx app.Context, a app.Action) {
	h.targetModal = a.Value.(string)
}

func (h *Handler) Render() app.UI {
	return app.Div().
		Body(
			app.If(
				h.targetModal == "checkinModal",
				&checkinModal{},
			),
		)

}
