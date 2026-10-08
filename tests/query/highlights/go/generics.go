package main
func Sum[T Number](xs []T) T {
//       ^ @type.parameter
//         ^ @type
	var f = obj.Method
//           ^ @variable.other.member
	return obj.field
//          ^ @variable.other.member
}

type response struct{}

func callGeneric[Response any](callbacks []func(), index int) {
    httpclient.DoJSONObserved[Response](client, req, 0, observer)
//             ^^^^^^^^^^^^^^ @function.method
//                            ^^^^^^^^ @type
    DoJSONObserved[Response](client, req, 0, observer)
//  ^^^^^^^^^^^^^^ @function
//                 ^^^^^^^^ @type
    httpclient.DoJSONObserved[response](client, req, 0, observer)
//             ^^^^^^^^^^^^^^ @function.method
//                            ^^^^^^^^ @type
    httpclient.DoJSONObserved[models.TokenResponse](client, req, 0, observer)
//             ^^^^^^^^^^^^^^ @function.method
//                                   ^^^^^^^^^^^^^ @type
    httpclient.DoJSONObserved[int](client, req, 0, observer)
//             ^^^^^^^^^^^^^^ @function.method
//                            ^^^ @type.builtin
    httpclient.DoJSONObserved[Response, int](client, req, 0, observer)
//             ^^^^^^^^^^^^^^ @function.method
//                            ^^^^^^^^ @type
//                                      ^^^ @type.builtin
    callbacks[index]()
//  ^^^^^^^^^ @variable.parameter
//            ^^^^^ @variable.parameter
    callbacks[0]()
//  ^^^^^^^^^ @variable.parameter
//            ^ @constant.numeric.integer
    _ = values[unknown]
//             ^^^^^^^ @variable
}
