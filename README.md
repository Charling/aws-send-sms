Sends SMS with Amazon AWS, using eu-west-1 region (hardcoded)

Getting Started on Ubuntu:
1) create AWS API user on eu-west-1 region
   AWS IAM policy must allow SNS:Publish
2) populate ~/.aws/credentials with the api user
   OR
   use environment variables AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY
3) apt-get -y install golang git
   git clone https://github.com/acortiana/aws-send-sms
   cd aws-send-sms
   export GOPATH="aws-send-sms directory absolute path"
   go get
   go build 

   aws-send-sms binary will be created

4) try to send an SMS:
   ./aws-send-sms -m "test message" -p +391234567890
