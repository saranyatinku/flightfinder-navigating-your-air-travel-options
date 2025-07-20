package aws

import (
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

var logger = log.New(os.Stdout, "[METRICS] ", log.Ldate|log.Ltime)

// MetricsClient implements MetricsClient interface
type MetricsClient struct {
	svc *cloudwatch.CloudWatch
}

func NewMetricsClient(aws_region string) (*MetricsClient, error) {
	mySession, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	cfg := aws.NewConfig().WithRegion(aws_region)
	svc := cloudwatch.New(mySession, cfg)
	client := &MetricsClient{svc: svc}
	if err = client.checkCloudWatchConnection(); err != nil {
		return nil, err
	}
	return client, nil
}

func (m *MetricsClient) PutRequestMetrics(pattern string, method string, executionTime time.Duration) {
	input := &cloudwatch.PutMetricDataInput{
		Namespace: aws.String("FlightFinder"),
		MetricData: []*cloudwatch.MetricDatum{
			{
				MetricName: aws.String("ReqDuration"),
				Unit:       aws.String("Seconds"), // https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html
				Value:      aws.Float64(float64(executionTime.Seconds())),
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String("URL"),
						Value: aws.String(pattern),
					},
					{
						Name:  aws.String("Method"),
						Value: aws.String(method),
					},
				},
			},
		},
	}

	_, err := m.svc.PutMetricData(input)

	// log on error, eg. Permission Denied - missing cloudwatch:PutMetricData in EC2 Instance Profile (the assigned IAM Role)
	if err != nil {
		logger.Println("Error sending metrics to AWS CloudWatch:", err.Error())
	}
}

func (m *MetricsClient) checkCloudWatchConnection() error {
	// send request to CloudWatch and see if error occurs
	input := &cloudwatch.PutMetricDataInput{
		Namespace: aws.String("FlightFinder"),
		MetricData: []*cloudwatch.MetricDatum{
			{
				MetricName: aws.String("CheckingCloudWatchConnectionOnStartup"),
				Unit:       aws.String("None"), // https://docs.aws.amazon.com/AmazonCloudWatch/latest/APIReference/API_MetricDatum.html
				Value:      aws.Float64(1.0),
			},
		},
	}

	_, err := m.svc.PutMetricData(input)
	return err
}
