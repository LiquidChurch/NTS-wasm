package modal

import (
	"log"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type modal struct {
	app.Compo

	Title 				string
	DisableFocus 		bool
	Body 				[]app.UI
	removeEventListener func()
}

func (m *modal) OnMount(ctx app.Context) {
	log.Println("Modal OnMount")
	m.removeEventListener = app.Window().AddEventListener("keyup", func(ctx app.Context, e app.Event) {
		if e.Get("key").String() == "Escape" {
			ctx.NewActionWithValue("modal", "")
		}
	})
}

func (m *modal) OnDismount() {
	log.Println("Modal On Dismount")
	if m.removeEventListener != nil {
		m.removeEventListener()
	}
	log.Println("trigger update")
	m.Update()
}

func (m *modal) onClose(ctx app.Context, e app.Event) {
	log.Println("Modal On Close")
	if e.Get("currentTarget").Call("isSameNode", e.Get("target")).Bool() {
		ctx.NewActionWithValue("modal", "")
	}
}

func (m *modal) closeModal(ctx app.Context, e app.Event) {
	log.Println("Close Modal")
	ctx.NewActionWithValue("modal", "")
}

func (m *modal) Render() app.UI {
	return app.Div().
		Class("bg-gray-300 bg-opacity-80 fixed inset-0 z-40 w-full h-full").
		Body(
			app.Div().
				Class("flex h-full justify-center items-center").
				OnClick(func(ctx app.Context, e app.Event) {
					// Close if we clicked outside the modal
					if e.Get("currentTarget").Call("isSameNode", e.Get("target")).Bool() {
						m.closeModal(ctx, e)
					}
				}).
				Body(
					&Autofocused{
						Disable: m.DisableFocus,
						Component: app.Div().
							TabIndex(-1).
							Class("z-50 overflow-x-hidden overflow-y-auto w-[90vw] sm:w-[80vw] md:w-[70vw] lg:w-[60vw]").
							Body(
								app.Div().
									Class("bg-white rounded-t border-2 border-blue-300 shadow min-h-[60vh] sm:min-h-[40vh]").
									Body(
										app.Div().
											Class("flex items-start justify-between p-5 border-b rounded-t").
											Body(
												app.H3().
													OnClick(m.closeModal).
													Class("text-gray-800 text-xl lg:text-2xl font-semibold").
													Text(m.Title),
												app.Button().
													Class("h-6 w-6 z-60 rounded hover:opacity-75 hover:bg-bluish focus:outline focus:outline-focus").
													TabIndex(0).
													OnClick(m.closeModal).
													Body(
														app.Img().
															Class("color-gray").
															Src("/web/asset/x-mark.svg"),
													),
											),
										app.Div().
											Class("p-4 flex flex-col").
											Body(
												app.Range(m.Body).Slice(func(j int) app.UI {
													return m.Body[j]
													}),
												),
									),
							),
					},
				),
		)
}
