import numpy as np 
import matplotlib.pyplot as plt 
import os, sys
import keras.backend as K 
from PIL import Image
import cv2
import psycopg2
import neat
import time 

from keras.models import Model, Sequential
from keras.layers import Dense, Conv2D, BatchNormalization, GaussianNoise
from keras.layers import Dropout, Activation

class Network:
    def __init__(self):
        self.inputshape = (8, 8)
    
    def build(self):
        pass

    def train(self):
        pass

    def predict(self):
        pass