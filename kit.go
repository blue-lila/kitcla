package kitcla

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/buttons"
	"github.com/blue-lila/kitcla/components/atoms/buttons_groups"
	"github.com/blue-lila/kitcla/components/atoms/card_bodies"
	"github.com/blue-lila/kitcla/components/atoms/card_headers"
	"github.com/blue-lila/kitcla/components/atoms/card_wrappers"
	"github.com/blue-lila/kitcla/components/atoms/cells"
	"github.com/blue-lila/kitcla/components/atoms/dropdowns"
	"github.com/blue-lila/kitcla/components/atoms/fields"
	"github.com/blue-lila/kitcla/components/atoms/headers"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/components/atoms/inputs"
	"github.com/blue-lila/kitcla/components/atoms/links"
	"github.com/blue-lila/kitcla/components/atoms/paginations"
	"github.com/blue-lila/kitcla/components/atoms/placeholders"
	"github.com/blue-lila/kitcla/components/atoms/resources"
	"github.com/blue-lila/kitcla/components/atoms/shows"
	"github.com/blue-lila/kitcla/components/atoms/tabs"
	"github.com/blue-lila/kitcla/components/molecules/alerts"
	"github.com/blue-lila/kitcla/components/molecules/messages"
	nabvars2 "github.com/blue-lila/kitcla/components/molecules/navbars"
	"github.com/blue-lila/kitcla/components/molecules/popovers"
	"github.com/blue-lila/kitcla/components/molecules/steppers"
	"github.com/blue-lila/kitcla/components/molecules/trees"
	"github.com/blue-lila/kitcla/components/organisms/cards"
	"github.com/blue-lila/kitcla/components/organisms/duo_cards"
	"github.com/blue-lila/kitcla/components/organisms/forms"
	"github.com/blue-lila/kitcla/components/organisms/grids"
	"github.com/blue-lila/kitcla/components/organisms/key_point"
	"github.com/blue-lila/kitcla/components/organisms/logs"
	"github.com/blue-lila/kitcla/components/organisms/menus"
	"github.com/blue-lila/kitcla/components/organisms/modals"
	"github.com/blue-lila/kitcla/components/organisms/navbars"
	"github.com/blue-lila/kitcla/components/organisms/navmenus"
	"github.com/blue-lila/kitcla/components/organisms/sidebars"
	"github.com/blue-lila/kitcla/components/organisms/tables"
)

type Kit struct {
	Component *components.Component
	Atoms     struct {
		Buttons struct {
			Button    *buttons.Button
			ButtonAlp *buttons.ButtonAlp
		}
		ButtonsGroups struct {
			ButtonsGroupAlp *buttons_groups.ButtonsGroupAlp
		}
		Headers struct {
			Header *headers.Header
		}
		Fields struct {
			Field *fields.Field
		}
		Dropdowns struct {
			Dropdown *dropdowns.Dropdown
		}
		Inputs struct {
			TextInput           *inputs.TextInput
			TextInputAlp        *inputs.TextInputAlp
			RichTextInput       *inputs.RichTextInput
			TextAreaInput       *inputs.TextAreaInput
			TextAreaInputAlp    *inputs.TextAreaInputAlp
			IntegerInput        *inputs.IntegerInput
			IntegerInputAlp     *inputs.IntegerInputAlp
			RadioInputAlp       *inputs.RadioInputAlp
			DecimalInput        *inputs.DecimalInput
			FileInput           *inputs.FileInput
			BooleanInput        *inputs.BooleanInput
			BooleanInputAlp     *inputs.BooleanInputAlp
			SwitchInput         *inputs.SwitchInput
			SwitchInputAlp      *inputs.SwitchInputAlp
			SelectInput         *inputs.SelectInput
			SelectInputAlp      *inputs.SelectInputAlp
			JsonInput           *inputs.JsonInput
			HiddenInput         *inputs.HiddenInput
			DatetimeInput       *inputs.DatetimeInput
			AdvancedSelectInput *inputs.AdvancedSelectInput
			GridIdInput         *inputs.GridIdInput
		}
		Shows struct {
			TextShow     *shows.TextShow
			RichTextShow *shows.RichTextShow
		}
		Cells struct {
			TextCell     *cells.TextCell
			PillCell     *cells.PillCell
			IntegerCell  *cells.IntegerCell
			DecimalCell  *cells.DecimalCell
			BooleanCell  *cells.BooleanCell
			TimeCell     *cells.TimeCell
			JsonCell     *cells.JsonCell
			LongTextCell *cells.LongTextCell
			RichTextCell *cells.RichTextCell
			RelationCell *cells.RelationCell
		}
		Placeholders struct {
			Placeholder *placeholders.Placeholder
		}
		Icons struct {
			Icon *icons.Icon
		}
		Links struct {
			Link *links.Link
		}
		Resources struct {
			Resource *resources.Resource
		}
		CardHeader struct {
			CardHeader *card_headers.CardHeader
		}
		CardBodies struct {
			DuoCardBody *card_bodies.DuoCardBody
		}
		CardWrapper struct {
			CardWrapper *card_wrappers.CardWrapper
		}
		Tabs struct {
			Tab *tabs.Tab
		}
		Paginations struct {
			Pagination *paginations.Pagination
		}
	}
	Molecules struct {
		Alerts struct {
			Alert *alerts.Alert
		}
		Messages struct {
			Message *messages.Message
		}
		Trees struct {
			Tree *trees.Tree
		}
		Popovers struct {
			Popover *popovers.Popover
		}
		Stepper struct {
			Stepper *steppers.Stepper
		}
		Navbars struct {
			Navbar *nabvars2.Navbar
		}
	}
	Organisms struct {
		Tables struct {
			Table         *tables.Table
			KeyValueTable *tables.KeyValueTable
		}
		Cards struct {
			Card *cards.Card
		}
		Modals struct {
			Modal *modals.Modal
		}
		Forms struct {
			Form       *forms.Form
			DeleteForm *forms.DeleteForm
		}
		Sidebars struct {
			Sidebar *sidebars.Sidebar
		}
		Navbars struct {
			Navbar     *navbars.Navbar
			DualNavbar *navbars.DualNavbar
		}
		Grids struct {
			IconCardGrid *grids.IconCardGrid
			KeyValueGrid *grids.KeyValueGrid
		}
		DuoCard struct {
			PeekDuoCard *duo_cards.PeekDuoCard
		}
		Menus struct {
			Menu *menus.Menu
		}
		Navmenus struct {
			Navmenu *navmenus.Navmenu
		}
		KeyPoint struct {
			KeyPoint *key_point.KeyPoint
		}
		Logs struct {
			Logs *logs.Logs
		}
	}
}

func New() *Kit {
	kit := &Kit{
		Component: &components.Component{},
	}

	kit.Atoms.Headers.Header = &headers.Header{Component: kit.Component}
	kit.Atoms.Icons.Icon = &icons.Icon{Component: kit.Component}
	kit.Atoms.Resources.Resource = &resources.Resource{Component: kit.Component}
	kit.Atoms.Links.Link = &links.Link{
		Component: kit.Component,
		Icon:      kit.Atoms.Icons.Icon,
	}
	kit.Atoms.Buttons.Button = &buttons.Button{
		Component: kit.Component,
		Icon:      kit.Atoms.Icons.Icon,
	}
	kit.Atoms.Buttons.ButtonAlp = &buttons.ButtonAlp{
		Component: kit.Component,
		Icon:      kit.Atoms.Icons.Icon,
		Button:    kit.Atoms.Buttons.Button,
	}
	kit.Atoms.ButtonsGroups.ButtonsGroupAlp = &buttons_groups.ButtonsGroupAlp{Component: kit.Component}
	kit.Atoms.Fields.Field = &fields.Field{Component: kit.Component}
	kit.Atoms.Dropdowns.Dropdown = &dropdowns.Dropdown{
		Component: kit.Component,
		Icon:      kit.Atoms.Icons.Icon,
	}
	kit.Atoms.Inputs.TextInput = &inputs.TextInput{Component: kit.Component}
	kit.Atoms.Inputs.TextInputAlp = &inputs.TextInputAlp{Component: kit.Component}
	kit.Atoms.Inputs.IntegerInput = &inputs.IntegerInput{Component: kit.Component}
	kit.Atoms.Inputs.IntegerInputAlp = &inputs.IntegerInputAlp{
		Component: kit.Component,
		Icon:      kit.Atoms.Icons.Icon,
	}
	kit.Atoms.Inputs.RadioInputAlp = &inputs.RadioInputAlp{Component: kit.Component}
	kit.Atoms.Inputs.DecimalInput = &inputs.DecimalInput{Component: kit.Component}
	kit.Atoms.Inputs.FileInput = &inputs.FileInput{Component: kit.Component}
	kit.Atoms.Inputs.BooleanInput = &inputs.BooleanInput{Component: kit.Component}
	kit.Atoms.Inputs.BooleanInputAlp = &inputs.BooleanInputAlp{Component: kit.Component}
	kit.Atoms.Inputs.SwitchInput = &inputs.SwitchInput{Component: kit.Component}
	kit.Atoms.Inputs.SwitchInputAlp = &inputs.SwitchInputAlp{Component: kit.Component}
	kit.Atoms.Inputs.SelectInput = &inputs.SelectInput{Component: kit.Component}
	kit.Atoms.Inputs.SelectInputAlp = &inputs.SelectInputAlp{Component: kit.Component}
	kit.Atoms.Inputs.JsonInput = &inputs.JsonInput{Component: kit.Component}
	kit.Atoms.Inputs.HiddenInput = &inputs.HiddenInput{Component: kit.Component}
	kit.Atoms.Inputs.DatetimeInput = &inputs.DatetimeInput{Component: kit.Component}
	kit.Atoms.Inputs.AdvancedSelectInput = &inputs.AdvancedSelectInput{Component: kit.Component}
	kit.Atoms.Inputs.GridIdInput = &inputs.GridIdInput{Component: kit.Component}
	kit.Atoms.Inputs.RichTextInput = &inputs.RichTextInput{
		Component: kit.Component,
		Resource:  kit.Atoms.Resources.Resource,
	}
	kit.Atoms.Inputs.TextAreaInput = &inputs.TextAreaInput{Component: kit.Component}
	kit.Atoms.Inputs.TextAreaInputAlp = &inputs.TextAreaInputAlp{Component: kit.Component}
	kit.Atoms.Placeholders.Placeholder = &placeholders.Placeholder{Component: kit.Component}
	kit.Atoms.Shows.TextShow = &shows.TextShow{Component: kit.Component}
	kit.Atoms.Shows.RichTextShow = &shows.RichTextShow{Component: kit.Component}
	kit.Atoms.Cells.TextCell = &cells.TextCell{Component: kit.Component}
	kit.Atoms.Cells.PillCell = &cells.PillCell{Component: kit.Component}
	kit.Atoms.Cells.RichTextCell = &cells.RichTextCell{Component: kit.Component}
	kit.Atoms.Cells.JsonCell = &cells.JsonCell{Component: kit.Component}
	kit.Atoms.Cells.LongTextCell = &cells.LongTextCell{Component: kit.Component}
	kit.Atoms.Cells.IntegerCell = &cells.IntegerCell{Component: kit.Component}
	kit.Atoms.Cells.DecimalCell = &cells.DecimalCell{Component: kit.Component}
	kit.Atoms.Cells.BooleanCell = &cells.BooleanCell{Component: kit.Component, Icon: kit.Atoms.Icons.Icon}
	kit.Atoms.Cells.TimeCell = &cells.TimeCell{Component: kit.Component}
	kit.Atoms.Cells.RelationCell = &cells.RelationCell{Component: kit.Component}
	kit.Atoms.CardBodies.DuoCardBody = &card_bodies.DuoCardBody{Component: kit.Component}
	kit.Atoms.CardHeader.CardHeader = &card_headers.CardHeader{Component: kit.Component}
	kit.Atoms.CardWrapper.CardWrapper = &card_wrappers.CardWrapper{Component: kit.Component}
	kit.Atoms.Tabs.Tab = &tabs.Tab{Component: kit.Component}
	kit.Atoms.Paginations.Pagination = &paginations.Pagination{Component: kit.Component}
	kit.Molecules.Alerts.Alert = &alerts.Alert{
		Component: kit.Component,
		Icon:      kit.Atoms.Icons.Icon,
	}
	kit.Molecules.Messages.Message = &messages.Message{Component: kit.Component}
	kit.Molecules.Trees.Tree = &trees.Tree{Component: kit.Component}
	kit.Molecules.Navbars.Navbar = &nabvars2.Navbar{Component: kit.Component}
	kit.Molecules.Stepper.Stepper = &steppers.Stepper{Component: kit.Component}
	kit.Molecules.Popovers.Popover = &popovers.Popover{
		Component: kit.Component,
		ButtonAlp: kit.Atoms.Buttons.ButtonAlp,
	}
	kit.Organisms.Tables.Table = &tables.Table{
		Component:      kit.Component,
		Button:         kit.Atoms.Buttons.Button,
		TextInput:      kit.Atoms.Inputs.TextInput,
		BooleanInput:   kit.Atoms.Inputs.BooleanInput,
		SelectInput:    kit.Atoms.Inputs.SelectInput,
		SelectInputAlp: kit.Atoms.Inputs.SelectInputAlp,
		IntegerInput:   kit.Atoms.Inputs.IntegerInput,
		Dropdown:       kit.Atoms.Dropdowns.Dropdown,
		Header:         kit.Atoms.Headers.Header,
		Pagination:     kit.Atoms.Paginations.Pagination,
	}
	kit.Organisms.Tables.KeyValueTable = &tables.KeyValueTable{Component: kit.Component}
	kit.Organisms.Cards.Card = &cards.Card{Component: kit.Component}
	kit.Organisms.Modals.Modal = &modals.Modal{
		Component: kit.Component,
		Button:    kit.Atoms.Buttons.Button,
	}
	kit.Organisms.Grids.IconCardGrid = &grids.IconCardGrid{
		Component: kit.Component,
		Icon:      kit.Atoms.Icons.Icon,
		Link:      kit.Atoms.Links.Link,
	}
	kit.Organisms.Grids.KeyValueGrid = &grids.KeyValueGrid{Component: kit.Component}
	kit.Organisms.Forms.Form = &forms.Form{
		Component:      kit.Component,
		Button:         kit.Atoms.Buttons.Button,
		SelectInputAlp: kit.Atoms.Inputs.SelectInputAlp,
	}
	kit.Organisms.Forms.DeleteForm = &forms.DeleteForm{
		Component: kit.Component,
		Button:    kit.Atoms.Buttons.Button,
		Link:      kit.Atoms.Links.Link,
	}
	kit.Organisms.Sidebars.Sidebar = &sidebars.Sidebar{Component: kit.Component}
	kit.Organisms.KeyPoint.KeyPoint = &key_point.KeyPoint{
		Component: kit.Component,
		Icon:      kit.Atoms.Icons.Icon,
	}
	kit.Organisms.DuoCard.PeekDuoCard = &duo_cards.PeekDuoCard{
		Component:   kit.Component,
		DuoCardBody: kit.Atoms.CardBodies.DuoCardBody,
		Table:       kit.Organisms.Tables.Table,
	}
	kit.Organisms.Navbars.Navbar = &navbars.Navbar{
		Component: kit.Component,
		Icon:      kit.Atoms.Icons.Icon,
		Link:      kit.Atoms.Links.Link,
	}
	kit.Organisms.Navbars.DualNavbar = &navbars.DualNavbar{
		Component: kit.Component,
		Icon:      kit.Atoms.Icons.Icon,
	}
	kit.Organisms.Menus.Menu = &menus.Menu{
		Component: kit.Component,
		Navbar:    kit.Molecules.Navbars.Navbar,
	}
	kit.Organisms.Logs.Logs = &logs.Logs{}
	kit.Organisms.Navmenus.Navmenu = &navmenus.Navmenu{
		Component: kit.Component,
		Icon:      kit.Atoms.Icons.Icon,
	}
	return kit
}
