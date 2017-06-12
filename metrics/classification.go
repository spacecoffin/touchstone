// Copyright © 2017 Marcus McCudy <marcus.mccurdy@gmail.com>
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

package metrics

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type ConfusionMatrix struct {
	TP    int64
	FP    int64
	TN    int64
	FN    int64
	Total int64
}

func (cm *ConfusionMatrix) Update(yText, yHatText, yWeightText string) {
	y, err := strconv.ParseInt(yText, 10, 8)
	if err != nil {
		log.Fatal("Error parsing int", err)
	}
	yHat, err := strconv.ParseInt(yHatText, 10, 8)
	if err != nil {
		log.Fatal("Error parsing int", err)
	}

	var yWeight int64
	invalidWeight := strings.Contains(yWeightText, "|")
	if !invalidWeight {
		yWeight, err = strconv.ParseInt(yWeightText, 10, 64)
		if err != nil {
			log.Fatal("Error parsing int", err)
		}
	} else {
		yWeight = 1
	}

	if y == -1 {
		y = 0
	}
	if yHat == -1 {
		yHat = 0
	}
	switch y {
	case 0:
		switch yHat {
		case 0:
			cm.TN += 1 * yWeight
		case 1:
			cm.FP += 1 * yWeight
		}
	case 1:
		switch yHat {
		case 0:
			cm.FN += 1 * yWeight
		case 1:
			cm.TP += 1 * yWeight
		}

	}
	cm.Total++
}

func (cm *ConfusionMatrix) FScore(beta float64) float64 {
	p := cm.Precision()
	r := cm.Recall()
	betaSquared := beta * beta
	f1 := (1 + betaSquared) * (p * r / ((betaSquared * p) + r))
	return f1
}

func (cm *ConfusionMatrix) Precision() float64 {
	return float64(cm.TP) / float64(cm.TP+cm.FP)
}

func (cm *ConfusionMatrix) Recall() float64 {
	return float64(cm.TP) / float64(cm.TP+cm.FN)
}
func (cm *ConfusionMatrix) MCC() float64 {
	denom := float64(cm.TP+cm.FP) * float64(cm.TP+cm.FN) * float64(cm.TN+cm.FP) * float64(cm.TN+cm.FN)
	if denom == 0.0 {
		return 0.0
	}
	numerator := float64(cm.TP*cm.TN) - float64(cm.FP*cm.FN)
	mcc := numerator / math.Sqrt(denom)
	return mcc
}
