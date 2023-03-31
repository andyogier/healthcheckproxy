# healthcheckproxy

Healthcheckproxy will receive the healthcheck call from your traffic manager health probe and forward that call to the healthcheck app. It will then return the response back to the traffic manager health probe ensuring that your app meets the required healthy state you set in the query and that connectivity all the way to the app is successful.
To use the healthcheckproxy, perform the following steps:

1. Import package "github.com/andyogier/healthcheckproxy"
2. Call healthcheckproxy.Start(path string, port string, appName string, query string) and pass in the following parameters:
   1. path is the path your traffic manager will target to call healthcheck. Must include leading /. Example: /health.
   2. port is the port you want to use for the connection.
   3. appName is the name of your app. This must match the name you set in your deployment yaml file.
   4. query is the query you want to use for determining the health of your app. You can use min={num} or percentageHealthyPods={num}. You then need to follow that with &namespace={namespace your app is in}. An example query would be min=1&namespace=myappsnamespace or percentageHealthyPods=80&namespace=myappsnamespace.
