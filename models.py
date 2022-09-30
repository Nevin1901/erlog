from dataclasses import dataclass
import string
from sqlalchemy import Column, Integer, String

from database import Base


@dataclass
class Log(Base):
    id: int
    text: string
    __tablename__ = "Log"

    id = Column(Integer, primary_key=True)
    text = Column(String, nullable=False)

    def __init__(self, text):
        self.text = text

    def __repr__(self):
        return f"<Log {self.text}>"
