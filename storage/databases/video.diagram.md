```mermaid
classDiagram
    VideoCategory --> Category
    VideoCategory --> Video
    
    VideoComment --> Comment
    VideoComment --> Video
    
    VideoReaction --> Video
    VideoReaction --> User
    
```