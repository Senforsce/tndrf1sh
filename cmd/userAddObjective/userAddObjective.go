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

//go:embed addObjectives.sparql
var AddObjectivesQuery string

var (
	userId                   string
	objectiveName            string
	objectiveKPIInitialDate  string
	objectiveKPIInitialValue string
	objectiveKPITargetValue  string
	objectiveKPIUnit         string
	objectiveKPIName         string
	objectiveTargetDate      string
)

func init() {
	flag.StringVar(&userId, "i", "http://senforsce.com/o8/brain/Some-ID-123", "The User Id created by `createUser` command")
	flag.StringVar(&objectiveName, "on", "My Objective", "The Objective Name of the User")
	flag.StringVar(&objectiveKPIInitialDate, "kid", "2024-31-12T10:00", "The Objective KPI initial date")
	flag.StringVar(&objectiveKPIInitialValue, "kiv", "10", "The Objective KPI initial value")
	flag.StringVar(&objectiveKPITargetValue, "ktv", "100", "The Objective KPI target value")
	flag.StringVar(&objectiveKPIUnit, "ku", "reps", "The Objective KPI unit")
	flag.StringVar(&objectiveKPIName, "kn", "My Objective", "The Objective Kpi name")
	flag.StringVar(&objectiveTargetDate, "td", "2025-31-12T10:00", "The Objective KPI target date")

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
	fmt.Printf("the objectiveName is %s\n", objectiveName)
	fmt.Printf("the objectiveKPIInitialDate is %s\n", objectiveKPIInitialDate)
	fmt.Printf("the objectiveKPIInitialValue is %s\n", objectiveKPIInitialValue)
	fmt.Printf("the objectiveKPITargetValue is %s\n", objectiveKPITargetValue)
	fmt.Printf("the objectiveKPIUnit is %s\n", objectiveKPIUnit)
	fmt.Printf("the objectiveKPIName is %s\n", objectiveKPIName)
	fmt.Printf("the objectiveTargetDate is %s\n", objectiveTargetDate)

	userObjectiveId, err := gonanoid.New()

	if err != nil {
		log.Fatalf("Unable to Generate A userObjectiveId")
	}

	infoIdentifier := fmt.Sprintf("http://senforsce.com/o8/brain/userObjective-%s", userObjectiveId)

	fmt.Printf("the infoIdentifier  is %s\n", infoIdentifier)
	userConfig := userconfig.NewUserConfig()

	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	formattedQuery := fmt.Sprintf(AddObjectivesQuery,
		infoIdentifier,
		userId,
		objectiveKPIInitialDate,
		objectiveKPIInitialValue,
		objectiveKPIName,
		objectiveKPITargetValue,
		objectiveKPIUnit,
		objectiveName,
		objectiveTargetDate,
		objectiveName,
	)

	log.Println(formattedQuery)

	r, err := repo.Update(formattedQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)

	fmt.Printf("\n SUCCESSFULLY CREATED : %s\n", infoIdentifier)

}
