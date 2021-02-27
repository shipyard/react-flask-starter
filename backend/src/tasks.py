from celery import Celery

from src.models import Counter
from src.settings import Settings as S


def make_celery(app):
    # Create and configure the celery app
    celery = Celery(
        app.import_name,
        backend=app.config['CELERY_RESULT_BACKEND'],
        broker=app.config['CELERY_BROKER_URL']
    )
    celery.conf.update(**S.CELERY_CONFIG)

    # Subclass the base Task so that tasks run in a Flask app context
    TaskBase = celery.Task
    class ContextTask(TaskBase):  # noqa
        def __call__(self, *args, **kwargs):  # noqa
            with app.app_context():
                return TaskBase.__call__(self, *args, **kwargs)
    celery.Task = ContextTask

    # Add the tasks
    Tasks.bind(celery)

    return celery


class Tasks:

    # How to add new tasks:
    # - Create the static method (like `ping_once()`)
    # - Add the corresponding bind in the bind() method

    @classmethod
    def bind(cls, celery):
        """Bind the Celery app to all the tasks"""
        # NOTE: DONT
        cls.ping_once = celery.task(cls.ping_once)

    @staticmethod
    def ping_once(amount):
        for counter in Counter.list():
            counter.increment(amount=amount)
