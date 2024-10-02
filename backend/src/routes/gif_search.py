from flask import (
    Blueprint,
    jsonify
)

import requests


blueprint = Blueprint('backend-go', __name__)

@blueprint.route('/app/v1/gif-search/', methods=['POST'])
def create_serach():
    response = requests.post('http://backend-go:8081/')
    print(response) # DEBUG
    print(response.text) # DEBUG
    return jsonify(response=response.text)
