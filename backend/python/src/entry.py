from src.app import Application


app = Application()
celery = app.celery
flask_app = app.flask_app
