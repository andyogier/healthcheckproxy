# Healthcheck Proxy

This is an example of how to setup a Healthcheck Proxy using go. Your app will receive the healthcheck call from your traffic manager health probe and forward that call to the healthcheck app. It will then return the response back to the traffic manager health probe ensuring that your app meets the required healthy state you set in the query and that connectivity all the way to the app is successful.
