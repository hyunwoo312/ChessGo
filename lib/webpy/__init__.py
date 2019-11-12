from flask import Flask
from flask_sqlalchemy import SQLAlchemy
from flask_bcrypt import Bcrypt
from flask_login import LoginManager
#from sqlalchemy.ext.declarative import declarative_base
#from sqlalchemy import create_engine, Column, String, Integer
app = Flask(__name__)
app.config['SECRET_KEY'] = '420a551128b9ae168b490771ecd1ed2b849a355a'
app.config['SQLALCHEMY_DATABASE_URI'] = 'sqlite:///protected/data/uwu.db'
# engine = create_engine('sqlite:////uwu.db', echo=True)
# basemodel = declarative_base()
db = SQLAlchemy(app)
bcrypt = Bcrypt(app)
login = LoginManager(app)
login.login_view = 'login'
login.login_message_category = 'info' 

from webpy import routes