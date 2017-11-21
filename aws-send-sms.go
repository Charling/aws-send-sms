package main
import "fmt"
import "os"
import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/service/sns"
import "github.com/pborman/getopt"

func main() {
    optMessage := getopt.StringLong("message", 'm', "", "Message to send")
    optPhoneNumber := getopt.StringLong("phonenumber", 'p', "", "Phonenumber")
    optHelp := getopt.BoolLong("help", 'h', "", "Help")
    getopt.Parse()


    if (*optHelp || len(*optMessage) == 0 || len(*optPhoneNumber) == 0) {
        fmt.Println("Sends an SMS via AWS SNS SMS service")
        fmt.Println("Region eu-west-1 is used (hardcoded)")
        fmt.Println("AWS credentials can be passed via")
        fmt.Println("      file ~/.aws/credentials")
        fmt.Println("      environment variables AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY")
        getopt.Usage()
        os.Exit(0)
    }

    svc := sns.New(session.New(&aws.Config{Region: aws.String("eu-west-1")}))

    params := &sns.PublishInput{
        Message: aws.String(*optMessage), // Required
        PhoneNumber:      aws.String(*optPhoneNumber),
    }
    resp, err := svc.Publish(params)

    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
    fmt.Println("success")
    fmt.Println(resp)
    os.Exit(0)
}
