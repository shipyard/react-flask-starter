import uuid
import time
import localstack_client.session

from flask import Blueprint
from flask import jsonify
from flask import redirect
from flask import request
from flask import url_for
from werkzeug.utils import secure_filename


BUCKET = 'test-bucket'

blueprint = Blueprint('aws', __name__)


def s3_client():
    return localstack_client.session.Session().client('s3')


@blueprint.route('/api/v1/files/')
def files():
    s3_client().create_bucket(Bucket=BUCKET)
    return jsonify(s3_client().list_objects(Bucket=BUCKET))


@blueprint.route('/api/v1/files/upload/', methods=['GET', 'POST'])
def upload():
    if request.method == 'POST':
        if 'file' not in request.files:
            return "No 'file' defined.", 400
        file = request.files['file']
        if file:
            if file.filename:
                filename = secure_filename(file.filename)
            else:
                timestr = time.strftime('%Y%m%d%H%M%S')
                ext = file.filename.split('.')[-1]
                # the suffix is necessary to avoid overwriting files when multiple uploads are done quickly
                suffix = ''.join(random.choices(string.ascii_letters, k=7))
                filename = f'upload-{timestr}-{suffix}.{ext}'
            s3_client().put_object(Bucket=BUCKET,
                                   Key=secure_filename(file.filename),
                                   ContentType=file.mimetype,
                                   Body=file)

            return jsonify(message="Uploaded '" + file.filename + "'.")


@blueprint.route('/api/v1/files/create/')
def add_file():
    s3_client().put_object(Bucket=BUCKET,
                           Key='{}.txt'.format(uuid.uuid1()),
                           Body=b'some content')
    return redirect(url_for('.files'))
