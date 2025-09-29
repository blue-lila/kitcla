package aria

// ARIA roles constants
const (
	// Landmark roles
	AriaRoleBanner        = "banner"
	AriaRoleNavigation    = "navigation"
	AriaRoleMain          = "main"
	AriaRoleContentinfo   = "contentinfo"
	AriaRoleSearch        = "search"
	AriaRoleComplementary = "complementary"

	// Document structure roles
	AriaRoleApplication  = "application"
	AriaRoleArticle      = "article"
	AriaRoleDocument     = "document"
	AriaRoleHeading      = "heading"
	AriaRoleImg          = "img"
	AriaRoleList         = "list"
	AriaRoleListitem     = "listitem"
	AriaRoleMath         = "math"
	AriaRoleNote         = "note"
	AriaRolePresentation = "presentation"
	AriaRoleRegion       = "region"
	AriaRoleSeparator    = "separator"
	AriaRoleToolbar      = "toolbar"

	// Widget roles
	AriaRoleButton     = "button"
	AriaRoleCheckbox   = "checkbox"
	AriaRoleCombobox   = "combobox"
	AriaRoleLink       = "link"
	AriaRoleMenuitem   = "menuitem"
	AriaRoleOption     = "option"
	AriaRoleRadio      = "radio"
	AriaRoleScrollbar  = "scrollbar"
	AriaRoleSlider     = "slider"
	AriaRoleSpinbutton = "spinbutton"
	AriaRoleSwitch     = "switch"
	AriaRoleTab        = "tab"
	AriaRoleTabpanel   = "tabpanel"
	AriaRoleTextbox    = "textbox"

	// Composite widget roles
	AriaRoleComboboxListbox = "listbox"
	AriaRoleMenu            = "menu"
	AriaRoleMenubar         = "menubar"
	AriaRoleRadiogroup      = "radiogroup"
	AriaRoleTablist         = "tablist"
	AriaRoleTree            = "tree"
	AriaRoleTreegrid        = "treegrid"
	AriaRoleGrid            = "grid"
	AriaRoleListbox         = "listbox"

	// Live region roles
	AriaRoleAlert       = "alert"
	AriaRoleAlertdialog = "alertdialog"
	AriaRoleLog         = "log"
	AriaRoleMarquee     = "marquee"
	AriaRoleStatus      = "status"
	AriaRoleTimer       = "timer"

	// Window roles
	AriaRoleDialog = "dialog"

	// Table roles
	AriaRoleTable        = "table"
	AriaRoleRow          = "row"
	AriaRoleRowgroup     = "rowgroup"
	AriaRoleCell         = "cell"
	AriaRoleColumnheader = "columnheader"
	AriaRoleRowheader    = "rowheader"
	AriaRoleGridcell     = "gridcell"

	// Form roles
	AriaRoleForm = "form"

	// Generic role
	AriaRoleGeneric = "generic"
)

// ARIA properties constants
const (
	// Global properties
	AriaLabel       = "aria-label"
	AriaLabelledby  = "aria-labelledby"
	AriaDescribedby = "aria-describedby"
	AriaHidden      = "aria-hidden"
	AriaLive        = "aria-live"
	AriaAtomic      = "aria-atomic"
	AriaRelevant    = "aria-relevant"
	AriaOwns        = "aria-owns"
	AriaControls    = "aria-controls"
	AriaFlowto      = "aria-flowto"

	// Widget properties
	AriaChecked         = "aria-checked"
	AriaDisabled        = "aria-disabled"
	AriaExpanded        = "aria-expanded"
	AriaHaspopup        = "aria-haspopup"
	AriaLevel           = "aria-level"
	AriaModal           = "aria-modal"
	AriaMultiline       = "aria-multiline"
	AriaMultiselectable = "aria-multiselectable"
	AriaOrientation     = "aria-orientation"
	AriaPressed         = "aria-pressed"
	AriaReadonly        = "aria-readonly"
	AriaRequired        = "aria-required"
	AriaSelected        = "aria-selected"
	AriaSort            = "aria-sort"
	AriaValuemax        = "aria-valuemax"
	AriaValuemin        = "aria-valuemin"
	AriaValuenow        = "aria-valuenow"
	AriaValuetext       = "aria-valuetext"

	// Relationship properties
	AriaActivedescendant = "aria-activedescendant"
	AriaColcount         = "aria-colcount"
	AriaColindex         = "aria-colindex"
	AriaColspan          = "aria-colspan"
	AriaDetails          = "aria-details"
	AriaErrormessage     = "aria-errormessage"
	AriaKeyshortcuts     = "aria-keyshortcuts"
	AriaPosinset         = "aria-posinset"
	AriaRowcount         = "aria-rowcount"
	AriaRowindex         = "aria-rowindex"
	AriaRowspan          = "aria-rowspan"
	AriaSetsize          = "aria-setsize"

	// Live region properties
	AriaBusy = "aria-busy"
)

// ARIA property values constants
const (
	// aria-live values
	AriaLiveOff       = "off"
	AriaLivePolite    = "polite"
	AriaLiveAssertive = "assertive"

	// aria-haspopup values
	AriaHaspopupFalse   = "false"
	AriaHaspopupTrue    = "true"
	AriaHaspopupMenu    = "menu"
	AriaHaspopupListbox = "listbox"
	AriaHaspopupTree    = "tree"
	AriaHaspopupGrid    = "grid"
	AriaHaspopupDialog  = "dialog"

	// aria-orientation values
	AriaOrientationVertical   = "vertical"
	AriaOrientationHorizontal = "horizontal"

	// aria-sort values
	AriaSortNone       = "none"
	AriaSortAscending  = "ascending"
	AriaSortDescending = "descending"
	AriaSortOther      = "other"

	// aria-checked values
	AriaCheckedTrue  = "true"
	AriaCheckedFalse = "false"
	AriaCheckedMixed = "mixed"

	// aria-expanded values
	AriaExpandedTrue  = "true"
	AriaExpandedFalse = "false"

	// aria-pressed values
	AriaPressedTrue  = "true"
	AriaPressedFalse = "false"
	AriaPressedMixed = "mixed"

	// aria-selected values
	AriaSelectedTrue  = "true"
	AriaSelectedFalse = "false"

	// aria-disabled values
	AriaDisabledTrue  = "true"
	AriaDisabledFalse = "false"

	// aria-hidden values
	AriaHiddenTrue  = "true"
	AriaHiddenFalse = "false"

	// aria-modal values
	AriaModalTrue  = "true"
	AriaModalFalse = "false"

	// aria-required values
	AriaRequiredTrue  = "true"
	AriaRequiredFalse = "false"

	// aria-readonly values
	AriaReadonlyTrue  = "true"
	AriaReadonlyFalse = "false"

	// aria-multiline values
	AriaMultilineTrue  = "true"
	AriaMultilineFalse = "false"

	// aria-multiselectable values
	AriaMultiselectableTrue  = "true"
	AriaMultiselectableFalse = "false"

	// aria-atomic values
	AriaAtomicTrue  = "true"
	AriaAtomicFalse = "false"

	// aria-busy values
	AriaBusyTrue  = "true"
	AriaBusyFalse = "false"

	// aria-relevant values
	AriaRelevantAdditions = "additions"
	AriaRelevantRemovals  = "removals"
	AriaRelevantText      = "text"
	AriaRelevantAll       = "all"
)
