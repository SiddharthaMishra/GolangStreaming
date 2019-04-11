import asyncio
import websockets
import cv2
import time
import pygame
import pygame.camera
import base64

async def hello():
    pygame.camera.init()
    pygame.camera.list_cameras() #Camera detected or not
    cam = pygame.camera.Camera("/dev/video0",(300,300))
    cam.start()
    async with websockets.connect(
        'ws://localhost:3000/broadcaster') as websocket:
        while True:
            for i in range(10):
                await websocket.send(base64.b64encode(pygame.image.tostring(cam.get_image(), 'RGBA')))
                await asyncio.sleep(1)
def world():
#    async with websockets.connect(
#        'ws://localhost:3000/broadcaster') as websocket:

    while True:
        pygame.camera.init()
        cam = pygame.camera.Camera("/dev/video0",(640,480))
        cam.start()
        img = cam.get_image()
        print(img)
        time.sleep(1)

#def test():

#hello()
asyncio.get_event_loop().run_until_complete(hello())