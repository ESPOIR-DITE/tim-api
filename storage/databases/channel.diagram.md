```mermaid
classDiagram
    
    

    ChannelVideo --> video
    ChannelVideo --> Channel 
    
    Channel --> ChannelType
    Channel --> Account

    channelSubscription --> Channel
    channelSubscription --> Account
    
```