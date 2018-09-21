// Copyright © 2018 Thrashing Code Team <thrashingcode@thrashingcode.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/Pallinder/go-randomdata"
	"github.com/spf13/cobra"
	"time"
	"github.com/icrowley/fake"
	"github.com/brianvoe/gofakeit"
)

// delugeCmd represents the deluge command
var delugeCmd = &cobra.Command{
	Use:   "deluge",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting.")

		repeats := 1000000

		randLibStart := time.Now()
		fmt.Println(randLibStart.Format(time.Stamp))
		for i := 0; i < repeats; i++ {
			randomdataLib()
		}
		randLibend := time.Now()
		randLibelapsed := randLibend.Sub(randLibStart)
		fmt.Printf("End randomdata Time: %s, %d elapsed.\n", randLibend.Format(time.Stamp), randLibelapsed.Nanoseconds())

		fakeLibStart := time.Now()
		fmt.Println(fakeLibStart.Format(time.Stamp))
		for i := 0; i < repeats; i++ {
			fakeLib()
		}
		fakeLibend := time.Now()
		fakeLibelapsed := fakeLibend.Sub(fakeLibStart)
		fmt.Printf("End fakelib Time: %s, %d elapsed.\n", fakeLibend.Format(time.Stamp), fakeLibelapsed.Nanoseconds())

		gofakeitStart := time.Now()
		fmt.Println(gofakeitStart.Format(time.Stamp))
		for i := 0; i < repeats; i++ {
			gofakeitLib()
		}
		gofakeitend := time.Now()
		gofakeitelapsed := gofakeitend.Sub(gofakeitStart)
		fmt.Printf("End gofakeite Time: %s, %d elapsed.\n", gofakeitend.Format(time.Stamp), gofakeitelapsed.Nanoseconds())


	},
}

func gofakeitLib() {
	gofakeit.Name()
	gofakeit.Email()
	gofakeit.Country()
	gofakeit.City()
	gofakeit.State()
	gofakeit.Street()
	gofakeit.StreetPrefix()
	gofakeit.StreetName()
	gofakeit.StreetNumber()
}

func fakeLib() {
	fake.FullName()
	fake.EmailAddress()
	fake.Country()
	fake.City()
	fake.State()
	fake.Street()
	fake.StreetAddress()
}

func randomdataLib() {
	randomdata.FullName(randomdata.RandomGender)
	randomdata.Email()
	randomdata.Country(randomdata.FullCountry)
	randomdata.City()
	randomdata.State(randomdata.Large)
	randomdata.Street()
	randomdata.Address()
}

func init() {
	rootCmd.AddCommand(delugeCmd)
}
