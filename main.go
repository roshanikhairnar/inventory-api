/*   file name : main.go
The Initialize method is responsible for create a database connection and wire up the routes,
and the Run method will simply start the application.
*/
package main

func main() {
	coditation := Coditation{}
	coditation.Initialize("root", "qwerty", "coditation")
	coditation.Run(":8080")
}
