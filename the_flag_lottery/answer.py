#!/bin/python2
import random, time, socket


s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
s.connect(("559dede4bc1c6129.247ctf.com", 50404))

print(s.recv(1024))
secret = random.Random()
secret.seed(int(time.time()))
winning_choice = str(secret.random())
s.sendall(winning_choice)
print(s.recv(1024))
print(s.recv(1024))
