from settings import * 
from Tetris import tetris
import sys

class App:
    def __init__(self):
        pg.init()
        pg.display.set_caption('Tetris')
        self.screen = pg.display.set_mode(FIELD_RES)
        self.clock = pg.time.Clock()
        self.tetris = tetris(self)
        
    def update(self):
        self.tetris.update()
        self.clock.tick(fps)
        
    def draw(self):
        self.screen.fill(color = FIELD_COLOR)
        self.tetris.draw()
        pg.display.flip()
        
        
        
    def check_event(self):
        for event in pg.event.get():
            if event.type == pg.QUIT or (event.type == pg.KEYDOWN and event.key == pg.K_ESCAPE):
                pg.quit()
                sys.exit()
            elif event.type == pg.KEYDOWN:
                self.tetris.control(pressed_key=event.key)
                  
                
    def run(self):
        while True:
            self.check_event()
            self.update()
            self.draw()
            
    
if __name__ == '__main__':
    app = App()
    app.run()
    

    