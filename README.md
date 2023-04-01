# Healthcheck Proxy

This is an example of how to setup a Healthcheck Proxy using go lang. You can follow this pattern in the language your app is coded in. Your app will receive the health probe from your traffic manager and forward that call to the healthcheck app. The healthcheck app will run your query and send a response back to your app. Your app will then send the response back to the traffic manager health probe ensuring that your app meets the required healthy state you set in the query and that connectivity all the way to the app is successful.
