package main

import (
	// "hospi_bed_stats/bed_stats"
	"hospi_bed_stats/api"

	"github.com/joho/godotenv"
)

func main() {
	// 	bed_num := bed_stats.GetBedStats()
	// 	fmt.Printf("%v beds are available\n", bed_num)
	err := godotenv.Load()
	if err != nil {

	}
	api.Api()
}
