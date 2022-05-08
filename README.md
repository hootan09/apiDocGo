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
- [ ] parse @apiPermission
- [ ] parse @apiQuery
- [ ] parse @apiSampleRequest
- [ ] parse @apiSuccessExample
- [ ] parse @apiSuccess