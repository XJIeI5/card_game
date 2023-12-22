import ctypes
import asyncio
from aioconsole import ainput
from client_helper import comand, client_logic


class GoString(ctypes.Structure):
    _fields_ = [("p", ctypes.c_char_p), ("n", ctypes.c_longlong)]


def get_gostring(line: str) -> GoString:
    return GoString(bytes(line, encoding='utf-8'), len(line))


loop_ = asyncio.get_event_loop()


async def main(client: client_logic.Client):
    comand.menu()
    while True:
        await asyncio.sleep(0.1)
        cmd = await ainput('>')
        cmd_cls = comand.CommandFactory.get_cmd(cmd)
        if not cmd_cls:
            print(f'Unknown: {cmd}')
        elif cmd_cls.asyn:
            await cmd_cls(client).run()
        else:
            cmd_cls(client).run()


if __name__ == '__main__':
    client = client_logic.Client(loop_)
    coro = loop_.create_connection(lambda: client, 'localhost', 9999)
    loop_.run_until_complete(coro)
    try:
        loop_.run_until_complete(main(client))
    except RuntimeError:
        pass
    except KeyboardInterrupt:
        client.disconnect()
