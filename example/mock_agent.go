package main

import (
	"encoding/json"
	"flag"
	"github.com/mdanzinger/policymanager"
	_ "github.com/mdanzinger/policymanager/policies"
	"io/ioutil"
	"log"
)

func main() {
	policyFile := flag.String("github.com/mdanzinger/policymanager_file", "./policy.json", "path to policy file")
	flag.Parse()

	rawPolicy, err := ioutil.ReadFile(*policyFile)
	if err != nil {
		log.Fatal("unable to read policy file")
	}

	pol := policy.DefaultPolicy{}
	if err := json.Unmarshal(rawPolicy, &pol); err != nil {
		log.Fatalf("error unmarshaling policy: %s", err)
	}

	enforcer, err := policy.ResolveEnforcer(pol)
	if err != nil {
		log.Fatalf("error resolving enforcer: %s", err)
	}

	if err := enforcer.Enforce(pol); err != nil {
		log.Fatalf("error enforcing policy: %s", err)
	}
}
