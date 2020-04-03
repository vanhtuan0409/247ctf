- Code

```
import os
from flask import Flask, request, session
from flag import flag

app = Flask(__name__)
app.config['SECRET_KEY'] = os.urandom(24)

def secret_key_to_int(s):
    try:
        secret_key = int(s)
    except ValueError:
        secret_key = 0
    return secret_key

@app.route("/flag")
def index():
    secret_key = secret_key_to_int(request.args['secret_key']) if 'secret_key' in request.args else None
    session['flag'] = flag
    if secret_key == app.config['SECRET_KEY']:
      return session['flag']
    else:
      return "Incorrect secret key!"

@app.route('/')
def source():
    return "
%s
" % open(__file__).read()

if __name__ == "__main__":
    app.run()
```

- After go to `/flag`, your web will be set a cookie on browser
  - Open Inspect tools, go to Application tab, select Cookies
- We can use `base64 -d` to decode the cookie stored on browser
  - `echo $COOKIE | base64 -d`
- Flag will also be encrypted with base64
