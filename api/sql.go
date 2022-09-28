package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"

	_ "github.com/mattn/go-sqlite3"

	"localhost/bplusbot/pkg/entity"
	"localhost/bplusbot/pkg/repository"
)

const fileName = "shoplist.db"

var service string

func init() {
	fmt.Println("starting....")
	service = "cmdShopList"
}

func main() {
	os.Remove(fileName)

	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("Service: ", service).
			Msgf("Cannot start %s", service)
	}

	shopListRepository := repository.NewSQLiteRepository(db)

	if err := shopListRepository.Migrate(); err != nil {
		log.Fatal().
			Err(err).
			Str("Service: ", service).
			Msgf("Cannot start %s", service)
	}

	gosamples := entity.ShopList{
		Description: "GOSAMPLES",
		Quantity:    2,
	}

	createdGosamples, err := shopListRepository.Create(gosamples)

	if err != nil {
		log.Fatal().
			Err(err).
			Str("Service: ", service).
			Msgf("Cannot start %s", service)
	}

	gotGosamples, err := shopListRepository.GetByName("GOSAMPLES")

	if err != nil {
		log.Fatal().
			Err(err).
			Str("Service: ", service).
			Msgf("Cannot start %s", service)
	}

	fmt.Printf("get by desciption: %+v\n", gotGosamples)

	createdGosamples.ID = 1

	if _, err := shopListRepository.Update(createdGosamples.ID, *createdGosamples); err != nil {
		log.Fatal().
			Err(err).
			Str("Service: ", service).
			Msgf("Cannot start %s", service)
	}

	all, err := shopListRepository.All()

	if err != nil {
		log.Fatal().
			Err(err).
			Str("Service: ", service).
			Msgf("Cannot start %s", service)
	}

	fmt.Printf("\nAll websites:\n")

	for _, website := range all {
		fmt.Printf("website: %+v\n", website)
	}

	if err := shopListRepository.Delete(createdGosamples.ID); err != nil {
		log.Fatal().
			Err(err).
			Str("Service: ", service).
			Msgf("Cannot start %s", service)
	}

	all, err = shopListRepository.All()

	if err != nil {
		log.Fatal().
			Err(err).
			Str("Service: ", service).
			Msgf("Cannot start %s", service)
	}

	fmt.Printf("\nAll websites:\n")

	for _, website := range all {
		fmt.Printf("website: %+v\n", website)
	}
}
