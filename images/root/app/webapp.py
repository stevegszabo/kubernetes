"""
Sample python webapp container
"""
import socket

from flask import Flask

GAPPLICATION = Flask(__name__)


@GAPPLICATION.route("/")
def indexpage():
    """
    Sample index page
    :return: String
    """
    hostname = socket.gethostname()
    address = socket.gethostbyname(hostname)
    return "{}:{}\n".format(hostname, address)


if __name__ == "__main__":
    GAPPLICATION.run(host='0.0.0.0', port=8080)
