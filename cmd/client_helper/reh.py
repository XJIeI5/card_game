import abc


class ResponceHandler(metaclass=abc.ABCMeta):
    def __init__(self, tcp_client) -> None:
        self.client = tcp_client

    @abc.abstractmethod
    def handle(self, data):
        raise NotImplementedError()


class EchoInStdoutHandler(ResponceHandler):
    def handle(self, data):
        print('Data received from server: \n===========\n{}\n===========\n'.format(data), flush=True)


class SetHandHandler(ResponceHandler):
    def handle(self, data):
        EchoInStdoutHandler(self.client).handle(data)


class ResponceHandlerFactory:
    _types = {"message": EchoInStdoutHandler,
              "set_hand": SetHandHandler}

    @classmethod
    def get_handler(cls, responce_type: str):
        responce_type = responce_type.strip()
        return cls._types.get(responce_type)
