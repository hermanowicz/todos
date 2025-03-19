import secrets
import sentry_sdk

from bottle import run, route, request
from itsdangerous import URLSafeSerializer
from upstash_redis import Redis

# redis setup for upstach
# redis = Redis.from_env()

# vars
ss_secret_key = secrets.token_urlsafe(48)

# setting up itsdangerous
ss = URLSafeSerializer(ss_secret_key, 'auth')


@route('/')
def index():
    return "hello, world -> index page"


@route("/register")
def register():
    if request.method == "GET":
        return "get register page"
    if request.method == "POST":
        return "form submit handler"


@route("/login")
def login():
    if request.method == "GET":
        return "login form response"
    if request.method == "POST":
        return "login form token sender"


@route("/auth/<token>")
def auth(token: str):
    return f"token from url: {token}"


@route('/privacy-policy')
def privacy_policy():
    return "privacy policy page"


run(server='gunicorn', port=8080, host='0.0.0.0', debug=True, reloader=True)
