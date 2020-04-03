#!/bin/python2
import random, time

secret = random.Random()
secret.seed(int(time.time()))
winning_choice = str(secret.random())
print(winning_choice)
