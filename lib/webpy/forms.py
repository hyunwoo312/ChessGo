from flask_wtf import FlaskForm
from flask_wtf.file import FileField, FileAllowed
from wtforms import StringField, PasswordField, SubmitField, BooleanField, TextAreaField
from wtforms.validators import DataRequired, Length, Email, EqualTo, ValidationError
from flaskblog.models import User
from flask_login import current_user


class RegistrationForm(FlaskForm):
    username = StringField('Username',
                           validators=[DataRequired(), Length(min=8, max=20)])
    email = StringField('Email', validators=[DataRequired(), Email()])
    password = PasswordField('Password',
                             validators=[DataRequired(), Length(min=8, max=20)])
    confirm_password = PasswordField('Confirm Password',
                                     validators=[DataRequired(), Length(min=8, max=20), EqualTo('password')])
    submit = SubmitField('Sign Up')
    
    def validate_username(self, username):
        user = User.query.filter_by(username=username.data).first()
        if user:
            raise ValidationError('That username is taken uwu.')
    
    def validate_email(self, email):
        user = User.query.filter_by(email=email.data).first()
        if user:
            raise ValidationError('That email is taken uwu.')


class LoginForm(FlaskForm):
    email = StringField('Email', validators=[DataRequired(), Email()])
    password = PasswordField('Password',
                             validators=[DataRequired(), Length(min=8, max=20)])
    remember = BooleanField('Remember Me')
    submit = SubmitField('Login')

class UpdateNameForm(FlaskForm):
    username = StringField('ユーザー名',
                           validators=[DataRequired(), Length(min=8, max=20)])
    submit = SubmitField('アップデイト')
    
    def validate_username(self, username):
        if username.data != current_user.username:
            user = User.query.filter_by(username=username.data).first()
            if user:
                raise ValidationError('That username is taken uwu.')

class PostForm(FlaskForm):
    title = StringField('Title', validators=[DataRequired()])
    content = TextAreaField('Type below uwu', validators=[DataRequired()])
    submit = SubmitField('Post Now UWU!!')

class RequestResetForm(FlaskForm):
    email = StringField('Email', validators=[DataRequired(), Email()])
    submit = SubmitField('Request Password Reset Kappa')

    def validate_email(self, email):
        user = User.query.filter_by(email=email.data).first()
        if user is None:
            raise ValidationError("that email ain't even in the database bro")

class ResetPasswordForm(FlaskForm):
    password = PasswordField('Password',
                             validators=[DataRequired(), Length(min=8, max=20)])
    confirm_password = PasswordField('Confirm Password',
                                     validators=[DataRequired(), Length(min=8, max=20), EqualTo('password')])
    submit = SubmitField('Request Password Reset Kappa')                                 

class SmartPasswordForm(FlaskForm):
    password = PasswordField('Old! Mot De Passe',
                             validators=[DataRequired(), Length(min=8, max=20)])
    newpassword = PasswordField('New! Mot De Passe',
                             validators=[DataRequired(), Length(min=8, max=20)])
    confirm_password = PasswordField('Confirm! Mot De Passe',
                                     validators=[DataRequired(), Length(min=8, max=20), EqualTo('newpassword')])
    submit = SubmitField('アップデイト')   
    
class UpdatePfpForm(FlaskForm):
    picture = FileField('Update Profile Pic', validators=[FileAllowed(['jpg','jpeg','png','gif'])])
    submit = SubmitField('アップデイト')


