var ph = require('password-hash');

if(process.argv.length < 6) {
  console.error("Usage: node generate.js <password> <algorithm> <salt length> <iter>");
  process.exit(2);
}

var password = process.argv[2];
var algorithm = process.argv[3];
var saltLength = parseInt(process.argv[4], 10);
var iter = parseInt(process.argv[5], 10);
console.log(ph.generate(password, {
  algorithm: algorithm,
  saltLength: saltLength,
  iterations: iter   
}));
