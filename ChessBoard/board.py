import pygame, sys
from time import sleep
from Button import Button
from pieces import Piece

import socket
import json
import threading
SERVER = 'localhost'
PORT = 5050

pygame.init()
WIDTH, HEIGHT = 700, 750
screen = pygame.display.set_mode((WIDTH, HEIGHT))
clock = pygame.time.Clock()
BLACK = (0, 0, 0)
GREY = (210, 210, 210)
WHITE = (255, 255, 255)
Font = pygame.font.SysFont('comicsans', 45)
SQUARE = 80
BUFFER_LEFT = 60
BUFFER_BOTTOM = 110
pygame.display.set_caption("ChessOnline")

img = pygame.image.load("./static/bishop.png")
img = pygame.transform.scale(img, (80, 80))

def get_pieces(buttons):
    pieces = []
    for button in buttons:
        if button.value[1] == "1" or button.value[1] == "2" or button.value[1] == "7" or button.value[1] == "8":
            pieces.append(Piece(img, (button.x, button.y), button.value))
    return pieces


def message_to_screen(msg, color, position, Font):
    text = Font.render(msg, True, color)
    screen.blit(text, position)

def draw_GRID():
    for i in range(1, 9):
            pygame.draw.line(screen, BLACK, (BUFFER_LEFT, i * SQUARE), (WIDTH, i * SQUARE))
            pygame.draw.line(screen, BLACK, (WIDTH - i*SQUARE, 0), (WIDTH - i*SQUARE, HEIGHT - BUFFER_BOTTOM))


def calculate_button_value_white(char, num):
    return chr(64 + 9 - char) + str(9 - num)

def calculate_button_value_black(char, num):
    return chr(64 + char) + str(num)

def get_buttons(calculate_button_value):
    buttons = []
    rest = 0
    for char in range(1, 9):
        count = rest % 2
        for num in range(1, 9):
            x = WIDTH - char*SQUARE
            y = num * SQUARE - SQUARE
            value = calculate_button_value(char, num)
            if num % 2 == count:
                buttons.append(Button(GREY, x, y, SQUARE, SQUARE, value))
            else:
                buttons.append(Button(BLACK, x, y, SQUARE, SQUARE, value))
        rest += 1
    return buttons

def draw_nums_white_view():
    for i in range(1, 9):
        message_to_screen(str(i), BLACK, (25, HEIGHT - i * SQUARE - 70), Font)

def draw_chars_white_view():
    for i in range(1, 9):
        message_to_screen(chr(64 + i), BLACK, (i*SQUARE + 10, HEIGHT - 90), Font)

def draw_nums_black_view():
    for i in range(1, 9):
        message_to_screen(str(i), BLACK, (25, i * SQUARE - 50), Font)

def draw_chars_black_view():
    for i in range(1, 9):
        message_to_screen(chr(64 + i), BLACK, (WIDTH - i*SQUARE + 20, HEIGHT - 90), Font)


def receive_first(buttons, pieces):
    ans = client.recv(64).decode('UTF-8')
    ans = json.loads(ans)
    for piece in pieces:
        if piece.value == ans["Pos1"]:
            for button in buttons:
                if button.value == ans["Pos2"]:
                    newPos = (button.x, button.y)
                    piece.pos = newPos
                    piece.value = button.value

def receive(buttons, pieces, turn):
    while True:
        ans = client.recv(64).decode('UTF-8')

        if ans == "SET":
            for piece in pieces:
                if piece.value == turn["pos1"]:
                    for button in buttons:
                        if button.value == turn["pos2"]:
                            newPos = (button.x, button.y)
                            piece.pos = newPos
                            piece.value = button.value
            continue
        try:
            ans = json.loads(ans)
            for piece in pieces:
                if piece.value == ans["Pos1"]:
                    for button in buttons:
                        if button.value == ans["Pos2"]:
                            newPos = (button.x, button.y)
                            piece.pos = newPos
                            piece.value = button.value
        except:
            print(ans)
        break

def main(buttons, draw_nums, draw_chars, client, color):
    first = True
    turn = {}
    while True:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                pygame.quit()
            if event.type == pygame.MOUSEBUTTONDOWN:
                mouse_pos = pygame.mouse.get_pos()
                for button in buttons:
                    if button.is_under(mouse_pos):
                        if len(turn) == 0:
                            turn["pos1"] = button.value
                        else:
                            turn["pos2"] = button.value
                            move = json.dumps(turn)
                            client.send(bytes(move, 'UTF-8'))
                            t = threading.Thread(target=receive, args=(buttons, pieces, turn))
                            t.start()
                            turn = {}

        screen.fill(WHITE)
        for button in buttons:
            button.draw(screen)
        draw_nums()
        draw_chars()
        draw_GRID
        for piece in pieces:
            screen.blit(piece.img, piece.pos)
        pygame.display.update()

        if color == "black" and first:
            t = threading.Thread(target=receive_first, args=(buttons, pieces))
            t.start()
        first = False
        clock.tick(80)


if __name__ == '__main__':
    client = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client.connect((SERVER, PORT))
    color = client.recv(64).decode('UTF-8')

    if color == 'white':
        buttons = get_buttons(calculate_button_value_white)
        pieces = get_pieces(buttons)
        main(buttons, draw_nums_white_view, draw_chars_white_view, client, color)
    else:
        buttons = get_buttons(calculate_button_value_black)
        pieces = get_pieces(buttons)
        main(buttons, draw_nums_black_view, draw_chars_black_view, client, color)
