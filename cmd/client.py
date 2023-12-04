from typing import Dict
import ctypes
import abc
import asyncio
from aioconsole import ainput


class GoString(ctypes.Structure):
    _fields_ = [("p", ctypes.c_char_p), ("n", ctypes.c_longlong)]


def get_gostring(line: str) -> GoString:
    return GoString(bytes(line, encoding='utf-8'), len(line))


class Request:
    def __init__(self, req_type: str, body: Dict):
        self.type = req_type
        self.body: Dict = body

    def __str__(self) -> str:
        type_str = f'"type":"{self.type}"'
        attr_list = []
        for key, value in self.body.items():
            match value:
                case bool() as bool_value:
                    curr_atr = f'"{key}":{str(bool_value).lower()}'
                case int() as int_value:
                    curr_atr = f'"{key}":{int_value}'
                case _:
                    curr_atr = f'"{key}":"{value}"'
            attr_list.append(curr_atr)
        attr_str = ','.join(attr_list)
        body_str = '"body":{' + attr_str + '}'
        return '{' + type_str + ',' + body_str + '}'


MENU = '''
---------------------------
start:\tStarts game
regist:\tRegist player with "name"
quit:\tDisconnect
---------------------------
'''

loop_ = asyncio.get_event_loop()


class Command(metaclass=abc.ABCMeta):
    asyn = False

    def __init__(self, tcp_client):
        self.client = tcp_client

    @abc.abstractmethod
    def run(self):
        raise NotImplementedError()


class StartGameCommand(Command):
    def run(self):
        req = str(Request("start_game", {}))
        self.client.send_data_to_tcp(req)


class RegistCommand(Command):
    asyn = True
    async def run(self):
        s = await ainput('name: ')
        req = str(Request("regist", {"name": s}))
        self.client.send_data_to_tcp(req)


class QuitCommand(Command):
    def run(self):
        print('Disconnected')
        self.client.disconnect()
        exit()


class PlayCardCommand(Command):
    asyn = True
    async def run(self):
        card_index = await int(ainput("peeked card: "))
        as_creature = await bool(ainput("as creature: "))
        if not as_creature:
            pass


class CommandFactory:
    _cmds = {'start_game': StartGameCommand,
             'regist': RegistCommand,
             'quit': QuitCommand}

    @classmethod
    def get_cmd(cls, cmd):
        cmd = cmd.strip()
        cmd_cls = cls._cmds.get(cmd)
        return cmd_cls


class Client(asyncio.Protocol):
    def __init__(self, loop):
        self.loop = loop
        self.transport = None

    def disconnect(self):
        self.loop.stop()

    def connection_made(self, transport):
        self.transport = transport

    def data_received(self, data):
        print('Data received from server: \n===========\n{}\n===========\n'.format(data.decode()), flush=True)

    def send_data_to_tcp(self, data):
        self.transport.write(data.encode())

    def connection_lost(self, exc):
        print('\nThe server closed the connection')
        print('Stop the event loop')
        self.loop.stop()


def menu():
    print(MENU)


async def main(client: Client):
    menu()
    while True:
        await asyncio.sleep(0.1)
        cmd = await ainput('>')
        cmd_cls = CommandFactory.get_cmd(cmd)
        if not cmd_cls:
            print(f'Unknown: {cmd}')
        elif cmd_cls.asyn:
            await cmd_cls(client).run()
        else:
            cmd_cls(client).run()


if __name__ == '__main__':
    client = Client(loop_)
    coro = loop_.create_connection(lambda: client, 'localhost', 9999)
    loop_.run_until_complete(coro)
    try:
        loop_.run_until_complete(main(client))
    except RuntimeError:
        pass
    except KeyboardInterrupt:
        client.disconnect()
