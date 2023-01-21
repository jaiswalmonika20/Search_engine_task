# **Search-Engine-Task**


## **Step:1 Clone the Repository-**

    git clone https://github.com/jaiswalmonika20/Search_engine_task  


## **Step:2 Running Dokcer-**
   (Check docker running properly)
  
    cd "Search_engine_task"

    docker-compose up


## **Step:3 Testing API-**

 - Use any API testing applications such as Postman, Insomnia or ThunderClinet (VS Code extension)
 - All available Routes
 
 
    i. To check api is working properly : **GET** REQUEST 
    
         http://localhost:8080/v1/
         
    ii. To store webpage in MongoDB : **POST** REQUEST 
    
        http://localhost:8080/v1/newpage
        
      Add this  to request body
        
        {
         "id":1,
         "key":"ford car sheldon "
        }
        
    iii. To get Query in MongoDB : **GET** REQUEST 
    
        http://localhost:8080/v1/ford car
        
        
    iv. To get all web pages in MongoDB : **GET** REQUEST  
    
        http://localhost:8080/v1/pages
        
       
       
## Step:4 **To stop running containers-**

      docker-compose down 
        
        
    
 
