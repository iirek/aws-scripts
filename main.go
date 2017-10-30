package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	_ "github.com/aws/aws-sdk-go/aws/awserr"
	_ "github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticache"
)

func main() {
	fmt.Printf("Are variable verified? %s\n", verifyEnvVariables())
	
	sess := session.Must(session.NewSession())

	svc := elasticache.New(sess, aws.NewConfig().WithRegion("eu-west-1"))

	input := &elasticache.DescribeCacheClustersInput{
		CacheClusterId: aws.String("potato-cluster"),
		ShowCacheNodeInfo: aws.Bool(true),
	}

	result, err := svc.DescribeCacheClusters(input)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, cluster_node := range result.CacheClusters[0].CacheNodes {
		fmt.Printf("%v:%v\n", *result.CacheClusters[0].CacheClusterId, *cluster_node.CacheNodeId)
	}

}

func verifyEnvVariables()(isVerified bool)  {
	variables := []string {"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY", "AWS_SESSION_TOKEN"}
	var verifiedLen = 0

	for _, variable := range variables {
		var envVariable = os.Getenv(variable)
		if len (envVariable) > 0 {
			verifiedLen +=1
		}
	}
	isVerified = verifiedLen == len(variables)

	return


}
