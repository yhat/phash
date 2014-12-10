#!/usr/bin/env python

import collections
import random
import subprocess as sp

with open("passwords.txt", "r") as f:
    passwords = f.read().strip().split("\n")

with open("hashes.txt", "r") as f:
    hashes = f.read().strip().split("\n")

print "[+] Downloading and compiling test dependencies"
sp.check_output(["npm", "install", "password-hash"], stderr=sp.STDOUT)
sp.check_output(["go", "get", "code.google.com/p/go.crypto/ripemd160"])
sp.check_output(["go", "get", "code.google.com/p/go.crypto/md4"])
sp.check_output(["go", "build", "generate.go"])
sp.check_output(["go", "build", "verify.go"])

def go_to_js(password, alg, saltn, i):
    h = sp.check_output(["./generate", password, alg, str(saltn), str(i)]).strip()
    sp.check_output(["node", "verify.js", password, h])

def js_to_go(password, alg, saltn, i):
    h = sp.check_output(["node", "generate.js", password, alg, str(saltn), str(i)]).strip()
    sp.check_output(["./verify", password, h])

def get_args():
    return [random.choice(hashes), random.randint(1, 10), random.randint(1, 10)]

hashes_tested = collections.defaultdict(int)

print "[+] Testing hashes (this is going to take a sec)"
for p in passwords:
    (h, s, i) = get_args()
    go_to_js(p, h, s, i)
    js_to_go(p, h, s, i)
    hashes_tested[h] += 1

print "[+] Tests passed"
print "[+] Hashes successfully run:"
for k, v in hashes_tested.items():
    print "%s: %d" % (k, v,)
