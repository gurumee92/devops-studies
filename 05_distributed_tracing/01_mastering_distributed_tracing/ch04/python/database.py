from sqlalchemy import create_engine
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
from sqlalchemy.schema import Column
from sqlalchemy.types import String

db_url = "mysql+pymysql://root:root@localhost:3306/test_db"
engine = create_engine(db_url)
Session = sessionmaker(bind=engine)
session = Session()
Base = declarative_base()

class Person(Base):
    __tablename__ = 'people'
    name = Column(String, primary_key=True)
    title = Column(String)
    description = Column(String)

    @staticmethod
    def get(name):
        return session.query(Person).get(name)