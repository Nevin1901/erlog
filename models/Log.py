from ..app import db


class Log(db.Model):
    id = db.Column(db.Integer)
    text = db.Column(db.String, nullable=False)
