from flask import (
    Blueprint,
    jsonify
)

from src.models import Counter

import datetime


blueprint = Blueprint('counter', __name__)

RESPONSE_TEMPLATE = 'Flask server running on port 8080. Pinged {count} {times}, \
most recently on {date}.'


@blueprint.route('/')
def index():
    response = 'Flask server'
    return jsonify(response=response)


@blueprint.route('/api/v1/reset/')
def reset():
    counter = Counter.get_create(label='Test')
    counter.reset()
    return jsonify(response=counter.count)


@blueprint.route('/api/v1/')
def api():
    counter = Counter.get_create(label='Test')
    counter.increment()
    date = datetime.datetime.now()
    dateStr = date.strftime('%c')
    times = 'time' if counter.count == 1 else 'times'
    response = RESPONSE_TEMPLATE.format(
        count=counter.count,
        times=times,
        date=dateStr)
    return jsonify(response=response)
