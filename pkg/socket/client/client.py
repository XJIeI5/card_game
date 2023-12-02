import socket
import ctypes


class GoString(ctypes.Structure):
    _fields_ = [("p", ctypes.c_char_p), ("n", ctypes.c_longlong)]


class Client:
    def __init__(self, name: str) -> None:
        self.name: str = name

    def listen(self, port: int) -> None:
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
            sock.connect(("localhost", port))
            data = '{"name":"' + self.name + '"}'
            sock.sendall(bytes(data, encoding='utf-8'))
            while True:
                data = sock.recv(1024)
                if data:
                    print(data)

    def host(self, port: int) -> None:
        lib = ctypes.cdll.LoadLibrary("./host.so")
        lib.Host.argtypes = [GoString]
        address = "localhost:" + str(port)
        lib.Host(GoString(bytes(address, encoding='utf-8'), len(address)))
        self.listen(port)


def main():
    client = Client(input("enter your name:"))
    command = input("host or connect:")
    if command == "host" or command == "h":
        client.host(int(input("port:")))
    elif command == "connect" or command == "c":
        client.listen(int(input("port:")))


if __name__ == "__main__":
    main()
