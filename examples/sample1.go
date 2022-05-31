package examples

/**
 * @apiDefine token Admins with token access only
 *
 * the token always send with Headers <code>authorization</code> and <code>Bearer</code> in Token String
 */

/**
* @api {get} /api/v1/admins/user/:id get user profile by admin
* @apiPermission token
* @apiGroup Admins
* @apiDescription get user profile by admin token
*
*
* only the <code>admin with token</code> can get it
* @apiHeader {String} authorization Users unique token <code>"Bearer XXXXXXXXXXXXX"</code>.
* @apiHeaderExample {json} Header-Example:
*     {
*       "authorization": "Bearer eyJhbGciOiJIUzI1..."
*     }
* @apiParam {String} id id of user (userId)
* @apiSuccess {Boolean} success true.
* @apiSuccess {Object} user Return user information.
* @apiSuccessExample {json} Success
* HTTP/1.1 200 ok
* {
*   "success": true,
*   "user": {...}
*
* @apiError {Boolean} success false.
* @apiError {String} error message
* @apiErrorExample {json} Error-Response:
*     HTTP/1.1 401 unautorized
*  {
*       "success": false,
*       "error": "Unautorized"
*   }
* @apiErrorExample {json} Error-Response:
*     HTTP/1.1 404 notFound
*  {
*       "success": false,
*       "error": "Invalid ObjectId"
*   }
* @apiErrorExample {json} Error-Response:
*     HTTP/1.1 400 NotFound
*  {
*       "success": false,
*       "error": "User Not Found"
*   }
 */
func sample() {

}
