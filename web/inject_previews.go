package main

import (
	"github.com/senforsce/coachmj/web/owl/about"
	"github.com/senforsce/coachmj/web/owl/abouthero"
	"github.com/senforsce/coachmj/web/owl/aboutlogo"
	"github.com/senforsce/coachmj/web/owl/aboutquote"
	"github.com/senforsce/coachmj/web/owl/adminsidebar"
	"github.com/senforsce/coachmj/web/owl/admintopbar"
	"github.com/senforsce/coachmj/web/owl/appointment"
	"github.com/senforsce/coachmj/web/owl/contactctapills"
	"github.com/senforsce/coachmj/web/owl/contactform"
	"github.com/senforsce/coachmj/web/owl/contactformwrapper"
	"github.com/senforsce/coachmj/web/owl/dashboardtopstats"
	"github.com/senforsce/coachmj/web/owl/eventlist"
	"github.com/senforsce/coachmj/web/owl/feature"
	"github.com/senforsce/coachmj/web/owl/footer"
	"github.com/senforsce/coachmj/web/owl/fullcalendar"
	"github.com/senforsce/coachmj/web/owl/header"
	"github.com/senforsce/coachmj/web/owl/headercarousel"
	"github.com/senforsce/coachmj/web/owl/headerstats"
	"github.com/senforsce/coachmj/web/owl/monthlystats"
	"github.com/senforsce/coachmj/web/owl/myvalues"
	"github.com/senforsce/coachmj/web/owl/navbar"
	"github.com/senforsce/coachmj/web/owl/quickchat"
	"github.com/senforsce/coachmj/web/owl/saleshistory"
	"github.com/senforsce/coachmj/web/owl/service"
	"github.com/senforsce/coachmj/web/owl/sitevisits"
	"github.com/senforsce/coachmj/web/owl/spinner"
	"github.com/senforsce/coachmj/web/owl/subscriptions"
	"github.com/senforsce/coachmj/web/owl/team"
	"github.com/senforsce/coachmj/web/owl/testimonial"
	"github.com/senforsce/coachmj/web/owl/topbar"
	"github.com/senforsce/coachmj/web/owl/values"
	"github.com/senforsce/coachmj/web/owl/vibelist"

	"github.com/senforsce/tndr0cean/router"
)

func WithHTMXPreviews(app *router.Tndr0cean) error {
	app.Get("/preview/about", about.Preview)
	app.Get("/preview/abouthero", abouthero.Preview)
	app.Get("/preview/aboutquote", aboutquote.Preview)
	app.Get("/preview/aboutlogo", aboutlogo.Preview)
	app.Get("/preview/adminsidebar", adminsidebar.Preview)
	app.Get("/preview/admintopbar", admintopbar.Preview)
	app.Get("/preview/appointment", appointment.Preview)
	app.Get("/preview/contactctapills", contactctapills.Preview)
	app.Get("/preview/contactform", contactform.Preview)
	app.Get("/preview/contactformwrapper", contactformwrapper.Preview)
	app.Get("/preview/dashboardtopstats", dashboardtopstats.Preview)
	app.Get("/preview/eventlist", eventlist.Preview)
	app.Get("/preview/feature", feature.Preview)
	app.Get("/preview/footer", footer.Preview)
	app.Get("/preview/fullcalendar", fullcalendar.Preview)
	app.Get("/preview/header", header.Preview)
	app.Get("/preview/headercarousel", headercarousel.Preview)
	app.Get("/preview/headerstats", headerstats.Preview)
	app.Get("/preview/monthlystats", monthlystats.Preview)
	app.Get("/preview/myvalues", myvalues.Preview)
	app.Get("/preview/navbar", navbar.Preview)
	app.Get("/preview/quickchat", quickchat.Preview)
	app.Get("/preview/saleshistory", saleshistory.Preview)
	app.Get("/preview/service", service.Preview)
	app.Get("/preview/sitevisits", sitevisits.Preview)
	app.Get("/preview/spinner", spinner.Preview)
	app.Get("/preview/subscriptions", subscriptions.Preview)
	app.Get("/preview/team", team.Preview)
	app.Get("/preview/testimonial", testimonial.Preview)
	app.Get("/preview/topbar", topbar.Preview)
	app.Get("/preview/values", values.Preview)
	app.Get("/preview/vibelist", vibelist.Preview)

	return nil
}
