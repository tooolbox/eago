
 /*
  * This file generated by `build.js` script
  */
 
package bootstrap
 
import (
  "github.com/robertkrimen/otto/ast"
  "github.com/robertkrimen/otto/parser"
)
 
// EagoBootstrapProgram :
var EagoBootstrapProgram *ast.Program
 
func init() {
  EagoBootstrapProgram, _ = parser.ParseFile(nil, "", eagoBootstrapJS, 0)
}
 
var eagoBootstrapJS string = `
  var println = function println() {
  print.apply(void 0, arguments);
  print("\n");
};
var require = (function () {
  function require(str) {
    if (str.substr(0, 2) == "./") {
      str = __dirname + str.substr(2);
    }
    return __require(str);
  }
  return require;
})();

`
  
