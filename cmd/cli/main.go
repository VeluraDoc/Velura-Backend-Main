// Copyright 2025 Velura Team
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

package main

import (
	"flag"
	"fmt"

	"github.com/VeluraOpenSource/Velura_Documents_Service/internal/config"
	pdf_strategy "github.com/VeluraOpenSource/Velura_Documents_Service/internal/module/doc/usecase"
	doc_strategy "github.com/VeluraOpenSource/Velura_Documents_Service/internal/module/doc/usecase/strategy"
)

var (
	inputFlag    = flag.String("i", "", "input file as PDF")
	strategyFlag = flag.String("s", "pdf", "conversion strategy (pdf)")
)

func main() {
	flag.Parse()
	config.LoadEnv()

	var strategy pdf_strategy.Strategy

	switch *strategyFlag {
	case "pdf":
		strategy = &doc_strategy.PdfToDocx{}

	default:
		strategy = &doc_strategy.PdfToDocx{}
	}

	if err := pdf_strategy.Convertor(strategy, *inputFlag); err != nil {
		fmt.Println("Something went wrong: ", err.Error())
	}
}
