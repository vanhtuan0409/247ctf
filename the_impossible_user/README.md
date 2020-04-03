- Code

```
from Crypto.Cipher import AES
from flask import Flask, request
from secret import flag, aes_key, secret_key

app = Flask(__name__)
app.config['SECRET_KEY'] = secret_key
app.config['DEBUG'] = False
flag_user = 'impossible_flag_user'

class AESCipher():
    def __init__(self):
        self.key = aes_key
        self.cipher = AES.new(self.key, AES.MODE_ECB)
        self.pad = lambda s: s + (AES.block_size - len(s) % AES.block_size) * chr(AES.block_size - len(s) % AES.block_size)
        self.unpad = lambda s: s[:-ord(s[len(s) - 1:])]

    def encrypt(self, plaintext):
        return self.cipher.encrypt(self.pad(plaintext)).encode('hex')

    def decrypt(self, encrypted):
        return self.unpad(self.cipher.decrypt(encrypted.decode('hex')))

@app.route("/")
def main():
    return "
%s
" % open(__file__).read()

@app.route("/encrypt")
def encrypt():
    try:
        user = request.args.get('user').decode('hex')
        if user == flag_user:
            return 'No cheating!'
        return AESCipher().encrypt(user)
    except:
        return 'Something went wrong!'

@app.route("/get_flag")
def get_flag():
    try:
        if AESCipher().decrypt(request.args.get('user')) == flag_user:
            return flag
        else:
            return 'Invalid user!'
    except:
        return 'Something went wrong!'

if __name__ == "__main__":
  app.run()
```

- AES block size is 128 bits = 16 bytes
- Using [converter](http://string-functions.com/string-hex.aspx), we know that `impossible_flag_user` is 20 bytes length ( = 20 * 8 = 160 bit )
  - Meaning that to encrypt `impossible_flag_user`, we need to pad the string to 3 block (128 * 2 = 256 bit)
- We can pad `impossible_flag_user` beginning with random 2 block (32 bytes)
  - `encrypted = AES("AAAAAA....A" + "hex(impossible_flag_user)")`
- Then we can get the last 256 bit (32 bytes) of the encrypted data as key
