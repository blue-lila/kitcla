package forms

import (
	"github.com/blue-lila/kitcla/aria"
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/buttons"
	"github.com/blue-lila/kitcla/components/atoms/inputs"
	"github.com/blue-lila/kitcla/dat"
	"github.com/blue-lila/kitcla/goc"
)

type Form struct {
	Button         *buttons.Button
	SelectInputAlp *inputs.SelectInputAlp
	Component      *components.Component
}

type FormMod struct {
	FormData   interface{}
	Fields     []*FormField
	Target     string
	Redirects  []Redirect
	FormBag    dat.FormBag
	FieldSlot  FieldSlot
	SubmitText string
	OnSubmit   string
}

type FormField struct {
	Key      string
	Label    string
	Hidden   bool
	Value    any
	Data     any
	Indexes  []int
	Children []*FormField
	Template bool
}

type Redirect struct {
	Text string
	Url  string
}

type FieldSlot func(form interface{}, field *FormField) goc.HTML

func (this *Form) Mod() *FormMod {
	return &FormMod{}
}

func (this *FormMod) AddFormField(key string, text string) {
	this.Fields = append(this.Fields, &FormField{Key: key, Label: text})
}

func (this *FormMod) AddFormFieldWithVisibility(key string, text string, hidden bool, hiddenNew bool, hiddenEdit bool) {
	this.Fields = append(this.Fields, &FormField{Key: key, Label: text, Hidden: hidden})
}

func (this *FormMod) AddRedirect(text string, url string) {
	this.Redirects = append(this.Redirects, Redirect{
		Text: text,
		Url:  url,
	})
}

func (this *Form) EmptyBag() dat.FormBag {
	return dat.FormBag{}
}

func (this *Form) Form(formData interface{}, fields []*FormField, target string, bag dat.FormBag, fieldSlot FieldSlot, mod *FormMod) goc.HTML {
	mod = this.modDefaulting(mod)
	mod.FormData = formData
	mod.Fields = fields
	mod.Target = target
	mod.FormBag = bag
	mod.FieldSlot = fieldSlot
	return this.H(mod)
}

func (this *Form) submitButton(mod *FormMod) goc.HTML {
	submitText := mod.SubmitText
	if submitText == "" {
		submitText = "Submit"
	}
	return this.Component.Wrap("pt-4", this.Button.FormSubmit(submitText, nil))
}

func (this *Form) modDefaulting(mod *FormMod) *FormMod {
	if mod == nil {
		return &FormMod{}
	}
	return mod
}

func (this *Form) H(mod *FormMod) goc.HTML {
	var subs []goc.HTML

	if len(mod.FormBag.Invalidations) > 0 {
		subs = append(subs, this.formErrors(mod))
	}

	subs = append(subs, this.fields(mod)...)
	subs = append(subs, this.redirectField(mod))
	subs = append(subs, this.submitButtons(mod))

	attr := goc.Attr{
		"action":  mod.Target,
		"method":  "POST",
		"enctype": "multipart/form-data",
		"class":   "flex flex-col max-w-3xl space-y-4",
		"role":    aria.AriaRoleForm,
	}

	if mod.OnSubmit != "" {
		attr["onsubmit"] = mod.OnSubmit
	}

	h := this.Component.Cas("form", attr, subs...)
	// Forced to move the x-data=form() from the form el to a wrapped div because the alpinejs toolbar has a bug
	// when the x-data is on a form and a form field is named id
	return this.Component.Wa(goc.Attr{"x-data": "form()"}, h)
}

func (this *Form) fields(mod *FormMod) []goc.HTML {
	var subs []goc.HTML
	for _, formField := range mod.Fields {
		field := mod.FieldSlot(mod.FormData, formField)
		set := this.invalidationFor(mod, formField)
		if len(set) > 0 {
			errorMessages := this.fieldsErrorMessages(set)
			field = this.Component.Ds(field, errorMessages)
		}
		if formField.Hidden == true {
			field = this.Component.Dcs("hidden", field)
		}
		subs = append(subs, field)
	}
	return subs
}

func (this *Form) fieldsErrorMessages(invalidations []*dat.Invalidation) goc.HTML {
	var subs []goc.HTML
	for _, invalidation := range invalidations {
		subs = append(subs, this.fieldErrorMessage(invalidation))
	}
	return this.Component.Cas("ul", goc.Attr{"class": "text-red-700 pt-1 list-disc list-inside", "role": aria.AriaRoleAlert, aria.AriaLive: aria.AriaLivePolite}, subs...)
}

func (this *Form) fieldErrorMessage(invalidation *dat.Invalidation) goc.HTML {
	return this.Component.Cv("li", invalidation.Message)
}

func (this *Form) invalidationFor(mod *FormMod, forField *FormField) []*dat.Invalidation {
	var set []*dat.Invalidation
	for _, invalidation := range mod.FormBag.Invalidations {
		if invalidation.Field == forField.Key {
			set = append(set, invalidation)
		}
	}
	return set
}

func (this *Form) formErrors(mod *FormMod) goc.HTML {
	text := "There are errors in the form"
	return this.Component.Cav("ul", goc.Attr{"class": "px-2 py-2 text-red-700 bg-red-100 border-red-500", "role": aria.AriaRoleAlert, aria.AriaLive: aria.AriaLiveAssertive}, text)
}

func (this *Form) submitButtons(mod *FormMod) goc.HTML {
	return this.Component.Dcs("flex flex-row justify-between", this.submitButton(mod), this.redirectsButtons(mod))
}

func (this *Form) redirectsButtons(mod *FormMod) goc.HTML {
	if len(mod.Redirects) == 0 {
		return goc.HTML{}
	}
	var set []goc.HTML
	for _, redirect := range mod.Redirects {
		set = append(set, this.Button.SecondaryLink(redirect.Text, redirect.Url, nil))
	}
	selectMod := this.SelectInputAlp.Mod()
	for _, redirect := range mod.Redirects {
		selectMod.AddOption(redirect.Url, redirect.Text)
	}
	selectRedirects := this.SelectInputAlp.SelectInputAlp("redirect", selectMod)

	return this.Component.Das(goc.Attr{"class": "flex flex-row space-x-2 pt-4 items-center"},
		this.Component.Dcv("text-gray-600 italic", "then"),
		this.Component.Ds(selectRedirects))
}

func (this *Form) redirectField(mod *FormMod) goc.HTML {
	if len(mod.Redirects) == 0 {
		return goc.HTML{}
	}
	return this.Component.Cav("input", goc.Attr{"type": "hidden", "name": "_redirect", "x-model": "redirect", "value": mod.Redirects[0].Url})
}
