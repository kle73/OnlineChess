"""import socket
import json
import threading

SERVER = 'localhost'
PORT = 5050


client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
client.connect((SERVER, PORT))
ans = client.recv(64).decode('UTF-8')
print(ans)

def receive():
    ans = client.recv(64).decode('UTF-8')
    print(ans)

answer_thread = threading.Thread(target=receive)
answer_thread.start()

while True:
    pos1 = input("Move From: ")
    pos2 = input("Move To: ")
    turn = {"pos1": pos1, "pos2": pos2}
    turn = json.dumps(turn)
    client.send(bytes(turn, 'UTF-8'))
    ans = client.recv(64).decode('UTF-8')
    print(ans)"""


s = "white CHECKMATE"
p = "jddsd CHECK"
list = [s, p]
for i in list:
    if "CHECKMATE" in i:
        print(i)
