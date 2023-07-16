### Tim-Api

#### Go & Mysql
The Back borne of Timtube System. The Api that handles all the requests from the client, agent and Admin.


## Instructions:
+ This Api depends on a MYSQL Database, Which should be running before running this app.
+ If running for the first time you should execute a http request to create need database tables cfr storage ReadMe.

## Services:
### SQS
Should be created before running the application.

`aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name tim_api_sqs_log`

### S3
First create a Bucket:
This is where videos will be saved.
aws s3api create-bucket --bucket tim-api-videos --endpoint-url http://localhost:4566 --create-bucket-configuration LocationConstraint=eu-west-2