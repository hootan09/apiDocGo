### Built in Code Documention
```sh
$ godoc -http=127.0.0.1:6060 
#or just the port 
$ godoc -http=:6060
```

Todo:
- [X] Parse @api
- [X] Parse @apiBody
- [X] Parse @apiDefine
- [X] Parse @apiExample
- [X] parse @apiGroup
- [X] parse @apiName
- [-] parse @apiParam (Must be Fixed in iterate map)
- [-] parse @apiErrorExample
- [X] parse @apiError
- [-] parse @apiHeaderExample
- [-] parse @apiHeader (must accept multiple fileds)
- [-] parse @apiParamExample
- [X] parse @apiUse
- [X] parse @apiPermission
- [X] parse @apiQuery
- [X] parse @apiSampleRequest
- [X] parse @apiSuccessExample
- [X] parse @apiSuccess
- [X] parse @apiVersion
- [X] Reader Functionality
- [X] Writer Functionality
- [X] ApiDOcGo yml config file
    - [X] Get cwd as path and looking for apidocgo.yml
- [X] Using go embed for saving static file in binary build
- [X] Using go flag for some things like serving - (builded content) | source path | destination path
    - [ ] Example `"$ apidocgo -i routes/ -o public/apidoc"`
    - [X] Example `"$ apidocgo init"` create apidocgo.yml in cwd
    - [X] Example `"$ apidocgo -p 8080"` Small Server for serving statics file in port and showing in console `"Listening on http://localhost:8080"`