from src.models.base import BaseModel, db


class Counter(BaseModel):

    count = db.Column(db.Integer, default=0, nullable=False)
    label = db.Column(db.String(), nullable=False)

    RESET_THRESHOLD = 1000000000

    def increment(self, amount=1):
        self.count += amount
        if self.count > self.RESET_THRESHOLD:
            self.count = 0
        self.save()

    def reset(self):
        self.count = 0
        self.save()
