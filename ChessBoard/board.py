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
checkmate_color = ""

whitePawn = pygame.image.load("./static/whitePawn.png")
blackPawn = pygame.image.load("./static/blackPawn.png")
whiteRook = pygame.image.load("./static/whiteRook.png")
blackRook = pygame.image.load("./static/blackRook.png")
whiteKnight = pygame.image.load("./static/whiteKnight.png")
blackKnight = pygame.image.load("./static/blackKnight.png")
whiteBishop = pygame.image.load("./static/whiteBishop.png")
blackBishop = pygame.image.load("./static/blackBishop.png")
whiteQueen = pygame.image.load("./static/whiteQueen.png")
blackQueen = pygame.image.load("./static/blackQueen.png")
whiteKing = pygame.image.load("./static/whiteKing.png")
blackKing = pygame.image.load("./static/blackKing.png")


def get_pieces(buttons):
    pieces = []
    for button in buttons:
        color = ""
        if button.value[1] == "1" or button.value[1] == "8":
            if button.value[0] == "A" or button.value[0] == "H":
                if button.value[1] == "1":
                    pieces.append(Piece(whiteRook, (button.x, button.y), button.value, "white", "rook"))
                else:
                    pieces.append(Piece(blackRook, (button.x, button.y), button.value, "white", "rook"))
            elif button.value[0] == "B" or button.value[0] == "G":
                if button.value[1] == "1":
                    pieces.append(Piece(whiteKnight, (button.x, button.y), button.value, "white", "knight"))
                else:
                    pieces.append(Piece(blackKnight, (button.x, button.y), button.value, "white", "knight"))
            elif button.value[0] == "C" or button.value[0] == "F":
                if button.value[1] == "1":
                    pieces.append(Piece(whiteBishop, (button.x, button.y), button.value, "white", "bishop"))
                else:
                    pieces.append(Piece(blackBishop, (button.x, button.y), button.value, "white", "bishop"))
            elif button.value[0] == "D":
                if button.value[1] == "1":
                    pieces.append(Piece(whiteQueen, (button.x, button.y), button.value, "white", "queen"))
                else:
                    pieces.append(Piece(blackQueen, (button.x, button.y), button.value, "white", "queen"))
            elif button.value[0] == "E":
                if button.value[1] == "1":
                    pieces.append(Piece(whiteKing, (button.x, button.y), button.value, "white", "king"))
                else:
                    pieces.append(Piece(blackKing, (button.x, button.y), button.value, "white", "king"))
        elif button.value[1] == "2":
            pieces.append(Piece(whitePawn, (button.x, button.y), button.value, "white", "pawn"))
        elif button.value[1] == "7":
            pieces.append(Piece(blackPawn, (button.x, button.y), button.value, "black", "pawn"))
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

def receive(buttons, pieces, turn, color):
    global checkmate_color
    while True:
        ans = client.recv(64).decode('UTF-8')
        if "SET" in ans and "Pos" not in ans:
            if "1" in ans:
                for piece in pieces:
                    if piece.value == turn['pos2']:
                        piece.pos = (1000, 1000)
            for piece in pieces:
                if piece.value == turn["pos1"]:
                    for button in buttons:
                        if button.value == turn["pos2"]:
                            newPos = (button.x, button.y)
                            piece.pos = newPos
                            piece.value = button.value
            continue
        elif "CHECKMATE" in ans and "Pos" not in ans:
            checkmate_color = "opponent"
            for piece in pieces:
                if piece.value == turn["pos1"]:
                    for button in buttons:
                        if button.value == turn["pos2"]:
                            newPos = (button.x, button.y)
                            piece.pos = newPos
                            piece.value = button.value
            break
        try:
            ans = json.loads(ans)
            if len(ans) == 4 or len(ans) == 3:
                print(ans['Option1'])
                if ans['Option1'] == "CHECKMATE":
                    checkmate_color = "me"
                if ans['Option2'] == "1":
                    for piece in pieces:
                        if piece.value == ans['Pos2']:
                            piece.pos = (1000, 1000)
            for piece in pieces:
                if piece.value == ans["Pos1"]:
                    for button in buttons:
                        if button.value == ans["Pos2"]:
                            newPos = (button.x, button.y)
                            piece.pos = newPos
                            piece.value = button.value
        except:
            print("except")
        break

def main(buttons, draw_nums, draw_chars, client, color):
    first = True
    turn = {}
    while True:
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                pygame.quit()
                sys.exit()
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
                            t = threading.Thread(target=receive, args=(buttons, pieces, turn, color))
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
        if checkmate_color == "opponent":
            message_to_screen("CHECKMATE! you win!", (0, 255, 0), (50, 50), Font)
        elif checkmate_color == "me":
            message_to_screen("CHECKMATE! you loose!", (255, 0, 0), (50, 50), Font)
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
