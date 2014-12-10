var ph = require('password-hash');

if(process.argv.length < 4) {
  console.error("Usage: node verify.js <password> <hash>")
  process.exit(2);
}

var password = process.argv[2];
var hash = process.argv[3];

if(!ph.verify(password, hash)) {
  console.error(password+" did not match "+hash);
  process.exit(2);
}
