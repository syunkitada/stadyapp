http_address = "127.0.0.1:4180"

redirect_url="https://myapp.localhost.test/oauth2/callback"

provider = "github"
email_domains = "*"

# Generate cookie_secret by `python -c 'import os,base64; print(base64.urlsafe_b64encode(os.urandom(32)).decode())'`
cookie_secret = "XXX"

# https://github.com/settings/applications/new
client_id = "XXX"
client_secret = "XXX"

cookie_expire = "1h"
