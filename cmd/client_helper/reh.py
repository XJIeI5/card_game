import abc


class ResponceHandler(metaclass=abc.ABCMeta):
    def __init__(self, tcp_client) -> None:
        self.client = tcp_client

    @abc.abstractmethod
    def handle(self, data):
        raise NotImplementedError()


class MessageHandler(ResponceHandler):
    def handle(self, data):
        print('Data received from server: \n===========\n{}\n===========\n'.format(data), flush=True)


class SetHandHandler(ResponceHandler):
    def handle(self, data):
        MessageHandler(self.client).handle(data)


class ResponceHandlerFactory:
    _types = {"message": MessageHandler,
              "set_hand": SetHandHandler,
              "error": MessageHandler}

    @classmethod
    def get_handler(cls, responce_type: str):
        responce_type = responce_type.strip()
        return cls._types.get(responce_type)
