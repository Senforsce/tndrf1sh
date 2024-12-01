package PDF

import (
	"fmt"
	"log"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
)

// func main() {
// 	m := GetMaroto()
// 	document, err := m.Generate()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = document.Save("docs/assets/pdf/simplestv2.pdf")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func GetNDA(c Config) core.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber("Page {current} of {total}", props.RightBottom).
		WithMargins(10, 15, 10).
		Build()

	// darkGrayColor := getDarkGrayColor()
	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	err := m.RegisterHeader(getPageHeader(c))
	if err != nil {
		log.Fatal(err.Error())
	}

	err = m.RegisterFooter(getPageFooter(c))
	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRow(40,
		text.NewCol(12, "Left-aligned text", props.Text{Top: 3, Left: 3, Align: align.Left}),
	)

	// sen:TODO: Create a mapping to ease the generation of pdf contracts

	return m
}

func getDarkGrayColor() *props.Color {
	return &props.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

type Config struct {
	ImgFilePath               string
	CompanyOrPersonName       string
	CompanyOrPersonAddress    string
	CompanyOrPersonCity       string
	CompanyOrPersonPhone      string
	CompanyOrPersonEmail      string
	CompanyOrPersonWebsiteURL string
}

func getPageHeader(c Config) core.Row {
	return row.New(20).Add(
		image.NewFromFileCol(3, c.ImgFilePath, props.Rect{
			Center:  true,
			Percent: 80,
		}),
		col.New(6),
		col.New(3).Add(
			text.New(fmt.Sprintf("%s %s %s", c.CompanyOrPersonName, c.CompanyOrPersonAddress, c.CompanyOrPersonCity), props.Text{
				Size:  8,
				Align: align.Right,
				Color: getRedColor(),
			}),
			text.New(fmt.Sprintf("Tel: %s, Email:", c.CompanyOrPersonPhone, c.CompanyOrPersonEmail), props.Text{
				Top:   12,
				Style: fontstyle.BoldItalic,
				Size:  10,
				Align: align.Right,
				Color: getBlueColor(),
			}),
			text.New(c.CompanyOrPersonWebsiteURL, props.Text{
				Top:   15,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Right,
				Color: getBlueColor(),
			}),
		),
	)
}

func getBlueColor() *props.Color {
	return &props.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

func getRedColor() *props.Color {
	return &props.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}

func getPageFooter(c Config) core.Row {
	return row.New(20).Add(
		col.New(12).Add(
			text.New(fmt.Sprintf("Tel:%s, E-mail: %s", c.CompanyOrPersonPhone, c.CompanyOrPersonEmail), props.Text{
				Top:   13,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: getBlueColor(),
			}),
			text.New(c.CompanyOrPersonWebsiteURL, props.Text{
				Top:   16,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: getBlueColor(),
			}),
		),
	)
}
