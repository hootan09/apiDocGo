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
- [X] parse @apiParam (Must be Fixed in iterate map)
- [-] parse @apiErrorExample
- [X] parse @apiError
- [-] parse @apiHeaderExample
- [X] parse @apiHeader
- [-] parse @apiParamExample
- [X] parse @apiUse
- [X] parse @apiPermission
- [X] parse @apiQuery
- [X] parse @apiSampleRequest
- [X] parse @apiSuccessExample
- [X] parse @apiSuccess
- [X] parse @apiVersion
- [X] Reader Functionality
- [ ] Writer Functionality
- [ ] ApiDOcGo yml config file
    - [ ] Get cwd as path and looking for apidocgo.yml
- [ ] Using go embed for saving static file in binary build
- [ ] Using go flag for some things like serving - (builded content) | source path | destination path
    - [ ] Example `"$ apidocgo -i routes/ -o public/apidoc"`
    - [ ] Example `"$ apidocgo init"` create apidocgo.yml in cwd
    - [ ] Example `"$ apidocgo server -p 8080"` Small Server for serving statics file in port and showing in console `"server running at http://localhost:port"`