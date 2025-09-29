package navmenus

import (
	"github.com/blue-lila/kitcla/components"
	"github.com/blue-lila/kitcla/components/atoms/icons"
	"github.com/blue-lila/kitcla/goc"
)

type Navmenu struct {
	Component *components.Component
	Icon      *icons.Icon
}

type NavmenuMenuItem struct {
	Label       string
	Link        string
	Description string
}

type NavmenuMenu struct {
	Key   string
	Label string
	Icon  string
	Items []*NavmenuMenuItem
}

type NavmenuMod struct {
	Menus []*NavmenuMenu
}

func (this *Navmenu) H(mod *NavmenuMod) goc.HTML {
	return this.Component.H("nav", goc.Attr{"x-data": this.component(), "class": "relative z-10 w-auto"},
		this.menusLabelsContainer(mod),
		this.navigationDropdown(mod),
	)
}

func (this *Navmenu) Mod() *NavmenuMod {
	return &NavmenuMod{}
}

func (this *Navmenu) menusLabelsContainer(mod *NavmenuMod) goc.HTML {
	return this.Component.H("div", goc.Attr{"class": "relative"}, this.menusLabels(mod))
}
func (this *Navmenu) menusLabels(mod *NavmenuMod) goc.HTML {
	var set []goc.HTML

	for _, item := range mod.Menus {
		set = append(set, this.menuLabel(mod, item))
	}

	return this.Component.Cas("ul", goc.Attr{
		"class": "flex items-center justify-center flex-1 p-1 space-x-1 list-none border rounded-md text-neutral-700 group border-neutral-200/80",
	},
		set...,
	)
}
func (this *Navmenu) menuLabel(mod *NavmenuMod, item *NavmenuMenu) goc.HTML {
	return this.Component.H("li", this.openMenuButton(mod, item))
}
func (this *Navmenu) openMenuButton(mod *NavmenuMod, item *NavmenuMenu) goc.HTML {
	return this.Component.H("button", goc.Attr{
		":class":      "{ 'bg-neutral-100' : navigationMenu=='" + item.Key + "', 'hover:bg-neutral-100' : navigationMenu!='" + item.Key + "' }",
		"@mouseover":  "navigationMenuOpen=true; navigationMenuReposition($el); navigationMenu='" + item.Key + "'",
		"@mouseleave": "navigationMenuLeave()",
		"class":       "inline-flex items-center justify-center h-8 px-4 py-1 text-sm font-medium transition-colors rounded-md hover:text-neutral-900 focus:outline-none disabled:opacity-50 disabled:pointer-events-none group w-max"},
		this.openMenuButtonLabel(mod, item), this.openMenuButtonIcon(mod, item),
	)
}
func (this *Navmenu) openMenuButtonLabel(mod *NavmenuMod, item *NavmenuMenu) goc.HTML {
	if item.Icon != "" {
		return this.Icon.Icon(item.Icon)
	}
	return this.Component.H("span", item.Label)
}
func (this *Navmenu) openMenuButtonIcon(mod *NavmenuMod, item *NavmenuMenu) goc.HTML {
	return this.Component.H("svg", goc.Attr{":class": "{ '-rotate-180' : navigationMenuOpen==true && navigationMenu == '" + item.Key + "' }", "class": "relative top-[1px] ml-1 h-3 w-3 ease-out duration-300", "xmlns": "http://www.w3.org/2000/svg", "viewBox": "0 0 24 24", "fill": "none", "stroke": "currentColor", "stroke-width": "2", "stroke-linecap": "round", "stroke-linejoin": "round", "aria-hidden": "true"}, this.polyline(mod))
}
func (this *Navmenu) polyline(mod *NavmenuMod) goc.HTML {
	return this.Component.H("polyline", goc.Attr{"points": "6 9 12 15 18 9"})
}

func (this *Navmenu) navigationDropdown(mod *NavmenuMod) goc.HTML {
	return this.Component.H("div", goc.Attr{
		"x-ref":                    "navigationDropdown",
		"x-show":                   "navigationMenuOpen",
		"x-transition:enter":       "transition ease-out duration-100",
		"x-transition:enter-start": "opacity-0 scale-90",
		"x-transition:enter-end":   "opacity-100 scale-100",
		"x-transition:leave":       "transition ease-in duration-100",
		"x-transition:leave-start": "opacity-100 scale-100",
		"x-transition:leave-end":   "opacity-0 scale-90",
		"@mouseover":               "navigationMenuClearCloseTimeout()",
		"@mouseleave":              "navigationMenuLeave()",
		"class":                    "absolute top-0 pt-3 duration-200 ease-out -translate-x-1/2 translate-y-11",
		"x-cloak":                  ""},
		this.container(mod))
}
func (this *Navmenu) container(mod *NavmenuMod) goc.HTML {
	var set []goc.HTML

	for _, item := range mod.Menus {
		set = append(set, this.menu(mod, item))
	}

	return this.Component.Cas("div", goc.Attr{
		"class": "flex justify-center w-auto h-auto overflow-hidden bg-white border rounded-md shadow-sm border-neutral-200/70",
	},
		set...,
	)
}
func (this *Navmenu) menu(mod *NavmenuMod, item *NavmenuMenu) goc.HTML {
	return this.Component.H("div", goc.Attr{"x-show": "navigationMenu == '" + item.Key + "'", "class": "flex items-stretch justify-center w-full max-w-2xl p-6 gap-x-3"},
		this.menuContainer(mod, item),
	)
}

func (this *Navmenu) menuContainer(mod *NavmenuMod, menu *NavmenuMenu) goc.HTML {
	var set []goc.HTML

	for _, item := range menu.Items {
		set = append(set, this.menuItem(mod, item))
	}

	return this.Component.Cas("div", goc.Attr{"class": "w-72"},
		set...,
	)
}
func (this *Navmenu) menuItem(mod *NavmenuMod, item *NavmenuMenuItem) goc.HTML {
	return this.Component.H("a",
		goc.Attr{"href": item.Link,
			"@click": "navigationMenuClose()",
			"class":  "block px-3.5 py-3 text-sm rounded hover:bg-neutral-100",
		},
		this.menuItemLabel(mod, item),
		this.menuItemDescription(mod, item),
	)
}
func (this *Navmenu) menuItemLabel(mod *NavmenuMod, item *NavmenuMenuItem) goc.HTML {
	return this.Component.H("span", goc.Attr{"class": "block mb-1 font-medium text-black"}, item.Label)
}
func (this *Navmenu) menuItemDescription(mod *NavmenuMod, item *NavmenuMenuItem) goc.HTML {
	if item.Description == "" {
		return goc.HTML{}
	}

	return this.Component.H("span",
		goc.Attr{"class": "block font-light leading-5 opacity-50"},
		item.Description)
}

func (this *Navmenu) component() string {
	return `{
        navigationMenuOpen: false,
        navigationMenu: '',
        navigationMenuCloseDelay: 200,
        navigationMenuCloseTimeout: null,
        navigationMenuLeave() {
            let that = this;
            this.navigationMenuCloseTimeout = setTimeout(() => {
                that.navigationMenuClose();
            }, this.navigationMenuCloseDelay);
        },
        navigationMenuReposition(navElement) {
            this.navigationMenuClearCloseTimeout();
            this.$refs.navigationDropdown.style.left = navElement.offsetLeft + 'px';
            this.$refs.navigationDropdown.style.marginLeft = (navElement.offsetWidth/2) + 'px';
        },
        navigationMenuClearCloseTimeout(){
            clearTimeout(this.navigationMenuCloseTimeout);
        },
        navigationMenuClose(){
            this.navigationMenuOpen = false;
            this.navigationMenu = '';
        }
    }`
}
