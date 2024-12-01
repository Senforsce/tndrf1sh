package main

import (
	"log"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"

	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/billingv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/billingv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	cfg := config.NewBuilder().
		WithPageNumber("Page {current} sur {total}", props.RightBottom).
		WithMargins(10, 15, 10).
		Build()

	darkGrayColor := getDarkGrayColor()
	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	err := m.RegisterHeader(getPageHeader())
	if err != nil {
		log.Fatal(err.Error())
	}

	err = m.RegisterFooter(getPageFooter())
	if err != nil {
		log.Fatal(err.Error())
	}

	m.AddRows(text.NewRow(10, "Devis DE-MON-00001", props.Text{
		Top:   3,
		Style: fontstyle.Bold,
		Align: align.Center,
	}))

	m.AddRow(7,
		text.NewCol(12, "Création de la présence en ligne de M.O.N Burger", props.Text{
			Top:   1.5,
			Size:  9,
			Style: fontstyle.Bold,
			Align: align.Center,
			Color: &props.WhiteColor,
		}),
	).WithStyle(&props.Cell{BackgroundColor: darkGrayColor})

	m.AddRows(getTransactions()...)

	return m
}

func getTransactions() []core.Row {
	rows := []core.Row{
		row.New(6).Add(
			text.NewCol(2, "Référence", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(5, "Désignation", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(1, "Qté", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(1, "PU HT", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
			text.NewCol(3, "Total HT", props.Text{Size: 9, Align: align.Center, Style: fontstyle.Bold}),
		),
	}

	var contentsRow []core.Row
	contents := getContents()
	/*for i := 0; i < 8; i++ {
	    contents = append(contents, contents...)
	}*/

	for i, content := range contents {
		r := row.New(6).Add(
			text.NewCol(2, content[0], props.Text{Size: 8, Align: align.Center}),
			text.NewCol(5, content[1], props.Text{Size: 8, Align: align.Center}),
			text.NewCol(1, content[2], props.Text{Size: 8, Align: align.Center}),
			text.NewCol(1, content[3], props.Text{Size: 8, Align: align.Center}),
			text.NewCol(3, content[4], props.Text{Size: 8, Align: align.Center}),
		)
		if i%2 == 0 {
			gray := getGrayColor()
			r.WithStyle(&props.Cell{BackgroundColor: gray})
		}

		contentsRow = append(contentsRow, r)
	}

	rows = append(rows, contentsRow...)

	rows = append(rows, row.New(20).Add(
		text.NewCol(5, "Durée de validité : 30 jours; Taux d'escompte : Pas d'escompte applicable. Date de livraison : 26-10-2024 Mode de paiement : Virement; Sous régime auto-entrepreneur, pas de TVA á facturer", props.Text{
			Style: fontstyle.Italic,
			Size:  8,
			Align: align.Left,
		}),

		col.New(2),
		text.NewCol(2, "Sous-Total HT:", props.Text{
			Top:   20,
			Size:  7,
			Align: align.Right,
		}),
		text.NewCol(3, "EUR 28000,00", props.Text{
			Top:   20,
			Style: fontstyle.Bold,
			Size:  7,
			Align: align.Center,
		}),
	))

	rows = append(rows, row.New(5).Add(
		col.New(7),

		text.NewCol(2, "Remise 1er Client (-75%):", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  7,
			Align: align.Right,
		}),
		text.NewCol(3, "EUR -21000,00", props.Text{
			Top:   5,
			Color: getRedColor(),
			Size:  7,
			Align: align.Center,
		}),
	))

	rows = append(rows, row.New(5).Add(
		col.New(7),

		text.NewCol(2, "Total:", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(3, "EUR 7000,00", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Center,
		}),
	))

	rows = append(rows, row.New(3).Add(
		text.NewCol(2, "Bon pour accord le:", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Left,
		}),
		text.NewCol(3, "23rd September 2024", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Center,
		}),
	))

	rows = append(rows, row.New(30).Add(
		text.NewCol(2, "Signature:", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
		}),
		signature.NewCol(4, "signature", props.Signature{}),
	))

	rows = append(rows, row.New(6).Add(
		text.NewCol(6, "Délai prévisionnel de réalisation :", props.Text{
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Left,
		}),
	))

	rows = append(rows, row.New(10).Add(
		text.NewCol(12, "Livraison (mise en production du site) sous 3 à 4 semaines à compter de la réception du devis dûment accepté, sous réserve de la fourniture des contenus et de l'accès à un espace d'hébergement opérationnel conformes aux spécifications.", props.Text{
			Style: fontstyle.Italic,
			Size:  8,
			Align: align.Left,
		}),
	))

	rows = append(rows, row.New(3).Add(
		text.NewCol(2, "Hebergement:", props.Text{
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Left,
		}),
	))

	rows = append(rows, row.New(10).Add(
		text.NewCol(12, "L'hébergement du site ne fait pas partie du présent devis. Les pré-requis pour l'hébergement sont : Serveur Unix (type Apache)", props.Text{
			Size:  8,
			Align: align.Left,
		}),
	))

	rows = append(rows, row.New(3).Add(
		text.NewCol(2, "Accompte:", props.Text{
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Left,
		}),
	))

	rows = append(rows, row.New(2).Add(
		text.NewCol(5, "14% à la commande", props.Text{
			Size:  8,
			Align: align.Left,
		}),
	))

	return rows
}

func getPageHeader() core.Row {
	return row.New(22).Add(
		image.NewFromFileCol(2, "docs/assets/images/abdoul-sy.png", props.Rect{
			Center:  true,
			Percent: 80,
		}),
		col.New(6),
		col.New(3).Add(
			text.New("Senforsce - SIRET: 929531903 00011, CODE: 7112B. \n 9 rue des Prairies, 75020 Paris, France.", props.Text{
				Size:  8,
				Align: align.Right,
				Color: getRedColor(),
			}),
			text.New("Tel: +33 666 5954 69", props.Text{
				Top:   12,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Right,
				Color: getBlueColor(),
			}),
			text.New("www.senforsce.com", props.Text{
				Top:   15,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Right,
				Color: getBlueColor(),
			}),
		),
	)
}

func getPageFooter() core.Row {
	return row.New(20).Add(
		col.New(12).Add(
			text.New("Tel: +33 666 5954 69", props.Text{
				Top:   13,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: getBlueColor(),
			}),
			text.New("www.senforsce.com", props.Text{
				Top:   16,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Left,
				Color: getBlueColor(),
			}),
		),
	)
}

func getDarkGrayColor() *props.Color {
	return &props.Color{
		Red:   55,
		Green: 55,
		Blue:  55,
	}
}

func getGrayColor() *props.Color {
	return &props.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
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

func getContents() [][]string {
	// per columns, second is span 2, but id should fill it
	return [][]string{
		{"MON-10", "Online presence of M.O.N Burger to increase visibility of resto and products", "21", "500 EUR", "10500 EUR"},
		{"MON-11", "Manage M.O.N Burger Promotions", "21", "500 EUR", "10500 EUR"},
		{"MON-14", "M.O.N Burger Business to Business", "7", "500 EUR", "3500 EUR"},
		{"MON-13", "Enhance M.O.N Burger's customer experience", "7", "500 EUR", "3500 EUR"},
		{"MON-12", "Make M.O.N Burger's Customer feel special", "21", "5 EUR", "105 EUR"},
	}
}
