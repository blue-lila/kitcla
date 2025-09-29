package sidebars

import (
	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
	"testing"
)

func TestSidebarLinkContainer(t *testing.T) {
	//kit := kitcla.New()
	//component := kit.Organisms.Sidebars.Sidebar
	//
	//h := component.LinkContainer(nil)
	//html := goc.RenderRoot(h)
	//
	//sup.UpdateEqualHtmlFromDataFile(t, html, "./data/sidebar_link_container_1.html")
	//sup.AssertEqualHtmlFromDataFile(t, html, "./data/sidebar_link_container_1.html")
}

func TestSidebarLogoContainer(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Sidebars.Sidebar

	h := component.LogoContainer()
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/sidebar_logo_container_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/sidebar_logo_container_1.html")
}

func TestSidebarLogoImg(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Sidebars.Sidebar

	h := component.LogoImg()
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/sidebar_logo_img_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/sidebar_logo_img_1.html")
}

func TestSidebarLogoTitle(t *testing.T) {
	kit := kitcla.New()
	component := kit.Organisms.Sidebars.Sidebar

	h := component.LogoTitle()
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/sidebar_logo_title_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/sidebar_logo_title_1.html")
}

func TestSidebarMenuItemLink(t *testing.T) {
	//kit := kitcla.New()
	//component := kit.Organisms.Sidebars.Sidebar
	//
	//h := component.MenuItemLink(nil)
	//html := goc.RenderRoot(h)
	//
	//sup.UpdateEqualHtmlFromDataFile(t, html, "./data/sidebar_menu_item_link_1.html")
	//sup.AssertEqualHtmlFromDataFile(t, html, "./data/sidebar_menu_item_link_1.html")
}

func TestSidebarMenuItems(t *testing.T) {
	//kit := kitcla.New()
	//component := kit.Organisms.Sidebars.Sidebar
	//
	//h := component.MenuItems(nil)
	//html := goc.RenderRoot(h)
	//
	//sup.UpdateEqualHtmlFromDataFile(t, html, "./data/sidebar_menu_items_1.html")
	//sup.AssertEqualHtmlFromDataFile(t, html, "./data/sidebar_menu_items_1.html")
}

func TestSidebarSidebar(t *testing.T) {
	//kit := kitcla.New()
	//component := kit.Organisms.Sidebars.Sidebar
	//
	//h := component.Sidebar()
	//html := goc.RenderRoot(h)
	//
	//sup.UpdateEqualHtmlFromDataFile(t, html, "./data/sidebar_sidebar_1.html")
	//sup.AssertEqualHtmlFromDataFile(t, html, "./data/sidebar_sidebar_1.html")
}
