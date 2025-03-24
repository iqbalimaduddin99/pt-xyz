package entities

type Admin struct {
   	Master    
    UserName string `db:"user_name"`   
    Password string `db:"password"`   
    FullName string `db:"full_name"`  
}