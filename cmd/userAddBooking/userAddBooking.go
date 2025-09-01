package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/senforsce/userconfig"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/senforsce/sparql"
)

//go:embed addBookingCoaching.sparql
var AddBookingCoaching string

//go:embed addBookingProgramme.sparql
var AddBookingProgramme string

var (
	userId              string
	bookingComment      string
	bookingRequestDate  string
	bookingStatusString string
	bookingType         string
)

func init() {
	flag.StringVar(&userId, "i", "http://senforsce.com/o8/brain/Some-ID-123", "The User Id created by `createUser` command")
	flag.StringVar(&bookingComment, "c", "Un Commentaire de Réservation", "The Booking Comment")
	flag.StringVar(&bookingRequestDate, "rd", "2025-31-12T10:00", "The Request Date od the booking (rfc:3339)")
	flag.StringVar(&bookingStatusString, "s", "DemandeEnvoyee", "The status of the Booking (DemandeEnvoyee|DemandeAcceptee|BookingEnded)")
	flag.StringVar(&bookingType, "t", "Coaching", "The booking type (Coaching|Programme)")

}

func main() {
	projectRoot := os.Getenv("TNDR_GENERATE_ROOT")
	generateKey := os.Getenv("TNDR_GENERATE_KEY")

	flag.Parse()

	if projectRoot == "" {
		log.Fatal("Please Set the environment variable TNDR_GENERATE_ROOT")
	}

	if generateKey == "" {
		log.Fatal("Please Set the environment variable TNDR_GENERATE_KEY")
	}

	fmt.Printf("the userId is %s\n", userId)
	fmt.Printf("the bookingComment is %s\n", bookingComment)
	fmt.Printf("the bookingRequestDate is %s\n", bookingRequestDate)
	fmt.Printf("the bookingStatusString is %s\n", bookingStatusString)
	fmt.Printf("the bookingType is %s\n", bookingType)

	userBookingId, err := gonanoid.New()

	if err != nil {
		log.Fatalf("Unable to Generate A userBookingId")
	}

	infoIdentifier := fmt.Sprintf("http://senforsce.com/o8/brain/Booking-%s", userBookingId)

	fmt.Printf("the infoIdentifier  is %s\n", infoIdentifier)
	userConfig := userconfig.NewUserConfig()

	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	var formattedQuery string
	switch bookingType {
	case "Coaching":
		{
			formattedQuery = fmt.Sprintf(AddBookingCoaching,
				userId,
				infoIdentifier,
				bookingComment,
				bookingRequestDate,
				bookingStatusString,
				bookingType,
			)
		}
	case "Programme":
		{
			formattedQuery = fmt.Sprintf(AddBookingProgramme,
				userId,
				infoIdentifier,
				bookingComment,
				bookingRequestDate,
				bookingStatusString,
				bookingType,
			)
		}
	}

	log.Println(formattedQuery)

	r, err := repo.Update(formattedQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)

	fmt.Printf("\n SUCCESSFULLY CREATED : %s\n", infoIdentifier)

}
