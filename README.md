### Built in Code Documention
```sh
$ godoc -http=127.0.0.1:6060 
#or just the port 
$ godoc -http=:6060
```

Todo:
- [X] Parse @api
- [ ] Parse @apiBody
- [X] Parse @apiDefine
- [X] Parse @apiExample
- [X] parse @apiGroup
- [X] parse @apiName
- [X] parse @apiParam (Must be Fixed in iterate map)
- [ ] parse @apiErrorExample
- [ ] parse @apiError
- [ ] parse @apiHeaderExample
- [ ] parse @apiHeader
- [ ] parse @apiParamExample
- [ ] parse @apiUse
- [ ] parse @apiPermission