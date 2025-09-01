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

//go:embed addObjectiveKPI.sparql
var AddObjectivesKPIQuery string

var (
	userId               string
	belongsToObjectiveId string
	objectiveKPIComment  string
	objectiveKPIDate     string
	objectiveKPIValue    string
)

func init() {
	flag.StringVar(&userId, "i", "http://senforsce.com/o8/brain/Some-ID-123", "The User Id created by `createUser` command")
	flag.StringVar(&belongsToObjectiveId, "o", "http://senforsce.com/o8/brain/Ojective-Some-ID-123", "The Objective Id created by `userAddObjectives` command")

	flag.StringVar(&objectiveKPIComment, "kc", "My Objective KPI Comment", "The Objective KPI Comment")
	flag.StringVar(&objectiveKPIDate, "kd", "2024-31-12T10:00", "The Objective KPI's date")
	flag.StringVar(&objectiveKPIValue, "kv", "25", "The Objective KPI's value")

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
	fmt.Printf("the belongsToObjectiveId is %s\n", belongsToObjectiveId)
	fmt.Printf("the objectiveKPIComment is %s\n", objectiveKPIComment)
	fmt.Printf("the objectiveKPIDate is %s\n", objectiveKPIDate)

	userObjectiveKPIId, err := gonanoid.New()

	if err != nil {
		log.Fatalf("Unable to Generate A userObjectiveKPIId")
	}

	infoIdentifier := fmt.Sprintf("http://senforsce.com/o8/brain/userObjectiveKPI-%s", userObjectiveKPIId)

	fmt.Printf("the infoIdentifier  is %s\n", infoIdentifier)
	userConfig := userconfig.NewUserConfig()

	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	formattedQuery := fmt.Sprintf(AddObjectivesKPIQuery,
		infoIdentifier,
		belongsToObjectiveId,
		userId,
		objectiveKPIComment,
		objectiveKPIDate,
		objectiveKPIValue,
	)

	log.Println(formattedQuery)

	r, err := repo.Update(formattedQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)

	fmt.Printf("\n SUCCESSFULLY CREATED : %s\n", infoIdentifier)

}
