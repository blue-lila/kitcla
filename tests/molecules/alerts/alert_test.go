package alerts

import (
	"testing"

	"github.com/blue-lila/kitcla"
	"github.com/blue-lila/kitcla/goc"
	"github.com/blue-lila/kitcla/sup"
)

func TestAlertSuccessAlert(t *testing.T) {
	kit := kitcla.New()
	component := kit.Molecules.Alerts.Alert

	h := component.SuccessAlert("Garden Successfully Watered", "All 12 garden beds have been watered. Next watering scheduled for tomorrow.", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/success_alert_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/success_alert_1.html")
	sup.AddPage("SuccessAlert", html)
}

func TestAlertWarningAlert(t *testing.T) {
	kit := kitcla.New()
	component := kit.Molecules.Alerts.Alert

	h := component.WarningAlert("Low Soil Moisture Detected", "Garden bed #3 has low soil moisture levels. Consider watering soon to prevent plant stress.", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/warning_alert_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/warning_alert_1.html")
	sup.AddPage("WarningAlert", html)
}

func TestAlertDangerAlert(t *testing.T) {
	kit := kitcla.New()
	component := kit.Molecules.Alerts.Alert

	h := component.DangerAlert("Pest Infestation Detected", "Critical: Aphids detected on tomato plants in bed #5. Immediate action required to prevent crop damage.", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/danger_alert_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/danger_alert_1.html")
	sup.AddPage("DangerAlert", html)
}

func TestAlertInfoAlert(t *testing.T) {
	kit := kitcla.New()
	component := kit.Molecules.Alerts.Alert

	h := component.InfoAlert("Planting Season Reminder", "Spring planting season begins in 2 weeks. Review your seed inventory and prepare garden beds.", nil)
	html := goc.RenderRoot(h)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/info_alert_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/info_alert_1.html")
	sup.AddPage("InfoAlert", html)
}

func TestAlertOverview(t *testing.T) {
	kit := kitcla.New()
	component := kit.Molecules.Alerts.Alert

	html := goc.RenderRoot(
		kit.Component.Dcs("space-y-4 p-6",
			kit.Component.H("h1", "class", "text-2xl font-bold mb-4", "Alerts Overview"),
			component.SuccessAlert("Harvest Complete", "Successfully harvested 15 lbs of tomatoes from bed #2.", nil),
			component.InfoAlert("Weather Update", "Sunny weather expected for the next 3 days. Perfect conditions for outdoor gardening.", nil),
			component.WarningAlert("Fertilizer Running Low", "Only 2 bags of organic fertilizer remaining. Reorder soon to maintain your schedule.", nil),
			component.DangerAlert("Frost Warning", "Temperature will drop below freezing tonight. Protect sensitive plants immediately.", nil),
		),
	)

	sup.UpdateEqualHtmlFromDataFile(t, html, "./data/alert_overview_1.html")
	sup.AssertEqualHtmlFromDataFile(t, html, "./data/alert_overview_1.html")
	sup.AddIndexPage("alerts", html)
}
