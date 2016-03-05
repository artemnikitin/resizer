# resizer
CLI tool for resizing images

##### Get it
``` 
go get github.com/artemnikitin/devicefarm-ci-tool 
``` 

   
Required launch parameters:   
```
devicefarm-ci-tool -project=name -app=/path/to/my/app.apk
```
By default, "BUILTIN_FUZZ" tests will be run for your app.

##### Optional parameters:   
- ```region``` set S3 region, by default region will be set to ```us-west-2```(At this moment, will be set to ```us-west-2``` in any case, because it's only supported region for the moment).          
Example:    
``` 
devicefarm-ci-tool -project=name -app=/path/to/my/app.apk -region=region-name 
```    
- ```devices``` specify name of device pool where app will be run.      
Example:   
``` 
devicefarm-ci-tool -project=name -app=/path/to/my/app.apk -devices=my-device-pool
```   
- ```config``` specify path to config in JSON format.      
Example:   
``` 
devicefarm-ci-tool -project=name -app=/path/to/my/app.apk -config=/path/to/config.json
```   
- ```wait``` will wait for end of run. Disabled by default. Useful for CI.     
Example:   
``` 
devicefarm-ci-tool -project=name -app=/path/to/my/app.apk -wait=true
```   
