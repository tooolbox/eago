
 /*
  * This file generated by `build.js` script
  */
 
package bootstrap
 
import (
  "github.com/robertkrimen/otto/ast"
  "github.com/robertkrimen/otto/parser"
)
 
// TestBootstrapProgram :
var TestBootstrapProgram *ast.Program
 
func init() {
  TestBootstrapProgram, _ = parser.ParseFile(nil, "", testBootstrapJS, 0)
}
 
var testBootstrapJS string = `
  var describe = (function () {
  function getAssertFunction(testName) {
    return function (cond) {
      if (!cond) {
        console.error("TEST FAILED ", testName);
      } else {
        console.log("TEST PASSED ", testName);
      }
    };
  }

  function getTestFunction(desc) {
    return function testFunction(description, cb) {
      var assert = getAssertFunction(desc + "/" + description);
      cb(assert);
    };
  }

  return function (description, cb) {
    cb(getTestFunction(description));
  };
})();

`
  
