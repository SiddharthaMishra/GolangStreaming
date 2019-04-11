import asyncio
import websockets
import time
import cv2
import base64

async def hello():

    cap = cv2.VideoCapture(0)
    async with websockets.connect(
        'ws://localhost:3000/broadcaster') as websocket:
        while True:
            retval, img = cap.read()
            retval, buffer = cv2.imencode('.png', img)
            img = base64.b64encode(buffer)
            await websocket.send(img)
            await asyncio.sleep(0.01)

asyncio.get_event_loop().run_until_complete(hello())