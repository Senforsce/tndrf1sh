package PDF

import (
	"fmt"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

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

func GetHebergement(c Cfg) core.Maroto {
	cfg := config.NewBuilder().
		WithMargins(10, 35, 10).
		Build()

	// darkGrayColor := getDarkGrayColor()
	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	m.AddRow(24,
		text.NewCol(12, fmt.Sprintf("Objet : %s", c.Objet), props.Text{
			Top: 0, Left: 6, Align: align.Left,
			Color: getAlmostBlackColor(),
		}),
	)

	m.AddRow(24)

	m.AddRow(16,
		text.NewCol(12, fmt.Sprintf(
			"Je soussignée %s %s %s, née le %s à %s",
			c.XTitre, c.XNom, c.XPrenom, c.XDateDeNaissance, c.XLieuDeNaissance,
		), props.Text{
			Top: 0, Left: 3, Align: align.Left,
			Color: getAlmostBlackColor(),
		}),
	)
	m.AddRow(16,

		text.NewCol(12, fmt.Sprintf(
			"atteste sur l'honneur héberger %s %s %s, née le %s à %s",
			c.YTitre, c.YNom, c.YPrenom, c.YDateDeNaissance, c.YLieuDeNaissance,
		), props.Text{
			Top: 0, Left: 3, Align: align.Left,
			Color: getAlmostBlackColor(),
		}),
	)

	m.AddRow(16,

		text.NewCol(12, fmt.Sprintf(
			"depuis le %s à l'adresse suivante",
			c.DepuisQuand,
		), props.Text{
			Top: 0, Left: 3, Align: align.Left,
			Color: getAlmostBlackColor(),
		}),
	)
	m.AddRow(16,

		text.NewCol(4, fmt.Sprintf(
			"%s",
			c.Addresse,
		), props.Text{
			Top: 0, Left: 3, Align: align.Left,
			Color: getAlmostBlackColor(),
		}),
	)
	m.AddRow(16,

		text.NewCol(12,
			"Veuillez recevoir, l'expression de mes sentiments distingués",
			props.Text{
				Top: 0, Left: 3, Align: align.Left,
				Color: getAlmostBlackColor(),
			}),
	)
	m.AddRow(16,
		text.NewCol(4, fmt.Sprintf(
			"%s %s,",
			c.XNom, c.XPrenom,
		), props.Text{
			Top: 0, Left: 9, Align: align.Left,
			Color: getAlmostBlackColor(),
		}),
	)

	m.AddRow(16,
		text.NewCol(4, fmt.Sprintf(
			"%s le %s,",
			c.LieuSignature, c.JourSignature,
		), props.Text{
			Top: 0, Left: 9, Align: align.Left,
			Color: getAlmostBlackColor(),
		}),
	)

	m.AddRow(24,
		image.NewFromFileCol(3, c.SignatureImgPath, props.Rect{
			Center:  true,
			Percent: 80,
		}),
	)

	return m
}

func getAlmostBlackColor() *props.Color {
	return &props.Color{
		Red:   10,
		Green: 10,
		Blue:  10,
	}
}

type Cfg struct {
	Objet            string
	XTitre           string
	XNom             string
	XPrenom          string
	XDateDeNaissance string
	XLieuDeNaissance string
	YTitre           string
	YNom             string
	YPrenom          string
	YDateDeNaissance string
	YLieuDeNaissance string
	Addresse         string
	DepuisQuand      string // todo date
	SignatureImgPath string
	LieuSignature    string
	JourSignature    string
}
