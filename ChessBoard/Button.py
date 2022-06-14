import pygame

class Button:
    def __init__(self, color, x, y, width, height, value):
        self.color = color
        self.x = x
        self.y = y
        self.width = width
        self.height = height
        self.value = value

    def draw(self, screen):
        pygame.draw.rect(screen, self.color, (self.x, self.y, self.width, self.height), 0)

    def is_under(self, pos):
        if pos[0] > self.x and pos[0] < self.x + self.width:
            if pos[1] > self.y and pos[1] < self.y + self.height:
                return True
        return False
