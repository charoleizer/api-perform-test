from flask import Flask
app = Flask(__name__)

import logging
log = logging.getLogger('werkzeug')
log.setLevel(logging.ERROR)
app.logger.disabled = True
log.disabled = True

@app.route("/ping")
def ping():
    return "<p>pong</p>"


if __name__ == '__main__':
      app.run(host='localhost', port=7771)
