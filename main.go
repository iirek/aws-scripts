package main

import (
	"fmt"
	"os"

	_ "github.com/aws/aws-sdk-go/aws"
	_ "github.com/aws/aws-sdk-go/aws/awserr"
	_ "github.com/aws/aws-sdk-go/aws/request"
	_ "github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/aws/aws-sdk-go/service/elasticache"
)

func main() {
	verifyEnvVariables()
}

func verifyEnvVariables() {
	variables := []string {"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_SESSION_TOKEN"}

	for _, variable := range variables {
		fmt.Printf("VARIABLE :%s is: %s", variable , os.Getenv(variable))
	}

}
