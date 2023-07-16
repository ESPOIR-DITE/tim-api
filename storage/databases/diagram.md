```mermaid
classDiagram
    
    class userDetail{
        
    }

    class UserAccount{

    }
    class User{

    }
    class Account{

    }
    class Role{

    }
    
    User --> Role
    UserAccount --> Account
    UserAccount --> User
    UserAccount --> userDetail
    UserAccount --> userBank
    
    UserSubscription -->Account
    VideoUser --> Account
 ```