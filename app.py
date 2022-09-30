import json
from flask import Flask, request, jsonify
from database import db_session, init_db
from models import Log
from flask_sqlalchemy import SQLAlchemy

db = SQLAlchemy()
app = Flask(__name__)

app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///project.db"

with app.app_context():
    init_db()


@app.route("/")
def index():
    logs = Log.query.all()
    return jsonify(logs)


@app.route("/create")
def create():
    text = request.args.get("text")
    log = Log(text)
    db_session.add(log)
    db_session.commit()
    return "ok"


@app.teardown_appcontext
def shutdown_session(exception=None):
    db_session.remove()
