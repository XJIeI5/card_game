import abc
from aioconsole import ainput
import client_helper.req as r


MENU = '''
---------------------------
regist:\t\tRegist player with "name"
start:\t\tStarts game
playcard:\tPlays card
quit:\t\tDisconnect
---------------------------
'''


def menu():
    print(MENU)


class Command(metaclass=abc.ABCMeta):
    asyn = False

    def __init__(self, tcp_client):
        self.client = tcp_client

    @abc.abstractmethod
    def run(self):
        raise NotImplementedError()


class StartGameCommand(Command):
    def run(self):
        req = str(r.Request("start_game", {}))
        self.client.send_data_to_tcp(req)


class RegistCommand(Command):
    asyn = True
    async def run(self):
        s = await ainput('name: ')
        req = str(r.Request("regist", {"name": s}))
        self.client.send_data_to_tcp(req)


class QuitCommand(Command):
    def run(self):
        print('Disconnected')
        self.client.disconnect()
        exit()


class PlayCardCommand(Command):
    asyn = True
    async def run(self):
        card_index = await ainput("peeked card: ")
        card_index = int(card_index)
        req = str(r.Request("play_card", {"card_index": card_index, "as_creature": True}))
        self.client.send_data_to_tcp(req)


class CommandFactory:
    _cmds = {'start': StartGameCommand,
             'regist': RegistCommand,
             'quit': QuitCommand,
             'playcard': PlayCardCommand}

    @classmethod
    def get_cmd(cls, cmd):
        cmd = cmd.strip()
        cmd_cls = cls._cmds.get(cmd)
        return cmd_cls
