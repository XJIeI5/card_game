import asyncio
import json
import client_helper.reh as resp
import client_helper.json_parse as jsparce


class Client(asyncio.Protocol):
    def __init__(self, loop):
        self.loop = loop
        self.transport = None

    def disconnect(self):
        self.loop.stop()

    def connection_made(self, transport):
        self.transport = transport

    def data_received(self, data):
        j = json.loads(data)
        try:
            r = jsparce.GameResponce(**j)
        except TypeError:
            return
        handler = resp.ResponceHandlerFactory.get_handler(r.type)
        handler(self).handle(str(r.body))
        # print('Data received from server: \n===========\n{}\n===========\n'.format(data.decode()), flush=True)

    def send_data_to_tcp(self, data):
        self.transport.write(data.encode())

    def connection_lost(self, exc):
        print('\nThe server closed the connection')
        print('Stop the event loop')
        self.loop.stop()
