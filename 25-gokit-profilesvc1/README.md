## sample http serve JSON response for call
 `/uppercase`  to transform lowcase string to uppercase string

    
    $curl -X POST -d '{"S": "high level"}' http://localhost:8080/uppercase

    {"v":"HIGH LEVEL"}
    
 `/count`  to count char in string input 

    
    $curl -X POST -d '{"S": "high level"}' http://localhost:8080/count

    {"v":10}
    
