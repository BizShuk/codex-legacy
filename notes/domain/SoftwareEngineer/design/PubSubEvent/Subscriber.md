# subscriber

This design is for Event Stream subscriber handler.

SQS client of aws-sdk-go didn't provider listener behavior. So with this desing, this client can handle multiple listener.

ps: Golang http client is safe for concurrent usage.

1. configure sqs arn arn and DLQ arn
2. configure account setting
3. writing handler for one listener
4. writing delete handler for one listener
5. pooling mechanism for listener and handler
6. one channle between listener and handler
7. base one experiment if channel is full, listeners should block

### issue discussion

1. where to put handler and delete handler? and how to attach to listener
2. how to design listener pool and handler pool
3. long polling for listener, Exponential Backoff And Jitter <https://aws.amazon.com/tw/blogs/architecture/exponential-backoff-and-jitter/>
4. retry and age mechanism
