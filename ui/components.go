package ui

import "charm.land/lipgloss/v2"

/*
	-----------------------------------------------------
	HEADER SECTION
	-----------------------------------------------------
*/

var heading = `
███╗   ██╗ ██████╗ ██╗███╗   ██╗██╗  ██╗ ██████╗████████╗██╗
████╗  ██║██╔════╝ ██║████╗  ██║╚██╗██╔╝██╔════╝╚══██╔══╝██║
██╔██╗ ██║██║  ███╗██║██╔██╗ ██║ ╚███╔╝ ██║        ██║   ██║
██║╚██╗██║██║   ██║██║██║╚██╗██║ ██╔██╗ ██║        ██║   ██║
██║ ╚████║╚██████╔╝██║██║ ╚████║██╔╝ ██╗╚██████╗   ██║   ███████╗
╚═╝  ╚═══╝ ╚═════╝ ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝ ╚═════╝   ╚═╝   ╚══════╝
`

var HeadingStyled = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#7D56F4")).
	PaddingTop(2).
	PaddingLeft(4).
	Render(heading)

var line = "----------------------------------------------------------------"
var LineStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#E762FE")).
	PaddingTop(1).
	PaddingLeft(4).
	Render(line)

var subHeading = " - NGINX CONFIGRATION MANAGER, BECAUSE WHY NOT - "
var SubHeadingStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#E762FE")).
	PaddingTop(1).
	PaddingLeft(10).
	Render(subHeading)

var BoxStyle = lipgloss.NewStyle().
	MarginTop(2).
	MarginLeft(2).
	MarginBottom(1).
	Foreground(lipgloss.Color("#7C3AED")).
	BorderStyle(lipgloss.NormalBorder()).
	BorderLeft(true).
	BorderForeground(lipgloss.Color("#7C3AED")).
	Background(lipgloss.Color("#121117")).
	Padding(1, 2).
	Width(68)

var Header = HeadingStyled + LineStyle + SubHeadingStyle + LineStyle

/*
	-----------------------------------------------------
	ITEMS SECTION
	-----------------------------------------------------
*/

var ItemStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#7C3AED")).
	Width(100)

var InputBoxStyle = lipgloss.NewStyle().
	MarginLeft(2).
	MarginBottom(1).
	Foreground(lipgloss.Color("#7C3AED")).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("#7C3AED")).
	Padding(1, 1).
	Width(68)

var resetConfigDisclaimer = `
	⚠ DANGER: This will permanently erase ALL existing Nginx configuration.

	The following will be deleted:
		• nginx.conf
		• sites-available/*
		• sites-enabled/*
		• Virtual host configurations
		• Reverse proxy settings
		• SSL/TLS configurations
		• Redirect rules
		• Custom Nginx settings

	This operation cannot be undone.

	Your websites will stop working immediately after the reset.

	PRESS 'CTRL S' to Confirm:
`

var ResetConfigDisclaimerStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("196")).
	Render(resetConfigDisclaimer)

var WarningBoxStyle = lipgloss.NewStyle().
	MarginTop(2).
	MarginLeft(2).
	MarginBottom(1).
	Foreground(lipgloss.Color("196")).
	BorderStyle(lipgloss.NormalBorder()).
	BorderLeft(true).
	BorderForeground(lipgloss.Color("196")).
	Background(lipgloss.Color("#121117")).
	Padding(1, 2).
	Width(68)
