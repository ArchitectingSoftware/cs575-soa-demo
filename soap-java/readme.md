##SOAP-Java

Sample project used in my software design class to showcase SOAP web services.  While the SOAP pattern is not used widely anymore, its interesting still to look at its design characteristics.  That is the purpose of this project.

This project uses gradle as its build system:

`gradle build` is used to build the project

`gradle bootRun` is used to run the project

#####Some sample XML Requests

To query a specific publication, change the `GetById` parameter.
```xml
<x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/" 
            xmlns:pub="http://drexel.edu/ws/messages">
    <x:Header/>
    <x:Body>
        <pub:PublicationRequest>
            <pub:RequestType>
                <pub:GetById>3</pub:GetById>
            </pub:RequestType>
        </pub:PublicationRequest>
    </x:Body>
</x:Envelope>
```
To get all publications:
```xml
<x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/" xmlns:pub="http://drexel.edu/ws/messages">
    <x:Header/>
    <x:Body>
        <pub:PublicationRequest>
            <pub:RequestType>
                <pub:GetAll></pub:GetAll>
            </pub:RequestType>
        </pub:PublicationRequest>
    </x:Body>
</x:Envelope>
```

If you use `curl` or some other testing tool make sure you set a content type header - `Content-Type:text/xml`

To get the WSDL document, you can go to this endpoint <http://localhost:8080/ws/publications.wsdl>.  Note you might need to change the default port number.

To execute these requests, via HTTP or CURL, Post to the endpoint <http://localhost:8080/ws/>