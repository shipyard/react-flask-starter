import uuid

from flask_sqlalchemy import SQLAlchemy
from sqlalchemy.orm.session import make_transient
from sqlalchemy_utils import EncryptedType
from sqlalchemy_utils.types.encrypted.encrypted_type import AesEngine

from src.settings import Settings as S


db = SQLAlchemy()
EncryptedString = EncryptedType(db.Unicode, S.AES_SECRET_KEY, AesEngine, 'pkcs5')


class BaseModel(db.Model):

    __abstract__ = True

    id = db.Column(db.Integer(), primary_key=True)
    uuid = db.Column(db.String(36))

    @classmethod
    def get_create(cls, *args, **kwargs):
        return cls.get(*args, **kwargs) or cls.create(*args, **kwargs)

    @classmethod
    def get(cls, *args, exception=False, **kwargs):
        result = cls.query.filter_by(*args, **kwargs).first()
        if not result and exception:
            raise Exception('No object matches this query!')
        return result

    @classmethod
    def list(cls, *args, **kwargs):
        return cls.query.filter_by(*args, **kwargs) \
                        .order_by('id') \
                        .all()

    @classmethod
    def create(cls, *args, **kwargs):
        instance = cls(*args, **kwargs)
        instance.generate_unique_uuid()
        instance.save()
        return instance

    def generate_unique_uuid(self):
        potential_uuid = str(uuid.uuid4())
        while self.__class__.get(uuid=potential_uuid):
            potential_uuid = str(uuid.uuid4())
        self.uuid = potential_uuid

    def refresh(self):
        db.session.refresh(self)

    def set(self, **kwargs):
        for field, value in kwargs.items():
            setattr(self, field, value)

    def clone(self, **fields):
        copy = self.get(id=self.id)
        make_transient(copy)

        copy.set(id=None, **fields)
        if hasattr(self, 'uuid'):
            self.generate_unique_uuid()

        copy.save()

        return copy

    def save(self):
        db.session.add(self)
        db.session.commit()

    def delete(self):
        db.session.delete(self)
        db.session.commit()
