// Copyright © 2017 Marcus McCurdy <marcus.mccurdy@gmail.com>
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
	"github.com/spf13/cobra"
	"github.com/volker48/touchstone/metrics"
	"log"
)

// r2Cmd represents the r2 command
var r2Cmd = &cobra.Command{
	Use:   "r2 y yhat",
	Short: "Coefficient of determination",
	Long: `Calculates the coefficient of determination also known as R^2:
Expects two files as arguments. A file of actual values as the first argument, and a file of predictions as the second argument.
The label is expected to be in the first column.
`,
	Run: func(cmd *cobra.Command, args []string) {
		residuals := &metrics.Residuals{}
		readFiles(args, residuals)

		log.Println("Number of samples: ", residuals.Count)
		log.Println("Mean y: ", residuals.Sum/float64(residuals.Count))
		log.Println("Mean Squared Error: ", residuals.MSE())
		log.Printf("R Squared: %f", residuals.RSquared())
	},
}

func init() {
	RootCmd.AddCommand(r2Cmd)

}
