# Redis pub/sub chatting service

## 1. What is redis pub/sub?
*  Publish/Subscribe messaging paradigm where senders (publishers) are not sending messages to specific receivers (subscribers). Rather, published messages are characterized into channels, without knowledge of what (if any) subscribers there may be. (Reference: https://redis.io/docs/manual/pubsub/)
* Once subscribe to a channel, client became subscriber and will become read-only.(Meaning no commands can be issued by clients)
* More than one channel can be subscribed to at a time.
