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

//go:embed editProgramExercice.sparql
var editProgramExercise string

var (
	userId             string
	exerciceId         string
	exerciseName       string
	exercicePRBest     string
	exercicePRLast     string
	exercicePRMin      string
	exercicePRUnit     string
	exerciceReps       string
	exerciceSets       string
	exerciceTimeUnit   string
	exerciceWeight     string
	exerciceWeightUnit string
	exerciseTime       string
	hasSportMovement   string
)

func init() {
	flag.StringVar(&userId, "i", "http://senforsce.com/o8/brain/Some-ID-123", "The User Id created by `createUser` command")
	flag.StringVar(&exerciceId, "ei", "http://senforsce.com/o8/brain/Some-ID-123", "The Exercise Id created by `programmAddExercise` command")
	flag.StringVar(&exerciseName, "en", "Exercise Squat", "The Name of the Exercise")
	flag.StringVar(&exercicePRBest, "eprb", "100", "The Best PR of the Exercise")
	flag.StringVar(&exercicePRLast, "eprl", "84", "The Last PR of the Exercise")
	flag.StringVar(&exercicePRMin, "eprm", "10", "The Minimum PR of the Exercise")
	flag.StringVar(&exercicePRUnit, "epru", "reps", "The Unit of PR for this Exercise (reps|kg|s|min)")
	flag.StringVar(&exerciceReps, "er", "1", "The number of reps for this exercise")
	flag.StringVar(&exerciceSets, "es", "1", "The number of sets for this exercise")
	flag.StringVar(&exerciceTimeUnit, "etu", "s", "the time unit of the exercise (s|min)")
	flag.StringVar(&exerciceWeight, "ew", "0", "the weight associated with is exercise")
	flag.StringVar(&exerciceWeightUnit, "ewu", "kg", "the weight unit associated with is exercise (g|kg)")
	flag.StringVar(&exerciseTime, "et", "0", "the time associated with is exercise")
	flag.StringVar(&hasSportMovement, "sm", "http://senforsce.com/o8/brain/Movement-Some-ID-123", "the predefined sport movement associated with this exercise")

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
	fmt.Printf("the exerciseName  is %s\n", exerciseName)
	fmt.Printf("the exercicePRBest is %s\n", exercicePRBest)
	fmt.Printf("the exercicePRLast is %s\n", exercicePRLast)
	fmt.Printf("the exercicePRMin is %s\n", exercicePRMin)
	fmt.Printf("the exercicePRUnit is %s\n", exercicePRUnit)
	fmt.Printf("the exerciceReps is %s\n", exerciceReps)
	fmt.Printf("the exerciceSets is %s\n", exerciceSets)
	fmt.Printf("the exerciceTimeUnit is %s\n", exerciceTimeUnit)
	fmt.Printf("the exerciceWeight is %s\n", exerciceWeight)
	fmt.Printf("the exerciceWeightUnit is %s\n", exerciceWeightUnit)
	fmt.Printf("the exerciseTime is %s\n", exerciseTime)
	fmt.Printf("the hasSportMovement is %s\n", hasSportMovement)

	mainInfoId, err := gonanoid.New()

	if err != nil {
		log.Fatalf("Unable to Generate A mainInfoId")
	}

	infoIdentifier := fmt.Sprintf("http://senforsce.com/o8/brain/Exercise-%s", mainInfoId)

	fmt.Printf("the infoIdentifier  is %s\n", infoIdentifier)
	userConfig := userconfig.NewUserConfig()

	repo, err := sparql.NewRepo(userConfig.SparQlUrl,
		sparql.DigestAuth(userConfig.SparQlDigestAuthUsername, userConfig.SparQlDigestAuthPassword),
		sparql.Timeout(time.Millisecond*time.Duration(userConfig.SparQlTimeoutMillisecond)),
	)

	if err != nil {
		log.Fatal(err)
	}

	formattedQuery := fmt.Sprintf(editProgramExercise,
		exerciseName,
		exercicePRBest,
		exercicePRLast,
		exercicePRMin,
		exercicePRUnit,
		exerciceReps,
		exerciceSets,
		exerciceTimeUnit,
		exerciceWeight,
		exerciceWeightUnit,
		exerciseTime,
		hasSportMovement,
		exerciceId,
	)

	log.Println(formattedQuery)

	r, err := repo.Update(formattedQuery)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(r)

	fmt.Printf("\n SUCCESSFULLY EDITED : %s\n", exerciceId)

}
