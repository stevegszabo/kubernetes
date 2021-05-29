"""
Sample python webapp container
"""
import os
import socket

from flask import Flask

GAPPLICATION = Flask(__name__)


@GAPPLICATION.route("/")
def indexpage():
    """
    Sample index page
    :return: String
    """
    version = os.getenv('WEBAPP_VERSION', 'v0.0.0')
    hostname = socket.gethostname()
    address = socket.gethostbyname(hostname)
    return "{}:{}:{}\n".format(version, hostname, address)


if __name__ == "__main__":
    GAPPLICATION.run(host='0.0.0.0', port=8080)
