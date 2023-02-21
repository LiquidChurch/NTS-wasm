package element

import (
	"log"
	"nts/component"
	"nts/rockapi"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func Person() *person {
	return &person{}
}

type person struct {
	app.Compo
	person rockapi.PersonDetail
	fullName string
	firstName string
}

func (p *person) OnMount(ctx app.Context) {
	ctx.Handle("person", p.setActive)
}

func (p *person) setActive(ctx app.Context, a app.Action) {
	p.person = interface{}(a.Value).(rockapi.PersonDetail)
	p.fullName = p.person.NickName + " " + p.person.LastName
}

func (p *person) Render() app.UI {
	var personSet bool
	if p.fullName == "" {
		log.Println("Person Info Not Loaded")
		personSet = false
		return app.Div().Body()
	} else {
		log.Println("Person Info Loaded")
		personSet = true
	}

	return app.Div().
		Class("").
		Body(
			app.If(personSet == true,
				component.Header().Level(3).Text("Name"),
				component.Header().Level(1).Text(p.fullName)),
		)
}