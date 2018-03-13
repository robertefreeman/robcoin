from flask import Flask, render_template, request, make_response, redirect
from flask_cors import CORS, cross_origin

import random, socket, time, json, os, sys, ast

version ='1.0'
hostname = socket.gethostname()

print "Starting web container %s" % hostname

app = Flask(__name__)
CORS(app)

@app.route('/')
def index():
    url = "robcoins.com"

    return render_template('index.html', hostname=hostname, version=version)


#curl -X PUT -H 'Content-Type: application/json' -d '{"healthy": "False"}' http://localhost:5000/health
#curl -X PUT -H 'Content-Type: application/json' -d '{"healthy": "True"}' http://localhost:8000/health
#curl -v http://localhost:8000/health

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=80, debug=True )
