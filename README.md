# S3 Bucket notifications

Me learning about s3 bucket notifications

## Approaches

### Native S3 notifications

- Offer the least amount of latency
- Less flexible in terms of filtering than the other solutions
- You cannot have multiple notifications with overlapping prefixes

### CloudTrail + EventBridge

- Higher latency than the native S3 notifications
- Much better filtering capabilities, though the suffix filtering is not allowed (not a feature in EventBridge)
- Higher costs than the native S3 notifications, you have to enable data plane events for a given bucket

### Native S3 + Lambda + EventBridge

- Combination of the previous 2 approaches
- Gives you the ability to filter on suffix since you are using the native S3 notifications
- Highest latency out of the previous solutions

## Learnings

### Circular reference with lambda function

Since the notification configuration on the s3 bucket is not a separate resource,
you will end up creating an circular reference when deploying the s3 + notifications + lambda stack (if you use references).

There are three solutions to that problem

#### Two step deployment

This is the approach I would favour. Yes, you have to do two step deployment and this is a bit awkward. One huge benefit though is that **you do not have to name your bucket explicitly**.

#### Naming your bucket explicitly

To avoid references to the bucket, you could name your bucket explicitly and create the arn
yourself. This probably is a good idea if you do not share aws account with other teams.

#### Using custom resource

This is an alternative to specifying the bucket name and still deploying in one go.
Here, we are doing the work _CloudFormation_ should have done, and move the notification settings creation into our own custom resource.

Here is a blog post explaining this approach
https://aws.amazon.com/blogs/mt/resolving-circular-dependency-in-provisioning-of-amazon-s3-buckets-with-aws-lambda-event-notifications/
